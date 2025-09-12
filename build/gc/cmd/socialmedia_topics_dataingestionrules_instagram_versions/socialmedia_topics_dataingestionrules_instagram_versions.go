package socialmedia_topics_dataingestionrules_instagram_versions

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
	Description = utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_instagram_versions", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions", )
	socialmedia_topics_dataingestionrules_instagram_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_instagram_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_topics_dataingestionrules_instagram_versionsCmd)
}

func Cmdsocialmedia_topics_dataingestionrules_instagram_versions() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "bool", "includeDeleted", "", "Determines whether to include soft-deleted item in the result.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions/{dataIngestionRuleVersion}", utils.FormatPermissions([]string{ "socialmedia:instagramDataIngestionRule:view",  }), utils.GenerateDevCentreLink("GET", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions/{dataIngestionRuleVersion}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "Successful operation.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/InstagramDataIngestionRuleVersionResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_instagram_versionsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "bool", "includeDeleted", "", "Determines whether to include soft-deleted items in the result.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions", utils.FormatPermissions([]string{ "socialmedia:instagramDataIngestionRule:view",  }), utils.GenerateDevCentreLink("GET", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "Successful operation.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_instagram_versionsCmd.AddCommand(listCmd)
	return socialmedia_topics_dataingestionrules_instagram_versionsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [topicId] [instagramIngestionRuleId] [dataIngestionRuleVersion]",
	Short: "Get a single Instagram data ingestion rule version.",
	Long:  "Get a single Instagram data ingestion rule version.",
	Args:  utils.DetermineArgs([]string{ "topicId", "instagramIngestionRuleId", "dataIngestionRuleVersion", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions/{dataIngestionRuleVersion}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		instagramIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{instagramIngestionRuleId}", fmt.Sprintf("%v", instagramIngestionRuleId), -1)
		dataIngestionRuleVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{dataIngestionRuleVersion}", fmt.Sprintf("%v", dataIngestionRuleVersion), -1)

		includeDeleted := utils.GetFlag(cmd.Flags(), "bool", "includeDeleted")
		if includeDeleted != "" {
			queryParams["includeDeleted"] = includeDeleted
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
	Use:   "list [topicId] [instagramIngestionRuleId]",
	Short: "Get the Instagram data ingestion rule versions.",
	Long:  "Get the Instagram data ingestion rule versions.",
	Args:  utils.DetermineArgs([]string{ "topicId", "instagramIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		instagramIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{instagramIngestionRuleId}", fmt.Sprintf("%v", instagramIngestionRuleId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		includeDeleted := utils.GetFlag(cmd.Flags(), "bool", "includeDeleted")
		if includeDeleted != "" {
			queryParams["includeDeleted"] = includeDeleted
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
