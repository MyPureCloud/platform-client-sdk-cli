package conversations_recordings

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
	Description = utils.FormatUsageDescription("conversations_recordings", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/recordings", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/recordings", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/recordings", )
	conversations_recordingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_recordings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_recordingsCmd)
}

func Cmdconversations_recordings() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "formatId", "WEBM", "The desired media format. Valid values:WAV,WEBM,WAV_ULAW,OGG_VORBIS,OGG_OPUS,MP3,NONE Valid values: WAV, WEBM, WAV_ULAW, OGG_VORBIS, OGG_OPUS, MP3, NONE")
	utils.AddFlag(getCmd.Flags(), "string", "emailFormatId", "EML", "The desired media format when downloading an email recording. Valid values:EML,NONE Valid values: EML, NONE")
	utils.AddFlag(getCmd.Flags(), "string", "chatFormatId", "ZIP", "The desired media format when downloading a chat recording. Valid values:ZIP,NONE  Valid values: ZIP, NONE")
	utils.AddFlag(getCmd.Flags(), "string", "messageFormatId", "ZIP", "The desired media format when downloading a message recording. Valid values:ZIP,NONE Valid values: ZIP, NONE")
	utils.AddFlag(getCmd.Flags(), "bool", "download", "false", "requesting a download format of the recording. Valid values:true,false Valid values: true, false")
	utils.AddFlag(getCmd.Flags(), "string", "fileName", "", "the name of the downloaded fileName")
	utils.AddFlag(getCmd.Flags(), "string", "locale", "", "The locale for the requested file when downloading, as an ISO 639-1 code")
	utils.AddFlag(getCmd.Flags(), "[]string", "mediaFormats", "", "All acceptable media formats. Overrides formatId. Valid values:WAV,WEBM,WAV_ULAW,OGG_VORBIS,OGG_OPUS,MP3")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/conversations/{conversationId}/recordings/{recordingId}", utils.FormatPermissions([]string{ "recording:recording:view", "recording:recordingSegment:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/conversations/{conversationId}/recordings/{recordingId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "Success - recording is transcoding",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Recording"
      }
    }
  }
}`)
	conversations_recordingsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "maxWaitMs", "5000", "The maximum number of milliseconds to wait for the recording to be ready. Must be a positive value.")
	utils.AddFlag(listCmd.Flags(), "string", "formatId", "WEBM", "The desired media format. Valid values:WAV,WEBM,WAV_ULAW,OGG_VORBIS,OGG_OPUS,MP3,NONE. Valid values: WAV, WEBM, WAV_ULAW, OGG_VORBIS, OGG_OPUS, MP3, NONE")
	utils.AddFlag(listCmd.Flags(), "[]string", "mediaFormats", "", "All acceptable media formats. Overrides formatId. Valid values:WAV,WEBM,WAV_ULAW,OGG_VORBIS,OGG_OPUS,MP3.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/conversations/{conversationId}/recordings", utils.FormatPermissions([]string{ "recording:recording:view", "recording:recordingSegment:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/conversations/{conversationId}/recordings")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "$ref" : "#/components/schemas/Recording"
        }
      }
    }
  }
}`)
	conversations_recordingsCmd.AddCommand(listCmd)

	utils.AddFlag(updateCmd.Flags(), "bool", "clearExport", "", "Whether to clear the pending export for the recording")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/conversations/{conversationId}/recordings/{recordingId}", utils.FormatPermissions([]string{ "recording:recording:view", "recording:recording:editRetention", "recording:screenRecording:editRetention",  }), utils.GenerateDevCentreLink("PUT", "Recording", "/api/v2/conversations/{conversationId}/recordings/{recordingId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "recording",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Recording"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Recording"
      }
    }
  }
}`)
	conversations_recordingsCmd.AddCommand(updateCmd)
	return conversations_recordingsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [conversationId] [recordingId]",
	Short: "Gets a specific recording.",
	Long:  "Gets a specific recording.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "recordingId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/recordings/{recordingId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		recordingId, args := args[0], args[1:]
		path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)

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
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
var listCmd = &cobra.Command{
	Use:   "list [conversationId]",
	Short: "Get all of a Conversation`s Recordings.",
	Long:  "Get all of a Conversation`s Recordings.",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/recordings"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

		maxWaitMs := utils.GetFlag(cmd.Flags(), "int", "maxWaitMs")
		if maxWaitMs != "" {
			queryParams["maxWaitMs"] = maxWaitMs
		}
		formatId := utils.GetFlag(cmd.Flags(), "string", "formatId")
		if formatId != "" {
			queryParams["formatId"] = formatId
		}
		mediaFormats := utils.GetFlag(cmd.Flags(), "[]string", "mediaFormats")
		if mediaFormats != "" {
			queryParams["mediaFormats"] = mediaFormats
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "list"
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
var updateCmd = &cobra.Command{
	Use:   "update [conversationId] [recordingId]",
	Short: "Updates the retention records on a recording.",
	Long:  "Updates the retention records on a recording.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "recordingId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Recording{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/recordings/{recordingId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		recordingId, args := args[0], args[1:]
		path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)

		clearExport := utils.GetFlag(cmd.Flags(), "bool", "clearExport")
		if clearExport != "" {
			queryParams["clearExport"] = clearExport
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "update"
		const httpMethod = "PUT"
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
