package main

import (
  "fmt"
  "time"
)

  var database = make(map[string]int)
  func AddClient() {
  var name string
  fmt.Scan(&name)
  database[name] = 0

	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")

}
 

	func account()  {
	var name string
	fmt.Scan(&name)	

	var n int
	fmt.Scan(&n)
  
	if n > 0 {
		database[name]+=n
	}
	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")
}

func ShowAccount() {
	var name string
	fmt.Scan(&name)	
	fmt.Println("Ваш баланс: ")
	fmt.Println(database[name])
	fmt.Println("______________")
}

func WithdrawMoney() {
	var name string
	fmt.Scan(&name)	
	var minus int
	fmt.Scan(&minus)
	if minus > database[name] {
		fmt.Println("Недостаточно средств")
	} else {
		database[name]-=minus
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