
import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	var resCoal int

	coalTransferPoint := make(chan int)
	var mtx =  sync.Mutex{}
	var wg = &sync.WaitGroup{}

	// контекст с отменой
	minerContext, minerCancel := context.WithCancel(context.Background())

	// ===== ШАХТЁР =====
	wg.Add(1)
	go func(n int) {
		defer wg.Done()

		for {
			select {
			case <-minerContext.Done():
				pp.Println("-------> Я шахтёр номер", n, "закончил работу")
				return

			case <-time.After(1 * time.Second):
				pp.Println("Я шахтёр номер", n, "начал добывать уголь")
				pp.Println("Я шахтёр номер", n, "добыл уголь")
				coalTransferPoint <- 1 // отправляем уголь в канал
				pp.Println("Я шахтёр номер", n, "отправил уголь на склад ©️")
			}
		}
	}(1)
	wg.Add(1)
	go func(n int) {
		defer wg.Done()

		for {
			select {
			case <-minerContext.Done():
				pp.Println("-------> Я шахтёр номер", n, "закончил работу")
				return

			case <-time.After(1 * time.Second):
				pp.Println("Я шахтёр номер", n, "начал добывать уголь")
				pp.Println("Я шахтёр номер", n, "добыл уголь")
				coalTransferPoint <- 1 // отправляем уголь в канал
				pp.Println("Я шахтёр номер", n, "отправил уголь на склад ®️")
			}
		}
	}(2)

	// ===== СКЛАД =====
	go func() {
		for val := range coalTransferPoint {
			mtx.Lock()
			resCoal += val
			fmt.Println("Склад получил уголь. Всего:", resCoal)
			mtx.Unlock()
		}
		fmt.Println("Склад закрыт")
	}()

	// ===== ОСТАНОВКА ШАХТЁРА =====
	go func() {
		time.Sleep(3 * time.Second)
		pp.Println("Завершаем работу шахтёра +++++++++++++++++")
		minerCancel()
	}()

	// ждём шахтёра
	wg.Wait()

	// закрываем канал, когда шахтёр закончил
	close(coalTransferPoint)

	fmt.Println("ИТОГО ДОБЫТО УГЛЯ:", resCoal)
}
