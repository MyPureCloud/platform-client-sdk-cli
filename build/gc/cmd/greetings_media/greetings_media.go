package greetings_media

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("greetings_media", "SWAGGER_OVERRIDE_/api/v2/greetings/{greetingId}/media", )
	greetings_mediaCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("greetings_media"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(greetings_mediaCmd)
}

func Cmdgreetings_media() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "formatId", "WAV", "The desired media format. Valid values: WAV, WEBM, WAV_ULAW, OGG_VORBIS, OGG_OPUS, MP3, NONE")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/greetings/{greetingId}/media", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/GreetingMediaInfo&quot;
  }
}`)
	greetings_mediaCmd.AddCommand(getCmd)
	
	return greetings_mediaCmd
}

var getCmd = &cobra.Command{
	Use:   "get [greetingId]",
	Short: "Get media playback URI for this greeting",
	Long:  "Get media playback URI for this greeting",
	Args:  utils.DetermineArgs([]string{ "greetingId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/greetings/{greetingId}/media"
		greetingId, args := args[0], args[1:]
		path = strings.Replace(path, "{greetingId}", fmt.Sprintf("%v", greetingId), -1)

		formatId := utils.GetFlag(cmd.Flags(), "string", "formatId")
		if formatId != "" {
			queryParams["formatId"] = formatId
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
