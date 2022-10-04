package orphanrecordings_media

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
	utils.AddFlag(getCmd.Flags(), "bool", "download", "false", "requesting a download format of the recording Valid values: true, false")
	utils.AddFlag(getCmd.Flags(), "string", "fileName", "", "the name of the downloaded fileName")
	utils.AddFlag(getCmd.Flags(), "string", "locale", "", "The locale for the requested file when downloading, as an ISO 639-1 code")
	utils.AddFlag(getCmd.Flags(), "[]string", "mediaFormats", "", "All acceptable media formats. Overrides formatId. Valid values:WAV,WEBM,WAV_ULAW,OGG_VORBIS,OGG_OPUS,MP3")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/orphanrecordings/{orphanId}/media", utils.FormatPermissions([]string{ "recording:orphan:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/orphanrecordings/{orphanId}/media")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Recording"
      }
    }
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
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

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
		mediaFormats := utils.GetFlag(cmd.Flags(), "[]string", "mediaFormats")
		if mediaFormats != "" {
			queryParams["mediaFormats"] = mediaFormats
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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
