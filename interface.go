/*	Пример работы с интерфейсом. Структуры Wallet, Card и ApplePay реализуют интерфейс Payer.
Функция Buy принимает на вход пустой интерфейс и пытается преобразовать его к плательщику (Payer).
Если преобразование не удалось, то выводим информацию о том, что входные данные платежным средством
не являются */

package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

type Card struct {
	Balance    int
	ValidUntil string
	Cardholder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("не хватает денег на карте")
	}
	c.Balance -= amount
	return nil
}

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("не хватает денег на аккаунте")
	}
	a.Money -= amount
	return nil
}

func Buy(in interface{}) {
	var p Payer
	var ok bool
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T не не является платежным средством\n\n", in)
		return
	}
	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

// func Buy(p Payer) {
// 	switch p.(type) {
// 	case *Wallet:
// 		fmt.Println("Оплата наличными?")
// 	case *Card:
// 		plasticCard, ok := p.(*Card)
// 		if !ok {
// 			fmt.Println("Не удалось преобразовать к типу *Card")
// 		}
// 		fmt.Println("Вставляйте карту,", plasticCard.Cardholder)
// 	default:
// 		fmt.Println("Что-то новое!")
// 	}
// 	err := p.Pay(10)
// 	if err != nil {
// 		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
// 		return
// 	}
// 	fmt.Printf("Спасибо за покупку через %T\n\n", p)
// }

func main() {

	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, Cardholder: "rvasily"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)
}
