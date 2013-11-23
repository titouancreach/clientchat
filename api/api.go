package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
//	"fmt"
)

const (
	ajaxUrl string = "http://chat.developpez.com/ajax.php"
	version string = "2.0.4"
)

type HandleConnection struct {
	Session string  `json:"session"`
	Passwd  string  `json:"motdepasse"`
	Salon   string  `json:"nomSalon"`
	Color   string  `json:"couleur"`
	State   float64 `json:"etat"`
	Message string  `json:"message"`
}

type HandleMessage struct {
	State int `json:"etat"`
	Message string `json:"message"`
	Lounge string `json:"salon"`
}

type HandleRefresh struct {
	State int `json:"etat"`
	Message string `json:"message"`
	Private interface{} `json:"pvs"`
	Loung string `json:"salon"`
}

func Connection(login string, passwd string, salon string) (*HandleConnection, error) {

	jsonParsedResponse := &HandleConnection{}

	resp, err := http.PostForm(ajaxUrl, url.Values{
		"q":               {"conn"},
		"v":               {version},
		"identifiant":     {login},
		"motdepasse":      {passwd},
		"mode":            {"0"},
		"decalageHoraire": {"0"},
		"options":         {"M"},
		"salon":           {salon},
	})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&jsonParsedResponse)

	if err != nil {
		return nil, err
	}

	if jsonParsedResponse.State == 0 {
		return nil, errors.New(jsonParsedResponse.Message)
	}

	return jsonParsedResponse, nil
}

func ConnectionAuto(login string, passwd string) (*HandleConnection, error) {

	jsonParsedResponse := &HandleConnection{}

	resp, err := http.PostForm(ajaxUrl, url.Values{
		"q":               {"conn"},
		"v":               {version},
		"identifiant":     {login},
		"motdepasse":      {passwd},
		"mode":            {"2"},
		"decalageHoraire": {"0"},
		"options":         {"M"},
		"salon":           {"1"},
	})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&jsonParsedResponse)

	if err != nil {
		return nil, err
	}

	if jsonParsedResponse.State == 0 {
		return nil, errors.New(jsonParsedResponse.Message)
	}

	return jsonParsedResponse, nil
}



func SendMessage(session string, message string, totalMessages int) (*HandleMessage, error) {

	jsonParsedResponse := &HandleMessage{}

	resp, err := http.PostForm(ajaxUrl, url.Values{
		"q": {"cmd"},
		"v": {version},
		"s": {session},
		"c": {message},
		"a": {string(totalMessages)},
	})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&jsonParsedResponse)
	
	if err != nil {
		return nil, err
	}

	if jsonParsedResponse.State == 0 {
		return nil, errors.New(jsonParsedResponse.Message)
	}

	return jsonParsedResponse, nil
}


func Refresh(session string, totalMessages int) (*HandleRefresh, error) {

	jsonParsedResponse := &HandleRefresh{}

	resp, err := http.PostForm(ajaxUrl, url.Values{
		"q": {"act"},
		"v": {version},
		"s": {session},
		"a": {string(totalMessages)},
	})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&jsonParsedResponse)

	if err != nil {
		return nil, err
	}

	if jsonParsedResponse.State != 1 {
		return nil, errors.New(jsonParsedResponse.Message)
	}

	return jsonParsedResponse, nil
}

