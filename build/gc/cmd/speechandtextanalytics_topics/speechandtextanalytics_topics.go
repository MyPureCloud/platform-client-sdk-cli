package speechandtextanalytics_topics

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
	Description = utils.FormatUsageDescription("speechandtextanalytics_topics", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/topics", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/topics", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/topics", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/topics", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/topics", )
	speechandtextanalytics_topicsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_topics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_topicsCmd)
}

func Cmdspeechandtextanalytics_topics() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/speechandtextanalytics/topics", utils.FormatPermissions([]string{ "speechAndTextAnalytics:topic:add",  }), utils.GenerateDevCentreLink("POST", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/topics")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "The topic to create",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TopicRequest"
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
        "$ref" : "#/components/schemas/Topic"
      }
    }
  }
}`)
	speechandtextanalytics_topicsCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/speechandtextanalytics/topics/{topicId}", utils.FormatPermissions([]string{ "speechAndTextAnalytics:topic:delete",  }), utils.GenerateDevCentreLink("DELETE", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/topics/{topicId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The topic was deleted successfully",
  "content" : { }
}`)
	speechandtextanalytics_topicsCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/speechandtextanalytics/topics/{topicId}", utils.FormatPermissions([]string{ "speechAndTextAnalytics:topic:view",  }), utils.GenerateDevCentreLink("GET", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/topics/{topicId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Topic"
      }
    }
  }
}`)
	speechandtextanalytics_topicsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "nextPage", "", "The key for listing the next page")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "20", "The page size for the listing")
	utils.AddFlag(listCmd.Flags(), "string", "state", "", "Topic state. Defaults to latest Valid values: latest, published")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Case insensitive partial name to filter by")
	utils.AddFlag(listCmd.Flags(), "[]string", "ids", "", "Comma separated Topic IDs to filter by. Cannot be used with other filters. Maximum of 50 IDs allowed.")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "Sort results by. Defaults to name Valid values: name")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "", "Sort order. Defaults to asc Valid values: asc, desc")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/speechandtextanalytics/topics", utils.FormatPermissions([]string{ "speechAndTextAnalytics:topic:view",  }), utils.GenerateDevCentreLink("GET", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/topics")))
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
	speechandtextanalytics_topicsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/speechandtextanalytics/topics/{topicId}", utils.FormatPermissions([]string{ "speechAndTextAnalytics:topic:edit",  }), utils.GenerateDevCentreLink("PUT", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/topics/{topicId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "The topic to update",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TopicRequest"
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
        "$ref" : "#/components/schemas/Topic"
      }
    }
  }
}`)
	speechandtextanalytics_topicsCmd.AddCommand(updateCmd)
	return speechandtextanalytics_topicsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Speech and Text Analytics topic",
	Long:  "Create new Speech and Text Analytics topic",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Topicrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/topics"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [topicId]",
	Short: "Delete a Speech and Text Analytics topic by id",
	Long:  "Delete a Speech and Text Analytics topic by id",
	Args:  utils.DetermineArgs([]string{ "topicId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/topics/{topicId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [topicId]",
	Short: "Get a Speech and Text Analytics topic by id",
	Long:  "Get a Speech and Text Analytics topic by id",
	Args:  utils.DetermineArgs([]string{ "topicId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/topics/{topicId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of Speech and Text Analytics topics",
	Long:  "Get the list of Speech and Text Analytics topics",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/topics"

		nextPage := utils.GetFlag(cmd.Flags(), "string", "nextPage")
		if nextPage != "" {
			queryParams["nextPage"] = nextPage
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		state := utils.GetFlag(cmd.Flags(), "string", "state")
		if state != "" {
			queryParams["state"] = state
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		ids := utils.GetFlag(cmd.Flags(), "[]string", "ids")
		if ids != "" {
			queryParams["ids"] = ids
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [topicId]",
	Short: "Update existing Speech and Text Analytics topic",
	Long:  "Update existing Speech and Text Analytics topic",
	Args:  utils.DetermineArgs([]string{ "topicId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Topicrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/topics/{topicId}"
		topicId, args := args[0], args[1:]
		path = strings.Replace(path, "{topicId}", fmt.Sprintf("%v", topicId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
