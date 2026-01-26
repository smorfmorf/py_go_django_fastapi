package miner

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

func Miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan <- int, //send only
	n int,
	power int,
){
	defer wg.Done()

	for {
		select{
			case <-ctx.Done():
				pp.Println("Я шахтёр номер", n, "закончил работу")
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

// пункт передачи 
func MinerPool(ctx context.Context, minerCount int) <-chan int{
	coalTransferPoint := make(chan int) // пункт передачи 
	wg := &sync.WaitGroup{} // чтобы подождать всех Mineroв

	for i:=1; i<=minerCount; i++{
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	// в отдельном потоке запускаем ожидание окончания Mineroв
	go func(){
		wg.Wait()
		fmt.Println("test check")
		close(coalTransferPoint) //закрываем канал
	}()

	return coalTransferPoint
}