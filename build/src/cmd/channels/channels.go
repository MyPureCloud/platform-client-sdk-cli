package channels
import (
	"fmt"
	"gc/retry"
	"gc/utils"
	"gc/services"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"strings"
)

var channelsCmd = &cobra.Command{
	Use:   "channels",
	Short: "Manages Genesys Cloud channels",
	Long:  `Manages Genesys Cloud channels`,
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(channelsCmd)
}

func Cmdchannels() *cobra.Command { 
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	channelsCmd.AddCommand(createCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "includechannels", "token", "Show user`s channels for this specific token or across all tokens for this user and app.  Channel Ids for other access tokens will not be shown, but will be presented to show their existence.")
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	channelsCmd.AddCommand(listCmd)
	
	return channelsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new channel",
	Long:  `Create a new channel`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/channels"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", "create", urlString)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "The list of existing channels",
	Long:  `The list of existing channels`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/channels"

		includechannels := utils.GetFlag(cmd.Flags(), "string", "includechannels")
		if includechannels != "" {
			queryParams["includechannels"] = includechannels
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", "list", urlString)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		utils.Render(results)
	},
}
