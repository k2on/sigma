package sigma

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {

	client, err := NewClient()
	if err != nil { panic(err) }

	// chats, err := client.Handles()
	// chats, err := client.Chats()
	chats, err := client.Messages(6, MessageFilter{Limit: 10})

	for _, c := range chats {
		fmt.Println(c)
	}
}
