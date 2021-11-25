package userrecordings_media

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("userrecordings_media", "SWAGGER_OVERRIDE_/api/v2/userrecordings/{recordingId}/media", )
	userrecordings_mediaCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("userrecordings_media"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(userrecordings_mediaCmd)
}

func Cmduserrecordings_media() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "formatId", "WEBM", "The desired media format. Valid values: WAV, WEBM, WAV_ULAW, OGG_VORBIS, OGG_OPUS, MP3, NONE")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/userrecordings/{recordingId}/media", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "User Recordings", "/api/v2/userrecordings/{recordingId}/media")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DownloadResponse"
  }
}`)
	userrecordings_mediaCmd.AddCommand(getCmd)
	
	return userrecordings_mediaCmd
}

var getCmd = &cobra.Command{
	Use:   "get [recordingId]",
	Short: "Download a user recording.",
	Long:  "Download a user recording.",
	Args:  utils.DetermineArgs([]string{ "recordingId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/userrecordings/{recordingId}/media"
		recordingId, args := args[0], args[1:]
		path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)

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

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
