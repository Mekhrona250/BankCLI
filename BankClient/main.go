package main

import (
  "fmt"
  "time"
)

var database = make(map[string]int)
var client_account int = 0

func AddClient() {
  var name string
  var age int

  fmt.Scan(&name)

  fmt.Scan(&age)

  database[name] = 2023 - age

  fmt.Println("________________")
  fmt.Println("Готово")
  fmt.Println("________________")

}

func account()  {
	var n int
	fmt.Scan(&n)
   
	if n > 0 {
		client_account+=n
	}
	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")
}

func ShowAccount() {
	fmt.Println("Ваш баланс: ")
	fmt.Println(client_account)
	 fmt.Println("______________")
}

func WithdrawMoney() {
	var minus int
	fmt.Scan(&minus)
	if minus > client_account {
		fmt.Println("Недостаточно средств")
	} else {
		client_account-=minus
		fmt.Println("________________")
		fmt.Println("Готово")
		fmt.Println("________________")
	}

	
}


func main() {

  for {
    var choice int
    fmt.Println("1. Создать клиента")
    fmt.Println("2. Пополнить счет клиента")
    fmt.Println("3. Показать баланс клиента")
	fmt.Println("4. Снять деньги клиента")
	fmt.Println("5. Выйти")
	

    fmt.Scan(&choice)

    if choice == 1 {
      AddClient()
    } else if choice == 2 {
      account()
    } else if choice == 3 {
      ShowAccount()
    } else if choice == 4 {
		WithdrawMoney()
	} else if choice == 5 {
		break
	}

    time.Sleep(3 * time.Second)

  }

}