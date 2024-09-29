package speechandtextanalytics_translations_languages_conversations

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
	Description = utils.FormatUsageDescription("speechandtextanalytics_translations_languages_conversations", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/translations/languages/{languageId}/conversations", )
	speechandtextanalytics_translations_languages_conversationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_translations_languages_conversations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_translations_languages_conversationsCmd)
}

func Cmdspeechandtextanalytics_translations_languages_conversations() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "communicationId", "", "Communication id associated with the conversation")
	utils.AddFlag(getCmd.Flags(), "string", "recordingId", "", "Recording id associated with the communication")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/speechandtextanalytics/translations/languages/{languageId}/conversations/{conversationId}", utils.FormatPermissions([]string{ "speechAndTextAnalytics:translation:view",  }), utils.GenerateDevCentreLink("GET", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/translations/languages/{languageId}/conversations/{conversationId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CommunicationTranslationList"
      }
    }
  }
}`)
	speechandtextanalytics_translations_languages_conversationsCmd.AddCommand(getCmd)
	return speechandtextanalytics_translations_languages_conversationsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [languageId] [conversationId]",
	Short: "Translate all communication(s) for an interaction.",
	Long:  "Translate all communication(s) for an interaction.",
	Args:  utils.DetermineArgs([]string{ "languageId", "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/translations/languages/{languageId}/conversations/{conversationId}"
		languageId, args := args[0], args[1:]
		path = strings.Replace(path, "{languageId}", fmt.Sprintf("%v", languageId), -1)
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

		communicationId := utils.GetFlag(cmd.Flags(), "string", "communicationId")
		if communicationId != "" {
			queryParams["communicationId"] = communicationId
		}
		recordingId := utils.GetFlag(cmd.Flags(), "string", "recordingId")
		if recordingId != "" {
			queryParams["recordingId"] = recordingId
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
