package connection

import (
	"fmt"
	"time"
)

type CableConnection struct {
	data []byte
}

func (b CableConnection) Connect() {
	fmt.Println("Cable connected!")
}

func (b CableConnection) Disconnect() {
	fmt.Println("Cable disconnected.")
}

func (b CableConnection) TransferData() {
	fmt.Println("Transfering data...")

	for _, i := range b.data {
		fmt.Print(string(i))
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("Data transfered!")
}

func (b *CableConnection) SetData(data []byte) {
	b.data = data
}

func (b CableConnection) Data() []byte {
	return b.data
}
