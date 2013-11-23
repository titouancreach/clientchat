package main

import "./bot"
import "fmt"

func main () {

	var me bot.Bot

	fmt.Println(me.Run())

}
