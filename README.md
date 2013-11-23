Client pour http://chat.developpez.com
=========


this project is the begining for a client for developpez.com
not letting me finish it, I post sources on my repisitory in order to a
fearless man continue it.


##PACKAGE DOCUMENTATION

###package parser

####FUNCTIONS

** func IsHL(msg string, pseudo string) (bool, string) ** 

return value

1st is true when message has adressed to us
2th returned value it's the nickname of the message sender

parameter

msg : The parsed message from the loung
pseudo : us pseudo


**func IsJoined(msg string) (bool, string) ** 

like isHL but return true if someone joined, and return his nickname


**func IsNoticed(msg string) (bool, string) ** 

the same if someone noticed you


**func ParseMessage(msg string) string ** 

Parse html msg to human reading message, and return it


--------------------------------------------------------------------------------------------------------

###package api

####TYPES

`

type HandleConnection struct {
    Session string  `json:"session"`
    Passwd  string  `json:"motdepasse"`
    Salon   string  `json:"nomSalon"`
    Color   string  `json:"couleur"`
    State   float64 `json:"etat"`
    Message string  `json:"message"`
}

`

Connection handle, directly parsed from JSON


** func Connection(login string, passwd string, salon string) (*HandleConnection, error) **

Ask for a new Connection


** func ConnectionAuto(login string, passwd string) (*HandleConnection, error) **

Ask for an automatic connextion (if you're were deconnected)



type HandleMessage struct {
    State   int    `json:"etat"`
    Message string `json:"message"`
    Lounge  string `json:"salon"`
}

Message response from server, directly parsed from JSON


** func SendMessage(session string, message string, totalMessages int) (*HandleMessage, error) **

Send message and return an the response of the server


type HandleRefresh struct {
    State   int         `json:"etat"`
    Message string      `json:"message"`
    Private interface{} `json:"pvs"`
    Loung   string      `json:"salon"`
}

the response of a refresh question


** func Refresh(session string, totalMessages int) (*HandleRefresh, error) **

give you new messages by the HandleRefresh type. you can find totalMessages in Bot.totalMessages


###package bot


####TYPES

type Bot struct {
    // contains filtered or unexported fields
}

bot information, view sources for any informations


** func (this *Bot) AutomaticAnswer(msg []string) error **

check if the message is addressed to you


** func (this *Bot) CheckForNewMessages() error **

print if new messages sended



** func (this *Bot) Run() error **

mainLoop


** func (this *Bot) SendMessage(msg string) error **

Just send a message using api package


** func (this *Bot) Timer(delay time.Duration) chan bool **

goroutine for checkForNewMessages



