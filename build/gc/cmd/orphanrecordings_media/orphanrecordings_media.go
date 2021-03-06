package orphanrecordings_media

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
	Description = utils.FormatUsageDescription("orphanrecordings_media", "SWAGGER_OVERRIDE_/api/v2/orphanrecordings/{orphanId}/media", )
	orphanrecordings_mediaCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orphanrecordings_media"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orphanrecordings_mediaCmd)
}

func Cmdorphanrecordings_media() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "formatId", "WEBM", "The desired media format. Valid values: WAV, WEBM, WAV_ULAW, OGG_VORBIS, OGG_OPUS, MP3, NONE")
	utils.AddFlag(getCmd.Flags(), "string", "emailFormatId", "EML", "The desired media format when downloading an email recording. Valid values: EML, NONE")
	utils.AddFlag(getCmd.Flags(), "string", "chatFormatId", "ZIP", "The desired media format when downloading a chat recording. Valid values: ZIP, NONE")
	utils.AddFlag(getCmd.Flags(), "string", "messageFormatId", "ZIP", "The desired media format when downloading a message recording. Valid values: ZIP, NONE")
	utils.AddFlag(getCmd.Flags(), "bool", "download", "false", "requesting a download format of the recording")
	utils.AddFlag(getCmd.Flags(), "string", "fileName", "", "the name of the downloaded fileName")
	utils.AddFlag(getCmd.Flags(), "string", "locale", "", "The locale for the requested file when downloading, as an ISO 639-1 code")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/orphanrecordings/{orphanId}/media", utils.FormatPermissions([]string{ "recording:orphan:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/orphanrecordings/{orphanId}/media")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Recording&quot;
  }
}`)
	orphanrecordings_mediaCmd.AddCommand(getCmd)
	
	return orphanrecordings_mediaCmd
}

var getCmd = &cobra.Command{
	Use:   "get [orphanId]",
	Short: "Gets the media of a single orphan recording",
	Long:  "Gets the media of a single orphan recording",
	Args:  utils.DetermineArgs([]string{ "orphanId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/orphanrecordings/{orphanId}/media"
		orphanId, args := args[0], args[1:]
		path = strings.Replace(path, "{orphanId}", fmt.Sprintf("%v", orphanId), -1)

		formatId := utils.GetFlag(cmd.Flags(), "string", "formatId")
		if formatId != "" {
			queryParams["formatId"] = formatId
		}
		emailFormatId := utils.GetFlag(cmd.Flags(), "string", "emailFormatId")
		if emailFormatId != "" {
			queryParams["emailFormatId"] = emailFormatId
		}
		chatFormatId := utils.GetFlag(cmd.Flags(), "string", "chatFormatId")
		if chatFormatId != "" {
			queryParams["chatFormatId"] = chatFormatId
		}
		messageFormatId := utils.GetFlag(cmd.Flags(), "string", "messageFormatId")
		if messageFormatId != "" {
			queryParams["messageFormatId"] = messageFormatId
		}
		download := utils.GetFlag(cmd.Flags(), "bool", "download")
		if download != "" {
			queryParams["download"] = download
		}
		fileName := utils.GetFlag(cmd.Flags(), "string", "fileName")
		if fileName != "" {
			queryParams["fileName"] = fileName
		}
		locale := utils.GetFlag(cmd.Flags(), "string", "locale")
		if locale != "" {
			queryParams["locale"] = locale
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
