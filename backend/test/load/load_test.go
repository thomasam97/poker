package load

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/gorilla/websocket"
)

var done chan interface{}

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
	}
}

var msg = NewChooseMessage("8")
var payload, _ = json.Marshal(msg)

func _TestWaitForAll(t *testing.T) {
	chan1 := make(chan struct{}, 1)
	chan2 := make(chan struct{}, 1)
	channels := make([]chan struct{}, 0)
	channels = append(channels, chan1, chan2)
	go func() {
		// time.Sleep(time.Second * 2)
		chan1 <- struct{}{}
		chan2 <- struct{}{}
	}()
	<-waitForAll(channels)
}

func TestLoad(t *testing.T) {
	const noOfUsers = 50
	connect := make(chan struct{}, 1)
	send := make(chan struct{}, 1)
	finishedChannels := make([]chan struct{}, 0)
	waitingChannels := make([]chan struct{}, 0)

	for ui := 0; ui < noOfUsers; ui++ {
		waiting, finished := user(ui, connect, send)
		finishedChannels = append(finishedChannels, finished)
		waitingChannels = append(waitingChannels, waiting)
	}

	<-waitForAll(waitingChannels)
	for ui := 0; ui < noOfUsers; ui++ {
		connect <- struct{}{}
	}
	// time.Sleep(time.Second * 2)
	for ui := 0; ui < noOfUsers; ui++ {
		send <- struct{}{}
	}
	<-waitForAll(finishedChannels)

}

func user(ui int, connect, send chan struct{}) (waiting, finished chan struct{}) {
	finished = make(chan struct{}, 1)
	waiting = make(chan struct{}, 1)
	go func(ui int, connect, send, finished, waiting chan struct{}) {
		// socketURL := fmt.Sprintf("ws://localhost:7788/poker/api/room/test-room/user-%d", ui)
		socketURL := fmt.Sprintf("wss://matrix.sprinteins.com/poker/api/room/test-room/user-%d", ui)
		waiting <- struct{}{}
		<-connect
		conn, _, err := websocket.DefaultDialer.Dial(socketURL, nil)
		if err != nil {
			log.Fatal("Error connecting to Websocket Server:", err)
		}
		// defer conn.Close()

		<-send
		err = conn.WriteMessage(websocket.TextMessage, []byte(payload))
		if err != nil {
			log.Println("Error during writing to websocket:", err)
			return
		}
		finished <- struct{}{}
	}(ui, connect, send, finished, waiting)

	return waiting, finished
}

func NewChooseMessage(card string) Message {
	return Message{
		Type:    "Choose",
		Payload: card,
	}
}

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func waitForAll(channels []chan struct{}) chan struct{} {
	out := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(len(channels))

	for ci, channel := range channels {
		go func(channel chan struct{}, ci int) {
			_ = <-channel
			wg.Done()
		}(channel, ci)
	}

	go func() {
		wg.Wait()
		out <- struct{}{}
		close(out)
	}()
	return out
}
