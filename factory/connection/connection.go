package connection

type Connection interface {
	Connect()
	Disconnect()
	TransferData()
	Data() []byte
	SetData([]byte)
}
