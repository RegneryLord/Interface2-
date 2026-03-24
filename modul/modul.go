package modul

import (
	"SecondProjectInterfaceHome/cannals"
	"fmt"
)

type Noti interface {
	Send(notification cannals.Notification) error
}

func GetType(change string)Noti{
switch change{
case "email":
	return &cannals.Email{}
case "push":
	return &cannals.Push{}
case "sms":
	return &cannals.Sms{}
default: 
return nil
}

}

// функция для отправки массива уведомлений
func MassText(notification []cannals.Notification){
	for _,v:=range notification{
		text:=GetType(v.Type)//смотрит какой тип мы написали sms email...
		if text == nil{
			fmt.Println("Нет обработчика для канала", v.Type)//не может найти тип, который есть у нас(push sms email)
			continue
		}
		err:= text.Send(v)
		if err!=nil{
			fmt.Println("Ошибка при отправке", err)
		}
	}
}


