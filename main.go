package main

import ("SecondProjectInterfaceHome/cannals"
"SecondProjectInterfaceHome/modul")


func main() {
	massNotification := []cannals.Notification{
{		Type: "email", Adress: "nhejhejdjebhs@jjh", Text:"dedennsnjhhefj"},
	{Type: "push", Adress: "sdsddsdssddssdsd@jjh", Text:"dsddsdawqwwwwwwwwqqqqqqqqqqqqqqqqqqqqqqhefj"},
	{Type: "", Adress: "dsdssd@jjh", Text:"dedensddssdsddssssssssssssssssssssssssssssssssssssnsnjhhefj"},
}
	modul.MassText(massNotification)
}
