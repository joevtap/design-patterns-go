package connection

import (
	"fmt"
	"time"
)

type BluetoothConnection struct {
	data []byte
}

func newBluetoothConnection() *BluetoothConnection {
	return new(BluetoothConnection)
}

func (b BluetoothConnection) Connect() {
	fmt.Println("Bluetooth connected!")
}

func (b BluetoothConnection) Disconnect() {
	fmt.Println("Bluetooth disconnected.")
}

func (b BluetoothConnection) TransferData() {
	fmt.Println("Transfering data...")

	for _, i := range b.data {
		fmt.Print(string(i))
		time.Sleep(1 * time.Second)
	}

	fmt.Println()
	fmt.Println("Data transfered!")
}

func (b *BluetoothConnection) SetData(data []byte) {
	b.data = data
}

func (b BluetoothConnection) Data() []byte {
	return b.data
}
