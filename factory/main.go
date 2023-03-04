package main

import (
	"fmt"
	"os"

	"github.com/joevtap/design-patterns-go/factory/connection"
)

func main() {
	bluetooth, _ := connection.GetConnection("bluetooth")
	bluetooth.SetData([]byte("Some data"))
	bluetooth.Connect()
	bluetooth.TransferData()
	bluetooth.Disconnect()

	fmt.Println()

	cable, _ := connection.GetConnection("cable")
	cable.SetData([]byte("More data"))
	cable.Connect()
	cable.TransferData()
	cable.Disconnect()

	fmt.Println()

	connection5G, err := connection.GetConnection("5G")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	connection5G.Connect()
}
