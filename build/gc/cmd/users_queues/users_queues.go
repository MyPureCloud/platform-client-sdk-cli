package users_queues

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
	Description = utils.FormatUsageDescription("users_queues", "SWAGGER_OVERRIDE_/api/v2/users/{userId}/queues", "SWAGGER_OVERRIDE_/api/v2/users/{userId}/queues", "SWAGGER_OVERRIDE_/api/v2/users/{userId}/queues", )
	users_queuesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_queues"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_queuesCmd)
}

func Cmdusers_queues() *cobra.Command { 
	joinCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", joinCmd.UsageTemplate(), "PATCH", "/api/v2/users/{userId}/queues/{queueId}", utils.FormatPermissions([]string{ "routing:queue:join", "routing:queueMember:manage",  }), utils.GenerateDevCentreLink("PATCH", "Routing", "/api/v2/users/{userId}/queues/{queueId}")))
	utils.AddFileFlagIfUpsert(joinCmd.Flags(), "PATCH", `{
  "description" : "Queue Member",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UserQueue"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(joinCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UserQueue"
      }
    }
  }
}`)
	users_queuesCmd.AddCommand(joinCmd)

	utils.AddFlag(joinsetCmd.Flags(), "[]string", "divisionId", "", "Division ID(s)")
	joinsetCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", joinsetCmd.UsageTemplate(), "PATCH", "/api/v2/users/{userId}/queues", utils.FormatPermissions([]string{ "routing:queue:join", "routing:queueMember:manage",  }), utils.GenerateDevCentreLink("PATCH", "Routing", "/api/v2/users/{userId}/queues")))
	utils.AddFileFlagIfUpsert(joinsetCmd.Flags(), "PATCH", `{
  "description" : "User Queues",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "$ref" : "#/components/schemas/UserQueue"
        }
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(joinsetCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UserQueueEntityListing"
      }
    }
  }
}`)
	users_queuesCmd.AddCommand(joinsetCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "bool", "joined", "true", "Is joined to the queue")
	utils.AddFlag(listCmd.Flags(), "[]string", "divisionId", "", "Division ID(s)")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/users/{userId}/queues", utils.FormatPermissions([]string{ "routing:queue:view", "routing:queue:join", "routing:queueMember:manage",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/users/{userId}/queues")))
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
	users_queuesCmd.AddCommand(listCmd)
	return users_queuesCmd
}

var joinCmd = &cobra.Command{
	Use:   "join [queueId] [userId]",
	Short: "Join or unjoin a queue for a user",
	Long:  "Join or unjoin a queue for a user",
	Args:  utils.DetermineArgs([]string{ "queueId", "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Userqueue{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/queues/{queueId}"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "join"
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
var joinsetCmd = &cobra.Command{
	Use:   "joinset [userId]",
	Short: "Join or unjoin a set of queues for a user",
	Long:  "Join or unjoin a set of queues for a user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Userqueue{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/queues"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "joinset"
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
var listCmd = &cobra.Command{
	Use:   "list [userId]",
	Short: "Get queues for user",
	Long:  "Get queues for user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/queues"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		joined := utils.GetFlag(cmd.Flags(), "bool", "joined")
		if joined != "" {
			queryParams["joined"] = joined
		}
		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
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
