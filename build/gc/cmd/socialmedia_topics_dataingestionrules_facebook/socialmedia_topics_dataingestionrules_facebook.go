package socialmedia_topics_dataingestionrules_facebook

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
	Description = utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_facebook", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook", )
	socialmedia_topics_dataingestionrules_facebookCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_facebook"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_topics_dataingestionrules_facebookCmd)
}

func Cmdsocialmedia_topics_dataingestionrules_facebook() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook", utils.FormatPermissions([]string{ "socialmedia:facebookDataIngestionRule:add",  }), utils.GenerateDevCentreLink("POST", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/FacebookDataIngestionRuleRequest"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Request accepted - the data ingestion rule will be created momentarily.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/FacebookDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_facebookCmd.AddCommand(createCmd)

	utils.AddFlag(deleteCmd.Flags(), "bool", "hardDelete", "false", "Determines whether a Facebook data ingestion rule should be soft-deleted (have it`s state set to deleted) or hard-deleted (permanently removed). Set to false (soft-delete) by default.")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:facebookDataIngestionRule:delete",  }), utils.GenerateDevCentreLink("DELETE", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Facebook data ingestion rule deleted.",
  "content" : { }
}`)
	socialmedia_topics_dataingestionrules_facebookCmd.AddCommand(deleteCmd)

	utils.AddFlag(getCmd.Flags(), "bool", "includeDeleted", "", "Determines whether to include soft-deleted items in the result.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:facebookDataIngestionRule:view",  }), utils.GenerateDevCentreLink("GET", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/FacebookDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_facebookCmd.AddCommand(getCmd)

	patchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", patchCmd.UsageTemplate(), "PATCH", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:facebookDataIngestionRule:edit",  }), utils.GenerateDevCentreLink("PATCH", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(patchCmd.Flags(), "PATCH", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/DataIngestionRuleStatusPatchRequest"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(patchCmd.Flags(), "PATCH", `{
  "description" : "Request accepted - status will be updated momentarily.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/FacebookDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_facebookCmd.AddCommand(patchCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:facebookDataIngestionRule:edit",  }), utils.GenerateDevCentreLink("PUT", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/FacebookDataIngestionRuleRequest"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "Request accepted - the data ingestion rule will be updated momentarily.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/FacebookDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_facebookCmd.AddCommand(updateCmd)
	return socialmedia_topics_dataingestionrules_facebookCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [topicId]",
	Short: "Create an Facebook data ingestion rule.",
	Long:  "Create an Facebook data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Facebookdataingestionrulerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)

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
	Use:   "delete [topicId] [facebookIngestionRuleId]",
	Short: "Delete a Facebook data ingestion rule.",
	Long:  "Delete a Facebook data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "facebookIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		facebookIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{facebookIngestionRuleId}", fmt.Sprintf("%v", facebookIngestionRuleId), -1)

		hardDelete := utils.GetFlag(cmd.Flags(), "bool", "hardDelete")
		if hardDelete != "" {
			queryParams["hardDelete"] = hardDelete
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
var getCmd = &cobra.Command{
	Use:   "get [topicId] [facebookIngestionRuleId]",
	Short: "Get a single Facebook data ingestion rule.",
	Long:  "Get a single Facebook data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "facebookIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		facebookIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{facebookIngestionRuleId}", fmt.Sprintf("%v", facebookIngestionRuleId), -1)

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
var patchCmd = &cobra.Command{
	Use:   "patch [topicId] [facebookIngestionRuleId]",
	Short: "Update the status of a Facebook data ingestion rule.",
	Long:  "Update the status of a Facebook data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "facebookIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Dataingestionrulestatuspatchrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		facebookIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{facebookIngestionRuleId}", fmt.Sprintf("%v", facebookIngestionRuleId), -1)

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

		const opId = "patch"
		const httpMethod = "PATCH"
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
	Use:   "update [topicId] [facebookIngestionRuleId]",
	Short: "Update the Facebook data ingestion rule.",
	Long:  "Update the Facebook data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "facebookIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Facebookdataingestionrulerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		facebookIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{facebookIngestionRuleId}", fmt.Sprintf("%v", facebookIngestionRuleId), -1)

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
