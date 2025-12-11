package conversations_recordingmetadata

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
	Description = utils.FormatUsageDescription("conversations_recordingmetadata", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/recordingmetadata", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/recordingmetadata", )
	conversations_recordingmetadataCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_recordingmetadata"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_recordingmetadataCmd)
}

func Cmdconversations_recordingmetadata() *cobra.Command { 
	conversationmetadataCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", conversationmetadataCmd.UsageTemplate(), "GET", "/api/v2/conversations/{conversationId}/recordingmetadata", utils.FormatPermissions([]string{ "recording:recording:view", "recording:recordingSegment:view", "recording:snippetRecording:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/conversations/{conversationId}/recordingmetadata")))
	utils.AddFileFlagIfUpsert(conversationmetadataCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(conversationmetadataCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "$ref" : "#/components/schemas/RecordingMetadata"
        }
      }
    }
  }
}`)
	conversations_recordingmetadataCmd.AddCommand(conversationmetadataCmd)

	recordingmetadataCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", recordingmetadataCmd.UsageTemplate(), "GET", "/api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}", utils.FormatPermissions([]string{ "recording:recording:view", "recording:recordingSegment:view", "recording:snippetRecording:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}")))
	utils.AddFileFlagIfUpsert(recordingmetadataCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(recordingmetadataCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/RecordingMetadata"
      }
    }
  }
}`)
	conversations_recordingmetadataCmd.AddCommand(recordingmetadataCmd)
	return conversations_recordingmetadataCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var conversationmetadataCmd = &cobra.Command{
	Use:   "conversationmetadata [conversationId]",
	Short: "Get recording metadata for a conversation. Does not return playable media nor system annotations. Bookmark annotations will be excluded if either recording:recording:view or recording:annotation:view permission is missing.",
	Long:  "Get recording metadata for a conversation. Does not return playable media nor system annotations. Bookmark annotations will be excluded if either recording:recording:view or recording:annotation:view permission is missing.",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/recordingmetadata"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

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

		const opId = "conversationmetadata"
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
var recordingmetadataCmd = &cobra.Command{
	Use:   "recordingmetadata [conversationId] [recordingId]",
	Short: "Get metadata for a specific recording. Does not return playable media.",
	Long:  "Get metadata for a specific recording. Does not return playable media.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "recordingId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		recordingId, args := args[0], args[1:]
		path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)

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

		const opId = "recordingmetadata"
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
