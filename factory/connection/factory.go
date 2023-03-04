package connection

import "fmt"

func GetConnection(connectionType string) (Connection, error) {
	switch connectionType {
	case "cable":
		return new(CableConnection), nil
	case "bluetooth":
		return new(BluetoothConnection), nil
	default:
		return nil, fmt.Errorf("invalid connectionType: %v", connectionType)
	}
}
