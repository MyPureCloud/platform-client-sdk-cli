package channels

import (
	"flag"
	"fmt"
	"gc/config"
	"gc/utils"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

func init() {
	channelsCmd.AddCommand(listenChannelCmd)
}

//processWSMessage is the function that processes a WebSocket channel.
func processWSMessage(c *websocket.Conn, done chan struct{}) {
	defer close(done)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		utils.Render(string(message))
	}
}

//waitForWSClose waits for an operating system interrupt to stop the websocket and cleanly close it down.
func waitForWSClose(c *websocket.Conn, done chan struct{}) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

var listenChannelCmd = &cobra.Command{
	Use:   "listen [channel id]",
	Short: "Listens to a channel and sends messages to standard out",
	Long:  `Listens to a channel and sends messages to standard out`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Root().Flags().GetString("profile")
		config, err := config.GetConfig(profileName)

		if err != nil {
			log.Fatal(err)
		}

		flag.Parse()
		log.SetFlags(0)

		//Set up the websocket connetion
		targetURI := fmt.Sprintf("wss://streaming.%s/channels/%s", config.Environment(), args[0])
		c, _, err := websocket.DefaultDialer.Dial(targetURI, nil)
		if err != nil {
			log.Fatal("Unable to connect to web socket:", err)
		}
		defer c.Close()
		done := make(chan struct{})

		//Spin off the function as an independently running go func
		go func() {
			processWSMessage(c, done)
		}()

		//Causes the CLI to wait for
		waitForWSClose(c, done)
	},
}
