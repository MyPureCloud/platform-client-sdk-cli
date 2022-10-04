package integrations_speech_tts_engines

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
	Description = utils.FormatUsageDescription("integrations_speech_tts_engines", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/tts/engines", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/tts/engines", )
	integrations_speech_tts_enginesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_tts_engines"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_tts_enginesCmd)
}

func Cmdintegrations_speech_tts_engines() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "bool", "includeVoices", "false", "Include voices for the engine")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/integrations/speech/tts/engines/{engineId}", utils.FormatPermissions([]string{ "integrations:integration:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/speech/tts/engines/{engineId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TtsEngineEntity"
      }
    }
  }
}`)
	integrations_speech_tts_enginesCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "bool", "includeVoices", "false", "Include voices for the engine")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Filter on engine name")
	utils.AddFlag(listCmd.Flags(), "string", "language", "", "Filter on supported language. If includeVoices=true then the voices are also filtered.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/integrations/speech/tts/engines", utils.FormatPermissions([]string{ "integrations:integration:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/speech/tts/engines")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	integrations_speech_tts_enginesCmd.AddCommand(listCmd)
	return integrations_speech_tts_enginesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [engineId]",
	Short: "Get details about a TTS engine",
	Long:  "Get details about a TTS engine",
	Args:  utils.DetermineArgs([]string{ "engineId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/speech/tts/engines/{engineId}"
		engineId, args := args[0], args[1:]
		path = strings.Replace(path, "{engineId}", fmt.Sprintf("%v", engineId), -1)

		includeVoices := utils.GetFlag(cmd.Flags(), "bool", "includeVoices")
		if includeVoices != "" {
			queryParams["includeVoices"] = includeVoices
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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of TTS engines enabled for org",
	Long:  "Get a list of TTS engines enabled for org",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/speech/tts/engines"

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		includeVoices := utils.GetFlag(cmd.Flags(), "bool", "includeVoices")
		if includeVoices != "" {
			queryParams["includeVoices"] = includeVoices
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		language := utils.GetFlag(cmd.Flags(), "string", "language")
		if language != "" {
			queryParams["language"] = language
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
