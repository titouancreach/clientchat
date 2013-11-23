package bot

import (
	"../api"
	"../parser"
	"fmt"
	"strings"
	"time"
)

type Bot struct {
	connectionToken  *api.HandleConnection
	lastMessageToken *api.HandleMessage
	totalMessage     int
	pseudo           string
}

func (this *Bot) Run() error {

	var err error

	this.totalMessage = 1

	this.pseudo = "login"

	this.connectionToken, err = api.Connection(this.pseudo, "passw", "loung")


	if err != nil {
		return err
	}

	stop := this.Timer(1 * time.Second)

	for true {

		var userInput string

		fmt.Scan(&userInput)

		if userInput == "/QUIT" {
			break
		}

		this.SendMessage(userInput)
	}

	stop <- true

	return nil
}

func (this *Bot) AutomaticAnswer(msg []string) error {

	var err error
	err = nil

	for ite := range msg {

		line := msg[ite]

		isHl, sender := parser.IsHL(line, this.pseudo)


		//It's what append when something HL us
		if isHl {

			err = this.SendMessage(sender + "> :kiss: :fleur:")
			if err == nil {
				println("Message sended back to " + sender)
			}
		}

		isNoticed, sender := parser.IsNoticed(line)

		//Noticed us
		if isNoticed {
			err = this.SendMessage("/NOTICE " + sender + " :rollepi:")
			if err == nil {
				println("Noticed sended back to " + sender)
			}
		}

		isJoined, sender := parser.IsJoined(line)

		//Or join current loung
		if isJoined {
			err = this.SendMessage(sender + "> Salut " + sender + " :salut: ")
			if err == nil {
				println("Joined sended to : " + sender)
			}
		}
	}

	return err
}

func (this *Bot) SendMessage(msg string) error {

	this.totalMessage++

	var err error

	this.lastMessageToken, err = api.SendMessage(this.connectionToken.Session, msg, this.totalMessage)

	//If message is not sended, ask for a new connection
	if err != nil {
		this.connectionToken, err = api.ConnectionAuto(this.pseudo, this.connectionToken.Passwd)
	}

	return err
}

func (this *Bot) CheckForNewMessages() error {

	this.totalMessage++

	handleRefresh, err := api.Refresh(this.connectionToken.Session, this.totalMessage)

	if err == nil {
		loungMsg := parser.ParseMessage(handleRefresh.Loung)

		//New message on the loung
		if loungMsg != "" {

			print(loungMsg)
			msg := strings.Split(loungMsg, "\n")
			err = this.AutomaticAnswer(msg)
		}

	} else {
		fmt.Println(err)
		this.connectionToken, err = api.ConnectionAuto(this.pseudo, this.connectionToken.Passwd)
	}

	return err
}

func (this *Bot) Timer(delay time.Duration) chan bool {

	stop := make(chan bool)

	go func() error {
		for {

			err := this.CheckForNewMessages()

			if err != nil {
				fmt.Println(err)
			}

			select {

			case <-time.After(delay):

			case <-stop:
				return nil
			}
		}
	}()

	return stop
}
