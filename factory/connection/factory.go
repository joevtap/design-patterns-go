package connection

import "fmt"

func GetConnection(connectionType string) (Connection, error) {
	switch connectionType {
	case "cable":
		return newCableConnection(), nil
	case "bluetooth":
		return newBluetoothConnection(), nil
	default:
		return nil, fmt.Errorf("invalid connectionType: %v", connectionType)
	}
}
