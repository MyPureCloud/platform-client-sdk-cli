package routing_queues_members

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
	Description = utils.FormatUsageDescription("routing_queues_members", "SWAGGER_OVERRIDE_/api/v2/routing/queues/{queueId}/members", "SWAGGER_OVERRIDE_/api/v2/routing/queues/{queueId}/members", "SWAGGER_OVERRIDE_/api/v2/routing/queues/{queueId}/members", "SWAGGER_OVERRIDE_/api/v2/routing/queues/{queueId}/members", "SWAGGER_OVERRIDE_/api/v2/routing/queues/{queueId}/members", )
	routing_queues_membersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_queues_members"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_queues_membersCmd)
}

func Cmdrouting_queues_members() *cobra.Command { 
	activateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", activateCmd.UsageTemplate(), "PATCH", "/api/v2/routing/queues/{queueId}/members", utils.FormatPermissions([]string{ "routing:queue:edit", "routing:queueMember:manage",  }), utils.GenerateDevCentreLink("PATCH", "Routing", "/api/v2/routing/queues/{queueId}/members")))
	utils.AddFileFlagIfUpsert(activateCmd.Flags(), "PATCH", `{
  "description" : "Queue Members",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "$ref" : "#/components/schemas/QueueMember"
        }
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(activateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/QueueMemberEntityListing"
      }
    }
  }
}`)
	routing_queues_membersCmd.AddCommand(activateCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/routing/queues/{queueId}/members/{memberId}", utils.FormatPermissions([]string{ "routing:queue:edit", "routing:queueMember:manage",  }), utils.GenerateDevCentreLink("DELETE", "Routing", "/api/v2/routing/queues/{queueId}/members/{memberId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Deleted.",
  "content" : { }
}`)
	routing_queues_membersCmd.AddCommand(deleteCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Max value is 100")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "asc", "Note: results are sorted by name. Valid values: asc, desc")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: routingStatus, presence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, dateLastLogin, authorization.unusedRoles, team, profileSkills, certifications, locations, groups, skills, languages, languagePreference, employerInfo, biography")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Filter by queue member name")
	utils.AddFlag(listCmd.Flags(), "[]string", "profileSkills", "", "Filter by profile skill")
	utils.AddFlag(listCmd.Flags(), "[]string", "skills", "", "Filter by skill")
	utils.AddFlag(listCmd.Flags(), "[]string", "languages", "", "Filter by language")
	utils.AddFlag(listCmd.Flags(), "[]string", "routingStatus", "", "Filter by routing status")
	utils.AddFlag(listCmd.Flags(), "[]string", "presence", "", "Filter by presence")
	utils.AddFlag(listCmd.Flags(), "string", "memberBy", "", "Filter by member type Valid values: user, group")
	utils.AddFlag(listCmd.Flags(), "bool", "joined", "", "Filter by joined status Valid values: true, false")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/routing/queues/{queueId}/members", utils.FormatPermissions([]string{ "routing:queue:view", "routing:queue:edit", "routing:queue:readonly", "routing:queueMember:manage",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/queues/{queueId}/members")))
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
	routing_queues_membersCmd.AddCommand(listCmd)

	utils.AddFlag(moveCmd.Flags(), "bool", "delete", "false", "True to delete queue members")
	moveCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", moveCmd.UsageTemplate(), "POST", "/api/v2/routing/queues/{queueId}/members", utils.FormatPermissions([]string{ "routing:queue:edit", "routing:queueMember:manage",  }), utils.GenerateDevCentreLink("POST", "Routing", "/api/v2/routing/queues/{queueId}/members")))
	utils.AddFileFlagIfUpsert(moveCmd.Flags(), "POST", `{
  "description" : "Queue Members",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "$ref" : "#/components/schemas/WritableEntity"
        }
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(moveCmd.Flags(), "POST", `{
  "description" : "The request could not be understood by the server due to malformed syntax.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ErrorBody"
      }
    }
  },
  "x-inin-error-codes" : {
    "bad.request" : "Invalid request data.  Make sure you submit a valid number of queue members.",
    "response.entity.too.large" : "The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable",
    "invalid.date" : "Dates must be specified as ISO-8601 strings. For example: yyyy-MM-ddTHH:mm:ss.SSSZ",
    "invalid.query.param.value" : "Value [%s] is not valid for parameter [%s]. Allowable values are: %s",
    "client.failed.request" : "The client did not produce a request with valid end of stream signaling. This can be caused by poor network connection and/or client behavior.",
    "invalid.property" : "Value [%s] is not a valid property for object [%s]",
    "queue.size.limit" : "Adding all requested members would exceed queue member limit.",
    "constraint.validation" : "%s",
    "invalid.value" : "Value [%s] is not valid for field type [%s]. Allowable values are: %s"
  }
}`)
	routing_queues_membersCmd.AddCommand(moveCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/routing/queues/{queueId}/members/{memberId}", utils.FormatPermissions([]string{ "routing:queue:edit", "routing:queueMember:manage",  }), utils.GenerateDevCentreLink("PATCH", "Routing", "/api/v2/routing/queues/{queueId}/members/{memberId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Queue Member",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/QueueMember"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "User update has been accepted",
  "content" : { }
}`)
	routing_queues_membersCmd.AddCommand(updateCmd)
	return routing_queues_membersCmd
}

var activateCmd = &cobra.Command{
	Use:   "activate [queueId]",
	Short: "Join or unjoin a set of users for a queue",
	Long:  "Join or unjoin a set of users for a queue",
	Args:  utils.DetermineArgs([]string{ "queueId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Queuemember{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/members"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "activate"
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
var deleteCmd = &cobra.Command{
	Use:   "delete [queueId] [memberId]",
	Short: "Delete a queue member.",
	Long:  "Delete a queue member.",
	Args:  utils.DetermineArgs([]string{ "queueId", "memberId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/members/{memberId}"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)
		memberId, args := args[0], args[1:]
		path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)

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
	Use:   "list [queueId]",
	Short: "Get the members of this queue.",
	Long:  "Get the members of this queue.",
	Args:  utils.DetermineArgs([]string{ "queueId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/members"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		profileSkills := utils.GetFlag(cmd.Flags(), "[]string", "profileSkills")
		if profileSkills != "" {
			queryParams["profileSkills"] = profileSkills
		}
		skills := utils.GetFlag(cmd.Flags(), "[]string", "skills")
		if skills != "" {
			queryParams["skills"] = skills
		}
		languages := utils.GetFlag(cmd.Flags(), "[]string", "languages")
		if languages != "" {
			queryParams["languages"] = languages
		}
		routingStatus := utils.GetFlag(cmd.Flags(), "[]string", "routingStatus")
		if routingStatus != "" {
			queryParams["routingStatus"] = routingStatus
		}
		presence := utils.GetFlag(cmd.Flags(), "[]string", "presence")
		if presence != "" {
			queryParams["presence"] = presence
		}
		memberBy := utils.GetFlag(cmd.Flags(), "string", "memberBy")
		if memberBy != "" {
			queryParams["memberBy"] = memberBy
		}
		joined := utils.GetFlag(cmd.Flags(), "bool", "joined")
		if joined != "" {
			queryParams["joined"] = joined
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
var moveCmd = &cobra.Command{
	Use:   "move [queueId]",
	Short: "Bulk add or delete up to 100 queue members",
	Long:  "Bulk add or delete up to 100 queue members",
	Args:  utils.DetermineArgs([]string{ "queueId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Writableentity{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/members"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)

		delete := utils.GetFlag(cmd.Flags(), "bool", "delete")
		if delete != "" {
			queryParams["delete"] = delete
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "move"
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
var updateCmd = &cobra.Command{
	Use:   "update [queueId] [memberId]",
	Short: "Update the ring number OR joined status for a queue member.",
	Long:  "Update the ring number OR joined status for a queue member.",
	Args:  utils.DetermineArgs([]string{ "queueId", "memberId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Queuemember{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/members/{memberId}"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)
		memberId, args := args[0], args[1:]
		path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
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
