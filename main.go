// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"math/rand/v2"
// 	"strconv"
// 	"study/greeting"
// 	"sync"
// 	"time"

// 	"github.com/k0kubun/pp"
// )

// type Person struct {
// 	Name   string
// 	Age    int
// 	Blance int
// }

// func (p Person) GetName() string {
// 	return p.Name
// }

// // 2 указателя под капотом на тип объекта и сам объект
// type User interface {
// 	GetName() string
// }

// func CheckUser(u User) {
// 	fmt.Printf(`функция которая принимает интерфейс User %v`, u.GetName(), "\n")
// 	fmt.Println("User Name:", u.GetName())
// }

// func Pay(user *Person, usd int) (int, error) {
// 	random := rand.IntN(100)
// 	pp.Println(random)
// 	if user.Blance-usd < 0 {
// 		return 0, errors.New("Недостаточно средств!")
// 	}

// 	return random, nil
// }

// func main() {

// 	defer func() {
// 		panic := recover()

// 		if panic != nil {
// 			fmt.Println("Произошла ошибка:", panic)
// 		}
// 	}()

// 	var person = Person{
// 		Name:   "Alice",
// 		Age:    30,
// 		Blance: 300,
// 	}
// 	val, err := Pay(&person, 100)
// 	if err != nil {
// 		fmt.Println("Ошибка при оплате:", err)
// 	} else {
// 		fmt.Println("Оплата прошла успешно, ваш код:", val)
// 	}

// 	greeting.SayGf()

// 	checkList := map[string]int{
// 		"apple":  5,
// 		"banana": 10,
// 		"orange": 7,
// 	}

// 	value, error := checkList["apple2"]

// 	if error != true {
// 		pp.Println("Error:", error)
// 	} else {
// 		fmt.Println("Success")
// 	}

// 	fmt.Println(value, error)

// 	whatAmI := func(item any) {
// 		switch t := item.(type) {
// 		case bool:
// 			fmt.Println("I'm a bool")
// 		case int:
// 			fmt.Println("I'm an int")
// 		default:
// 			fmt.Printf("Don't know type %T\n", t)
// 		}
// 	}
// 	whatAmI(true)
// 	whatAmI("hey")

// 	intChan := make(chan int)

// 	go func() {
// 		intChan <- 666
// 	}()

// 	select {
// 	case msg := <-intChan:
// 		fmt.Println(msg)
// 	}

// 	mapa := map[string]int{
// 		"dd": 21,
// 		"ff": 22,
// 	}
// 	mapa["x"] = 2
// 	val, ok := mapa["yy"]
// 	if ok != true {
// 		fmt.Println("Error:", ok)
// 		pp.Println("mapa", true)
// 	}

// 	type Message struct {
// 		Name string
// 	}
// 	var robot = []any{}

// 	//! Закрытие каналов. Аксиомы каналов.
// 	// создаю новый открытый канал (канал нужен для передачи данных между каналами)
// 	var messageChan = make(chan string)

// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		for i := 0; i < 6; i++ {
// 			messageChan <- "name" + strconv.Itoa(i)
// 		}
// 		//! закрываем канал
// 		close(messageChan)
// 	}()

// 	// // блокируем поток и ждем значение от канала
// 	// v1, ok := <-messageChan
// 	// fmt.Println(v1, ok)

// 	//! как только канал будет закрыт тогда автоматически цикл range transerPoint завершится и не нужно следить за ok(статусом)
// 	for mess := range messageChan {
// 		robot = append(robot, mess)
// 	}
// 	pp.Println("robot", robot)

// 	//*1) Context -----------------------------------------
// 	// можем отдельно контролировать выполнение каких-то узлов нашей программы (те засунуть в контекст группу go-рутин и когда нам нужно завершить их закрываем контекст)

// 	parentContext, parentCloseContext := context.WithCancel(context.Background())
// 	go foo(parentContext)
// 	time.Sleep(3 * time.Second)
// 	// закрываем контекст (группу гоурутин)
// 	parentCloseContext()
// 	time.Sleep(3 * time.Second)

// 	//!2) инструмент синхронизации WaitGroup -------------------------------------(дождаться окончания выполнения)
// 	// ждем пока все горутины завершатся, и потом выполняем код в Main
// 	// а через каналы синхронизация происходит когда рутина отдает какие-то данные!

// 	//? берем указатель на WaitGroup (она как каналы под капотом сама не делает указатель на себя)
// 	wg := &sync.WaitGroup{}

// 	wg.Add(3) //? обязательно(счетчик) - говорим щас запущу 2 горутину
// 	postman_WaitGrop("новости", wg)
// 	postman_WaitGrop("Auto", wg)
// 	postman_WaitGrop("Sport", wg)

// 	//? блокируемся на вызове пока счетчик не станет 0
// 	wg.Wait()
// 	pp.Print("test")
// 	//! -------------------------------------

// 	//*3) состояние гонки атомики, Mutex -------------------------------------
// 	// Если есть какой-то конкурентный доступ к какой-то переменной его нужно облакдывать Mutex чтобы не происходила состояние гонки
// 	wg.Add(3)

// 	go inc_Mutex(wg)
// 	go inc_Mutex(wg)
// 	go inc_Mutex(wg)

// 	wg.Wait()

// 	pp.Println("number", number)

// 	//* -------------------------------------

// 	//*4) состояние гонки RWMutex -------------------------------------
// 	wg.Add(1)
// 	go setLike(wg)

// 	for i:=1; i<=3; i++{
// 		wg.Add(1)
// 		go getLike(wg, "1")
// 		wg.Add(1)
// 		go getLike(wg , "2")
// 		wg.Add(1)
// 		go getLike(wg , "3")
// 	}

// 	wg.Wait()

// 	fmt.Println("likes", likes)

// 	//* -------------------------------------

// }

// var number int = 0
// var mutex = sync.Mutex{}

// func inc_Mutex(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 10000; i++ {
// 		// тут как бы мы ставим блокировку гарантируя, что только одна go изменяет number
// 		mutex.Lock()
// 		number += 1
// 		mutex.Unlock()
// 	}
// }

// var likes int = 0
// var mtx sync.RWMutex

// func setLike(wg *sync.WaitGroup){
// 	defer wg.Done()

// 	for i:=1; i<=100; i++{
// //! только 1 Go-рутина может щас выполнять этот код
// 		mtx.Lock()
// 		pp.Println( "likes", likes,"setLike ++", i)
// 		likes++
// 		mtx.Unlock()
// //! Go планировщик решает, какая горутина следующей возьмёт мьютекс.
// 	}
// }
// func getLike(wg *sync.WaitGroup, name string){
// 	defer wg.Done()

// 	for i:=1; i<=100; i++{
// 		mtx.RLock()
// 	    pp.Println("я взял выполнение мьютекса", name, likes)
// 		// _=likes
// 		mtx.RUnlock()
// 	}
// }

// func postman_WaitGrop(text string, wg *sync.WaitGroup) {
// 	defer wg.Done() //? Важно (счетчик) каждый раз уменьшает на 1
// 	// когда func запущеная в горутине завершится, всегда в конце вызывается defer
// 	for i := 1; i <= 3; i++ {
// 		pp.Println("Я понес газету", text, i)
// 		time.Sleep(250 * time.Millisecond)
// 	}
// }

// func foo(ctx context.Context) {
// 	// в бесконечном цикле for будем следить отменен ли контекст или нет через select
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			pp.Println("контекст удален")
// 			return
// 		default:
// 			fmt.Println("foo")
// 		}
// 		time.Sleep(100 * time.Millisecond)

// 	}

// }

package main

import (
	"context"
	"fmt"
	"study/miner"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)
func main(){
	// -race детект состояние гонки
var resCoal int
minerCotext, minecrCancel := context.WithCancel(context.Background())
coalTransferPoint := miner.MinerPool(minerCotext, 1)

var mtx = sync.Mutex{}
var wg = &sync.WaitGroup{}

go func(){
	time.Sleep(3 * time.Second)
	minecrCancel()
}()

wg.Add(1)
go func(){
	defer wg.Done()

	for val:= range coalTransferPoint {
	pp.Println("Количество угля в складе", val)
	mtx.Lock()
	fmt.Println("lock ->>>")
	time.Sleep(5 * time.Second)
	resCoal += val
	mtx.Unlock()	
}
}()

wg.Add(1)
go func(){
	defer wg.Done()

	for val:= range coalTransferPoint {
	pp.Println("Количество угля в складе", val)
	mtx.Lock()
	fmt.Println("lock 2 ->>>")
	time.Sleep(5 * time.Second)
	resCoal += val
	mtx.Unlock()	
}
}()

go func(){
	for{
	time.Sleep(1 * time.Second)
	fmt.Println("read: ->>>", resCoal)
	}
}()

wg.Wait()

// var isCoalClose bool = false

// for !isCoalClose{
// 	select{
// 	case coal, ok := <-coalTransferPoint:
// 	if !ok {
// 		isCoalClose = true
// 		pp.Println("Уголь закончился")
// 	}
// 	resCoal += coal

// 	}
// }


minecrCancel()
fmt.Println("Общее количество угля", resCoal)
}