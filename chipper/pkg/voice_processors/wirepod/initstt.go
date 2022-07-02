package wirepod

import (
	"fmt"
	"log"
	"os"
	"time"

	leopard "github.com/Picovoice/leopard/binding/go"
	rhino "github.com/Picovoice/rhino/binding/go/v2"
)

var leopardSTT1 leopard.Leopard
var leopardSTT2 leopard.Leopard
var leopardSTT3 leopard.Leopard
var leopardSTT4 leopard.Leopard
var leopardSTT5 leopard.Leopard
var rhinoSTI1 rhino.Rhino
var rhinoSTI2 rhino.Rhino
var rhinoSTI3 rhino.Rhino
var rhinoSTI4 rhino.Rhino
var rhinoSTI5 rhino.Rhino

func initPicovoice1(picovoiceKey string) {
	leopardSTT1 = leopard.Leopard{AccessKey: picovoiceKey}
	err := leopardSTT1.Init()
	rhinoSTI1 = rhino.NewRhino(picovoiceKey, "./intents.rhn")
	rhinoSTI1.Init()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Initialized Picovoice Instance 1")
}

func initPicovoice2(picovoiceKey string) {
	leopardSTT2 = leopard.Leopard{AccessKey: picovoiceKey}
	err := leopardSTT2.Init()
	if err != nil {
		log.Println(err)
	}
	rhinoSTI2 = rhino.NewRhino(picovoiceKey, "./intents.rhn")
	rhinoSTI2.Init()
	fmt.Println("Initialized Picovoice Instance 2")
}
func initPicovoice3(picovoiceKey string) {
	leopardSTT3 = leopard.Leopard{AccessKey: picovoiceKey}
	err := leopardSTT3.Init()
	if err != nil {
		log.Println(err)
	}
	rhinoSTI3 = rhino.NewRhino(picovoiceKey, "./intents.rhn")
	rhinoSTI3.Init()
	fmt.Println("Initialized Picovoice Instance 3")
}

func initPicovoice4(picovoiceKey string) {
	leopardSTT4 = leopard.Leopard{AccessKey: picovoiceKey}
	err := leopardSTT4.Init()
	if err != nil {
		log.Println(err)
	}
	rhinoSTI4 = rhino.NewRhino(picovoiceKey, "./intents.rhn")
	rhinoSTI4.Init()
	fmt.Println("Initialized Picovoice Instance 4")
}

func initPicovoice5(picovoiceKey string) {
	leopardSTT5 = leopard.Leopard{AccessKey: picovoiceKey}
	err := leopardSTT5.Init()
	if err != nil {
		log.Println(err)
	}
	rhinoSTI5 = rhino.NewRhino(picovoiceKey, "./intents.rhn")
	rhinoSTI5.Init()
	fmt.Println("Initialized Picovoice Instance 5")
}

func InitPicovoice() {
	var picovoiceKey string
	picovoiceKeyOS := os.Getenv("PICOVOICE_APIKEY")
	leopardKeyOS := os.Getenv("LEOPARD_APIKEY")
	if picovoiceKeyOS == "" {
		if leopardKeyOS == "" {
			fmt.Println("You must set PICOVOICE_APIKEY to a value.")
			os.Exit(1)
		} else {
			fmt.Println("PICOVOICE_APIKEY is not set, using LEOPARD_APIKEY")
			picovoiceKey = leopardKeyOS
		}
	} else {
		picovoiceKey = picovoiceKeyOS
	}
	fmt.Println("Initializing Picovoice Instances...")
	go initPicovoice1(picovoiceKey)
	go initPicovoice2(picovoiceKey)
	go initPicovoice3(picovoiceKey)
	go initPicovoice4(picovoiceKey)
	initPicovoice5(picovoiceKey)
	time.Sleep(time.Millisecond * 1500)
}
