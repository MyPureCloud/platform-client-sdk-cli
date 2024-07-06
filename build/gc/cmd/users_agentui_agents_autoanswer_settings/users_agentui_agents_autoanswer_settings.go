package users_agentui_agents_autoanswer_settings

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
	Description = utils.FormatUsageDescription("users_agentui_agents_autoanswer_settings", "SWAGGER_OVERRIDE_/api/v2/users/agentui/agents/autoanswer/{agentId}/settings", "SWAGGER_OVERRIDE_/api/v2/users/agentui/agents/autoanswer/{agentId}/settings", "SWAGGER_OVERRIDE_/api/v2/users/agentui/agents/autoanswer/{agentId}/settings", "SWAGGER_OVERRIDE_/api/v2/users/agentui/agents/autoanswer/{agentId}/settings", )
	users_agentui_agents_autoanswer_settingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_agentui_agents_autoanswer_settings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_agentui_agents_autoanswer_settingsCmd)
}

func Cmdusers_agentui_agents_autoanswer_settings() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings", utils.FormatPermissions([]string{ "agentUI:agents:delete",  }), utils.GenerateDevCentreLink("DELETE", "Agent UI", "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Auto answer settings deleted successfully",
  "content" : { }
}`)
	users_agentui_agents_autoanswer_settingsCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings", utils.FormatPermissions([]string{ "agentUI:agents:view",  }), utils.GenerateDevCentreLink("GET", "Agent UI", "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "Auto answer settings retrieved successfully",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AutoAnswerSettings"
      }
    }
  }
}`)
	users_agentui_agents_autoanswer_settingsCmd.AddCommand(getCmd)

	patchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", patchCmd.UsageTemplate(), "PATCH", "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings", utils.FormatPermissions([]string{ "agentUI:agents:edit",  }), utils.GenerateDevCentreLink("PATCH", "Agent UI", "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings")))
	utils.AddFileFlagIfUpsert(patchCmd.Flags(), "PATCH", `{
  "description" : "AutoAnswerSettings",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AutoAnswerSettings"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(patchCmd.Flags(), "PATCH", `{
  "description" : "Auto answer settings updated successfully",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AutoAnswerSettings"
      }
    }
  }
}`)
	users_agentui_agents_autoanswer_settingsCmd.AddCommand(patchCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings", utils.FormatPermissions([]string{ "agentUI:agents:edit",  }), utils.GenerateDevCentreLink("PUT", "Agent UI", "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "AutoAnswerSettings",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AutoAnswerSettings"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "Auto answer settings set successfully",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AutoAnswerSettings"
      }
    }
  }
}`)
	users_agentui_agents_autoanswer_settingsCmd.AddCommand(updateCmd)
	return users_agentui_agents_autoanswer_settingsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [agentId]",
	Short: "Delete agent auto answer settings",
	Long:  "Delete agent auto answer settings",
	Args:  utils.DetermineArgs([]string{ "agentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings"
		agentId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)

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
	Use:   "get [agentId]",
	Short: "Get agent auto answer settings",
	Long:  "Get agent auto answer settings",
	Args:  utils.DetermineArgs([]string{ "agentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings"
		agentId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)

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
	Use:   "patch [agentId]",
	Short: "Update agent auto answer settings",
	Long:  "Update agent auto answer settings",
	Args:  utils.DetermineArgs([]string{ "agentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Autoanswersettings{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings"
		agentId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)

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
	Use:   "update [agentId]",
	Short: "Set agent auto answer settings",
	Long:  "Set agent auto answer settings",
	Args:  utils.DetermineArgs([]string{ "agentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Autoanswersettings{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/users/agentui/agents/autoanswer/{agentId}/settings"
		agentId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)

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
