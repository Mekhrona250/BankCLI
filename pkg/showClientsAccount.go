package pkg

import "fmt"

func ShowClientsAccount() {
	var name string

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	var balance float64
	var has bool
	// проверка на наличие клиента
	for _, val := range Accounts {
		if name == val.Name {
			balance = val.Balance
			has = true
		}
	}

	if !has {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
		return
	}

	fmt.Printf("Баланс счета %v является %v \n", name, balance)
}
