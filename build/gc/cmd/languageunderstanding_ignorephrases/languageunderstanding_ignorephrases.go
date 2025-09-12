package languageunderstanding_ignorephrases

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
	Description = utils.FormatUsageDescription("languageunderstanding_ignorephrases", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/ignorephrases", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/ignorephrases", )
	languageunderstanding_ignorephrasesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding_ignorephrases"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstanding_ignorephrasesCmd)
}

func Cmdlanguageunderstanding_ignorephrases() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/languageunderstanding/ignorephrases/{languageCode}", utils.FormatPermissions([]string{ "languageUnderstanding:ignoredPhrase:add",  }), utils.GenerateDevCentreLink("POST", "Language Understanding", "/api/v2/languageunderstanding/ignorephrases/{languageCode}")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Request body containing phrases to be ignored",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/IgnorePhrasesRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/IgnorePhrasesResponse"
      }
    }
  }
}`)
	languageunderstanding_ignorephrasesCmd.AddCommand(createCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "The page size for the listing. The max that will be returned is 200.")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The page number for the listing")
	utils.AddFlag(listCmd.Flags(), "string", "text", "", "The phrase text filter applied to the listing")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "desc", "The sort order for the listing Valid values: asc, desc")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "dateModified", "The field to sort by for the listing Valid values: dateModified, text")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/languageunderstanding/ignorephrases/{languageCode}", utils.FormatPermissions([]string{ "languageUnderstanding:ignoredPhrase:view",  }), utils.GenerateDevCentreLink("GET", "Language Understanding", "/api/v2/languageunderstanding/ignorephrases/{languageCode}")))
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
	languageunderstanding_ignorephrasesCmd.AddCommand(listCmd)
	return languageunderstanding_ignorephrasesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [languageCode]",
	Short: "Add phrases to the ignored phrases list",
	Long:  "Add phrases to the ignored phrases list",
	Args:  utils.DetermineArgs([]string{ "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Ignorephrasesrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/ignorephrases/{languageCode}"
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
var listCmd = &cobra.Command{
	Use:   "list [languageCode]",
	Short: "Get list of all ignored phrases of the specified language code",
	Long:  "Get list of all ignored phrases of the specified language code",
	Args:  utils.DetermineArgs([]string{ "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/ignorephrases/{languageCode}"
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		text := utils.GetFlag(cmd.Flags(), "string", "text")
		if text != "" {
			queryParams["text"] = text
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
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
