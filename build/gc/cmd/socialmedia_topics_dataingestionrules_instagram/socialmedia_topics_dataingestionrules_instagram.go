package socialmedia_topics_dataingestionrules_instagram

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
	Description = utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_instagram", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram", )
	socialmedia_topics_dataingestionrules_instagramCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_instagram"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_topics_dataingestionrules_instagramCmd)
}

func Cmdsocialmedia_topics_dataingestionrules_instagram() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram", utils.FormatPermissions([]string{ "socialmedia:instagramDataIngestionRule:add",  }), utils.GenerateDevCentreLink("POST", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/InstagramDataIngestionRuleRequest"
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
        "$ref" : "#/components/schemas/InstagramDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_instagramCmd.AddCommand(createCmd)

	utils.AddFlag(deleteCmd.Flags(), "bool", "hardDelete", "false", "Determines whether a Instagram data ingestion rule should be soft-deleted (have it`s state set to deleted) or hard-deleted (permanently removed). Set to false (soft-delete) by default.")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:instagramDataIngestionRule:delete",  }), utils.GenerateDevCentreLink("DELETE", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Instagram data ingestion rule deleted.",
  "content" : { }
}`)
	socialmedia_topics_dataingestionrules_instagramCmd.AddCommand(deleteCmd)

	utils.AddFlag(getCmd.Flags(), "bool", "includeDeleted", "", "Determines whether to include soft-deleted items in the result.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:instagramDataIngestionRule:view",  }), utils.GenerateDevCentreLink("GET", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "Successful operation.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/InstagramDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_instagramCmd.AddCommand(getCmd)

	patchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", patchCmd.UsageTemplate(), "PATCH", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:instagramDataIngestionRule:edit",  }), utils.GenerateDevCentreLink("PATCH", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}")))
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
        "$ref" : "#/components/schemas/InstagramDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_instagramCmd.AddCommand(patchCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}", utils.FormatPermissions([]string{ "socialmedia:instagramDataIngestionRule:edit",  }), utils.GenerateDevCentreLink("PUT", "Social Media", "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/InstagramDataIngestionRuleRequest"
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
        "$ref" : "#/components/schemas/InstagramDataIngestionRuleResponse"
      }
    }
  }
}`)
	socialmedia_topics_dataingestionrules_instagramCmd.AddCommand(updateCmd)
	return socialmedia_topics_dataingestionrules_instagramCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [topicId]",
	Short: "Create an Instagram data ingestion rule.",
	Long:  "Create an Instagram data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Instagramdataingestionrulerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram"
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
	Use:   "delete [topicId] [instagramIngestionRuleId]",
	Short: "Delete a Instagram data ingestion rule.",
	Long:  "Delete a Instagram data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "instagramIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		instagramIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{instagramIngestionRuleId}", fmt.Sprintf("%v", instagramIngestionRuleId), -1)

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
	Use:   "get [topicId] [instagramIngestionRuleId]",
	Short: "Get a single Instagram data ingestion rule.",
	Long:  "Get a single Instagram data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "instagramIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		instagramIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{instagramIngestionRuleId}", fmt.Sprintf("%v", instagramIngestionRuleId), -1)

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
	Use:   "patch [topicId] [instagramIngestionRuleId]",
	Short: "Update the status of a Instagram data ingestion rule.",
	Long:  "Update the status of a Instagram data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "instagramIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Dataingestionrulestatuspatchrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		instagramIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{instagramIngestionRuleId}", fmt.Sprintf("%v", instagramIngestionRuleId), -1)

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
	Use:   "update [topicId] [instagramIngestionRuleId]",
	Short: "Update the Instagram data ingestion rule.",
	Long:  "Update the Instagram data ingestion rule.",
	Args:  utils.DetermineArgs([]string{ "topicId", "instagramIngestionRuleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Instagramdataingestionrulerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)
		instagramIngestionRuleId, args := args[0], args[1:]
		path = strings.Replace(path, "{instagramIngestionRuleId}", fmt.Sprintf("%v", instagramIngestionRuleId), -1)

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
