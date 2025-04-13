package wscli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/wscli/clicmd"
)

type WsCli struct {
	userId string

	host string
	port int

	conn       *websocket.Conn
	outputChan chan string
	errChan    chan error
	doneChan   chan bool
}

func NewWsCli(userId, host string, port int) *WsCli {
	return &WsCli{
		userId:   userId,
		host:     host,
		port:     port,
		conn:     &websocket.Conn{},
		errChan:  make(chan error),
		doneChan: make(chan bool),
	}
}

func (cli *WsCli) Run() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	err := cli.connect()
	if err != nil {
		return err
	}
	defer cli.conn.Close()

	go cli.runUntilDone(cli.handleServerMessage)
	go cli.runUntilDone(cli.handleInput)

	err = <-cli.errChan
	close(cli.doneChan)
	return err
}

func (cli *WsCli) connect() error {
	url := url.URL{
		Scheme:   "ws",
		Host:     fmt.Sprintf("%s:%d", cli.host, cli.port),
		Path:     "/ws",
		RawQuery: url.Values{"userId": []string{cli.userId}}.Encode(),
	}
	log.Println("Connecting to server at", url.String())

	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		return err
	}
	cli.conn = conn
	return nil
}

func (cli *WsCli) runUntilDone(function func() error) error {
	for {
		select {
		case <-cli.doneChan:
			return nil
		default:
			err := function()
			if err != nil {
				cli.errChan <- err
				return err
			}
		}
	}
}

func (cli *WsCli) handleServerMessage() error {
	_, msg, err := cli.conn.ReadMessage()
	if err != nil {
		return err
	}

	var prettyMsg bytes.Buffer
	err = json.Indent(&prettyMsg, msg, "", "  ")
	if err != nil {
		log.Println(msg)
		return nil
	}

	log.Printf("Received message:\n%s\n", string(prettyMsg.Bytes()))
	return nil
}

func (cli *WsCli) handleInput() error {
	printPrompt()
	cmd, err := clicmd.ReadCommand()
	if err != nil {
		log.Println("Error reading command:", err)
		return nil
	}

	wsMsg, err := cmd.ToWsMsg()
	if err != nil {
		log.Println("Error reading command:", err)
		return nil
	}

	err = cli.conn.WriteMessage(websocket.TextMessage, []byte(wsMsg))
	return nil
}

func printPrompt() {
	log.Println(" -- Available commands:")
	for _, cmd := range clicmd.AvailableParsers {
		log.Println(" * ", cmd.GetFormat())
	}
	log.Println(" -- Enter command: ")
}
