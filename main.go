package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/wim07101993/fly_swatting_contest/participants"
)

const (
	Port     = "8080"
	IpAddres = ""
)

func readSettings(filePath string) Settings {
	log.Println("Reading settings at:", filePath)
	bsettings, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var settings Settings
	err = json.Unmarshal(bsettings, &settings)
	if err != nil {
		panic(err)
	}

	return settings
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Print("Request:Idex from", r.RemoteAddr)
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	var s Settings
	if len(os.Args) > 1 {
		s = readSettings(os.Args[1])
	} else {
		s = readSettings("settings.json")
	}

	log.Println("File set to:", s.ParticiPantsFilePath)

	ps := participants.NewService(s.ParticiPantsFilePath)
	pc := participants.NewController(ps)

	router := httprouter.New()

	router.POST("/participants", pc.Create)

	router.GET("/", Index)
	router.GET("/participants/", pc.GetAll)
	router.GET("/participants/:"+participants.NameParameter, pc.Get)

	router.PUT("/participants/:"+participants.NameParameter+"/score", pc.UpdateScore)
	router.PUT("/participants/:"+participants.NameParameter+"/name", pc.UpdateName)
	router.PUT("/participants/:"+participants.NameParameter+"/increaseScore", pc.IncreaseScore)
	router.PUT("/participants/:"+participants.NameParameter+"/decreaseScore", pc.DecreaseScore)

	router.DELETE("/participants/:"+participants.NameParameter, pc.Delete)

	log.Println("Start listening at", IpAddres+":"+Port)
	log.Fatal(http.ListenAndServe(IpAddres+":"+Port, router))
}
