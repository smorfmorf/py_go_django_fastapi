package miner

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)


func Miner(
	ctx context.Context, //context
	wg *sync.WaitGroup,  // await 
	transferPoint chan <- int, //send only (только писать)
	n int,
	power int,
){
	defer wg.Done()

	for {
		select{
			case <-ctx.Done():
				pp.Println("-------> Я шахтёр номер", n, "закончил работу")
				return 
			default:
				pp.Println("Я шахтёр номер", n, "начал добывать уголь")
				time.Sleep(1 * time.Second)
				pp.Println("Я шахтёр номер", n, "добыл уголь")
				transferPoint <- power
				pp.Println("Я шахтёр номер", n, "отправил уголь на склад")
		}
	}
}

// пункт передачи  хотим только читать <-
func MinerPool(ctx context.Context, minerCount int) <-chan int{
	coalTransferPoint := make(chan int) // канал передачи угля
	wg := &sync.WaitGroup{} // чтобы подождать всех Mineroв

	for i:=1; i<=minerCount; i++{
		wg.Add(1)
		// Функция, описывающая работу одного отдельного шахтёра
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	// в отдельном потоке запускаем ожидание окончания Mineroв
	// чтобы не блокировать канал для передачи угля
	go func(){
		wg.Wait()
		fmt.Println("end miner await")
		close(coalTransferPoint) //закрываем канал
	}()


	return coalTransferPoint
}