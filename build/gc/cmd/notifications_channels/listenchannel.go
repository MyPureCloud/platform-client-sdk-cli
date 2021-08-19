package notifications_channels

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

func init() {
	notifications_channelsCmd.AddCommand(listenChannelCmd)
}

//processWSMessage is the function that processes a WebSocket channel.
func processWSMessage(c *websocket.Conn, done chan struct{}, heartbeatSuppressed bool) {
	defer close(done)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			logger.Warn("Websocket read:", err)
			return
		}
		if isWebSocketHeartbeat(string(message)) && heartbeatSuppressed {
			continue
		}
		utils.Render(string(message))
	}
}

func isWebSocketHeartbeat(message string) bool {
	return strings.Contains(message, "WebSocket Heartbeat") && strings.Contains(message, "channel.metadata")
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
				logger.Warn("Websocket write:", err)
				return
			}
		case <-interrupt:
			logger.Warn("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logger.Warn("Websocket write close:", err)
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
	Use:   "listen [channelId]",
	Short: "Listens to a channel and sends messages to standard out",
	Long:  `Listens to a channel and sends messages to standard out`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Root().Flags().GetString("profile")
		heartbeatSuppressed, _ := cmd.Root().Flags().GetBool("noheartbeat")

		config, err := config.GetConfig(profileName)
		if err != nil {
			logger.Fatal(err)
		}

		//Set up the websocket connection
		targetURI := fmt.Sprintf("wss://streaming.%s/channels/%s", config.Environment(), args[0])
		c, _, err := websocket.DefaultDialer.Dial(targetURI, nil)
		if err != nil {
			logger.Fatal("Unable to connect to web socket:", err)
		}
		defer c.Close()
		done := make(chan struct{})

		//Spin off the function as an independently running go func
		go func() {
			processWSMessage(c, done, heartbeatSuppressed)
		}()

		//Causes the CLI to wait for
		waitForWSClose(c, done)
	},
}
