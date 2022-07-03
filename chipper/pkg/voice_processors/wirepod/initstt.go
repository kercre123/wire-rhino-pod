package wirepod

import (
	"fmt"
	"os"
	"strconv"

	rhino "github.com/Picovoice/rhino/binding/go/v2"
)

var rhinoSTIArray []rhino.Rhino
var picovoiceInstancesOS string = os.Getenv("PICOVOICE_INSTANCES")
var picovoiceInstances int

func InitPicovoice() {
	var picovoiceKey string
	picovoiceKeyOS := os.Getenv("PICOVOICE_APIKEY")
	leopardKeyOS := os.Getenv("LEOPARD_APIKEY")
	if picovoiceInstancesOS == "" {
		fmt.Println("PICOVOICE_INSTANCES is not set, using default value of 5")
		picovoiceInstances = 5
	} else {
		picovoiceInstances, err := strconv.Atoi(picovoiceInstancesOS)
		if err != nil {
			fmt.Println("PICOVOICE_INSTANCES is not a valid integer, using default value of 5")
			picovoiceInstances = 5
		}
		fmt.Println("Initializing " + strconv.Itoa(picovoiceInstances) + " Picovoice Instances...")
	}
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
	for i := 0; i < picovoiceInstances; i++ {
		fmt.Println("Initializing Picovoice Instance " + strconv.Itoa(i))
		rhinoSTIArray = append(rhinoSTIArray, rhino.NewRhino(picovoiceKey, "./intents.rhn"))
		rhinoSTIArray[i].Init()
	}
}
