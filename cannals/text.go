package cannals

import (
	
	"fmt"
)

type Notification struct {
	Type   string
	Adress string
	Text   string
	
}

type Email struct {
}

func (c *Email) Send(n Notification) error {
	fmt.Println("Отправка email-уведомления по адресу: ", n.Adress)
	fmt.Println("Текст email-уведомления: ", n.Text)
	return nil
}

type Push struct {
}

func (c *Push) Send(n Notification) error{
	fmt.Println("Отправка Push-уведомления по адресу: ", n.Adress)
	fmt.Println("Текст Push-уведомления: ", n.Text)
	return nil
}

type Sms struct{
	
}

func (c *Sms) Send(n Notification) error {
	fmt.Println("Отправка sms-уведомления по адресу: ", n.Adress)
	fmt.Println("Текст sms-уведомления: ", n.Text)
	return nil
}
