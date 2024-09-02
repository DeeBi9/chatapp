package emiter

const Exchange = "DirectMessage"

// Client structure to connect to the websocket connection
type EmitClient struct {
	sender_username   string
	reciever_username string
}
