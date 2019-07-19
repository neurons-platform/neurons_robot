package client_interface


type Sender interface {
	// Send(msg string) bool
	Send(map[string]string) bool
}

