package socialmedia_topics_dataingestionrules_twitter

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
	Description = utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_twitter", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter", )
	socialmedia_topics_dataingestionrules_twitterCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_twitter"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_topics_dataingestionrules_twitterCmd)
}

func Cmdsocialmedia_topics_dataingestionrules_twitter() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter", utils.FormatPermissions([]string{ "socialmedia:twitterDataIngestionRule:add",  }), utils.GenerateDevCentreLink("POST", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TwitterDataIngestionRuleRequest"
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
        "$ref" : "#/components/schemas/TwitterDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_twitterCmd.AddCommand(createCmd)

	utils.AddFlag(deleteCmd.Flags(), "bool", "hardDelete", "false", "Determines whether a X (formally Twitter) data ingestion rule should be soft-deleted (have it`s state set to deleted) or hard-deleted (permanently removed). Set to false (soft-delete) by default.")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:twitterDataIngestionRule:delete",  }), utils.GenerateDevCentreLink("DELETE", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "X (formally Twitter) data ingestion rule deleted.",
  "content" : { }
}`)
	socialmedia_topics_dataingestionrules_twitterCmd.AddCommand(deleteCmd)

	utils.AddFlag(getCmd.Flags(), "bool", "includeDeleted", "", "Determines whether to include soft-deleted items in the result.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:twitterDataIngestionRule:view",  }), utils.GenerateDevCentreLink("GET", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TwitterDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_twitterCmd.AddCommand(getCmd)

	patchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", patchCmd.UsageTemplate(), "PATCH", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:twitterDataIngestionRule:edit",  }), utils.GenerateDevCentreLink("PATCH", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}")))
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
        "$ref" : "#/components/schemas/TwitterDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_twitterCmd.AddCommand(patchCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:twitterDataIngestionRule:edit",  }), utils.GenerateDevCentreLink("PUT", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TwitterDataIngestionRuleRequest"
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
        "$ref" : "#/components/schemas/TwitterDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_twitterCmd.AddCommand(updateCmd)
	return socialmedia_topics_dataingestionrules_twitterCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [topicId]",
	Short: "Create an twitter data ingestion rule.",
	Long:  "Create an twitter data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Twitterdataingestionrulerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter"
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
	Use:   "delete [topicId] [twitterIngestionRuleId]",
	Short: "Delete a X (formally Twitter) data ingestion rule.",
	Long:  "Delete a X (formally Twitter) data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "twitterIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		twitterIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{twitterIngestionRuleId}", fmt.Sprintf("%v", twitterIngestionRuleId), -1)

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
	Use:   "get [topicId] [twitterIngestionRuleId]",
	Short: "Get a single X (formally Twitter) data ingestion rule.",
	Long:  "Get a single X (formally Twitter) data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "twitterIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		twitterIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{twitterIngestionRuleId}", fmt.Sprintf("%v", twitterIngestionRuleId), -1)

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
	Use:   "patch [topicId] [twitterIngestionRuleId]",
	Short: "Update the status of a X (formally Twitter) data ingestion rule.",
	Long:  "Update the status of a X (formally Twitter) data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "twitterIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Dataingestionrulestatuspatchrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		twitterIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{twitterIngestionRuleId}", fmt.Sprintf("%v", twitterIngestionRuleId), -1)

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
	Use:   "update [topicId] [twitterIngestionRuleId]",
	Short: "Update the X (formally Twitter) data ingestion rule.",
	Long:  "Update the X (formally Twitter) data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "twitterIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Twitterdataingestionrulerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		twitterIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{twitterIngestionRuleId}", fmt.Sprintf("%v", twitterIngestionRuleId), -1)

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
