package postman

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

func Postman(wg *sync.WaitGroup, ctx context.Context,transferPoint chan<- string, n int, mail string) {

	defer wg.Done() //уменьшает на 1 

	for {
		select{
		case <-ctx.Done():
			pp.Println("Я postman номер", n, "закончил работу")
				return 
		default:		
				pp.Println("Я postman номер", n, "взял письмо")
				time.Sleep(1 * time.Second)
				pp.Println("Я postman номер", n, "донес до почты")
				transferPoint <- mail
				pp.Println("Я postman номер", n, "передал письмо", mail)
		}

	}
}

func PostmanPool(ctx context.Context,postmanNumber int, mail string) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanNumber; i++ {
		wg.Add(1)
		go Postman(wg, ctx, mailTransferPoint, i, mailSend(postmanNumber))
	}

	go func(){
		wg.Wait()
		fmt.Println("end miner await")
		close(mailTransferPoint) //закрываем канал
	}()

	return mailTransferPoint
}

func mailSend(postmanNumber int) string {
		postmanMap := map[int]string{
		1: "Журнал",
		2: "Журнал 2",
		3: "Журнал 3",
	}
	mail, ok := postmanMap[postmanNumber]
	if (!ok){
		mail = "Письмо"
	}
	return mail
}