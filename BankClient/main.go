package main

import (
  "fmt"
  "time"
)
var database = make(map[string]int)

func AddClient(name string) {
	database[name] = 0

}
 

func account(name string)  {
	var n int
	fmt.Scan(&n)
  
	if n > 0 {
		database[name]+=n
	}
}

func ShowAccount(name string) {
	fmt.Print("Ваш баланс: ")
	fmt.Print(database[name], "\n")
	fmt.Println("______________")
}

func WithdrawMoney(name string) {
	var minus int
	fmt.Scan(&minus)
	if minus > database[name] {
		fmt.Println("Недостаточно средств")
	} else {
		database[name]-=minus
	}

	
}


func main() {
	

  for {
    var choice int
	var name string

    fmt.Println("1. Создать клиента")
    fmt.Println("2. Пополнить счет клиента")
    fmt.Println("3. Показать баланс клиента")
	fmt.Println("4. Снять деньги клиента")
	fmt.Println("5. Выйти")
	

    fmt.Scan(&choice)

	if choice >= 5 {
		break
	}

	fmt.Scan(&name)

    if choice == 1 {
    	AddClient(name)
    } else if choice == 2 {
    	account(name)
    } else if choice == 3 {
    	ShowAccount(name)
    } else if choice == 4 {
		WithdrawMoney(name)
	}
	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")

    time.Sleep(3 * time.Second)

  }

}