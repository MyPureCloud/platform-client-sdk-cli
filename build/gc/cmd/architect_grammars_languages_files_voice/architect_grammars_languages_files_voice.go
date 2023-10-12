package architect_grammars_languages_files_voice

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
	Description = utils.FormatUsageDescription("architect_grammars_languages_files_voice", "SWAGGER_OVERRIDE_/api/v2/architect/grammars/{grammarId}/languages/{languageCode}/files/voice", "SWAGGER_OVERRIDE_/api/v2/architect/grammars/{grammarId}/languages/{languageCode}/files/voice", )
	architect_grammars_languages_files_voiceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("architect_grammars_languages_files_voice"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(architect_grammars_languages_files_voiceCmd)
}

func Cmdarchitect_grammars_languages_files_voice() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/architect/grammars/{grammarId}/languages/{languageCode}/files/voice", utils.FormatPermissions([]string{ "architect:grammar:edit",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/architect/grammars/{grammarId}/languages/{languageCode}/files/voice")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "query",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/GrammarFileUploadRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Presigned URL successfully created.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UploadUrlResponse"
      }
    }
  }
}`)
	architect_grammars_languages_files_voiceCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/architect/grammars/{grammarId}/languages/{languageCode}/files/voice", utils.FormatPermissions([]string{ "architect:grammar:edit",  }), utils.GenerateDevCentreLink("DELETE", "Architect", "/api/v2/architect/grammars/{grammarId}/languages/{languageCode}/files/voice")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "File was cleared.",
  "content" : { }
}`)
	architect_grammars_languages_files_voiceCmd.AddCommand(deleteCmd)
	return architect_grammars_languages_files_voiceCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [grammarId] [languageCode]",
	Short: "Creates a presigned URL for uploading a grammar voice mode file",
	Long:  "Creates a presigned URL for uploading a grammar voice mode file",
	Args:  utils.DetermineArgs([]string{ "grammarId", "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Grammarfileuploadrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/grammars/{grammarId}/languages/{languageCode}/files/voice"
		grammarId, args := args[0], args[1:]
		path = strings.Replace(path, "{grammarId}", fmt.Sprintf("%v", grammarId), -1)
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)

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

		const opId = "create"
		const httpMethod = "POST"
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
var deleteCmd = &cobra.Command{
	Use:   "delete [grammarId] [languageCode]",
	Short: "Clear the voice mode file for the grammar language if there is one",
	Long:  "Clear the voice mode file for the grammar language if there is one",
	Args:  utils.DetermineArgs([]string{ "grammarId", "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/grammars/{grammarId}/languages/{languageCode}/files/voice"
		grammarId, args := args[0], args[1:]
		path = strings.Replace(path, "{grammarId}", fmt.Sprintf("%v", grammarId), -1)
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)

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

		const opId = "delete"
		const httpMethod = "DELETE"
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
