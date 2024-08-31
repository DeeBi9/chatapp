package emiter

// Client structure to connect to the websocket connection
type EmitClient struct {
	ID                int
	Exchange          string
	sender_username   string
	reciever_username string
}
