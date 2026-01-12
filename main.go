package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"study/greeting"
	"time"

	"github.com/k0kubun/pp"
)

type Person struct {
	Name   string
	Age    int
	Blance int
}

func (p Person) GetName() string {
	return p.Name
}

// 2 указателя под капотом на тип объекта и сам объект
type User interface {
	GetName() string
}

func CheckUser(u User) {
	fmt.Printf(`функция которая принимает интерфейс User %v`, u.GetName(), "\n")
	fmt.Println("User Name:", u.GetName())
}

func Pay(user *Person, usd int) (int, error) {
	random := rand.IntN(100)
	pp.Println(random)
	if user.Blance-usd < 0 {
		return 0, errors.New("Недостаточно средств!")
	}

	return random, nil
}

func main() {

	defer func() {
		panic := recover()

		if panic != nil {
			fmt.Println("Произошла ошибка:", panic)
		}
	}()

	var person = Person{
		Name:   "Alice",
		Age:    30,
		Blance: 300,
	}
	val, err := Pay(&person, 100)
	if err != nil {
		fmt.Println("Ошибка при оплате:", err)
	} else {
		fmt.Println("Оплата прошла успешно, ваш код:", val)
	}

	greeting.SayGf()

	checkList := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 7,
	}

	value, error:= checkList["apple2"]

	if error != true {
		pp.Println("Error:", error)
	} else {
		fmt.Println("Success")
	}

	fmt.Println(value, error)

	
	

	whatAmI := func(item any) {
		switch t := item.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI("hey")

	intChan := make(chan int)

	go func() {
		intChan <- 666
	}()

	select {
		case msg := <-intChan:
			fmt.Println(msg)
	}

	mapa := map[string]int{
		"dd": 21,
		"ff": 22,
	}
	mapa["x"] = 2
	val, ok := mapa["yy"]
	if ok != true {
		fmt.Println("Error:", ok)
		pp.Println("mapa", true)
	}




	type Message struct {
		Name string
	}
	var robot = []any{}
	// создаю новый открытый канал
	var messageChan = make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		for i := 0; i < 6; i++ {
			messageChan <- "name" + strconv.Itoa(i)
		}
		// закрываем канал
		close(messageChan)
	}()

	// // блокируем поток и ждем значение от канала
	// v1, ok := <-messageChan
	// fmt.Println(v1, ok)

	// как только канал будет закрыт тогда автоматически цикл range transerPoint завершится и не нужно следить за ok(статусом)
	for mess := range messageChan {
		robot = append(robot, mess)
	}

	pp.Println("robot", robot)

}
