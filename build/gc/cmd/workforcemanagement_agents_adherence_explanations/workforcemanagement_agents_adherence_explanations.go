package workforcemanagement_agents_adherence_explanations

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
	Description = utils.FormatUsageDescription("workforcemanagement_agents_adherence_explanations", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations", )
	workforcemanagement_agents_adherence_explanationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_adherence_explanations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_adherence_explanationsCmd)
}

func Cmdworkforcemanagement_agents_adherence_explanations() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations", utils.FormatPermissions([]string{ "wfm:adherenceExplanation:add",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "The request body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AddAdherenceExplanationAdminRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "The adherence explanation is processing",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AdherenceExplanationAsyncResponse"
      }
    }
  }
}`)
	workforcemanagement_agents_adherence_explanationsCmd.AddCommand(createCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}", utils.FormatPermissions([]string{ "wfm:adherenceExplanation:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AdherenceExplanationResponse"
      }
    }
  }
}`)
	workforcemanagement_agents_adherence_explanationsCmd.AddCommand(getCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}", utils.FormatPermissions([]string{ "wfm:adherenceExplanation:edit",  }), utils.GenerateDevCentreLink("PATCH", "Workforce Management", "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "The request body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UpdateAdherenceExplanationStatusRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AdherenceExplanationAsyncResponse"
      }
    }
  }
}`)
	workforcemanagement_agents_adherence_explanationsCmd.AddCommand(updateCmd)
	return workforcemanagement_agents_adherence_explanationsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [agentId]",
	Short: "Add an adherence explanation for the requested user",
	Long:  "Add an adherence explanation for the requested user",
	Args:  utils.DetermineArgs([]string{ "agentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Addadherenceexplanationadminrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations"
		agentId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
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
var getCmd = &cobra.Command{
	Use:   "get [agentId] [explanationId]",
	Short: "Get an adherence explanation",
	Long:  "Get an adherence explanation",
	Args:  utils.DetermineArgs([]string{ "agentId", "explanationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}"
		agentId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)
		explanationId, args := args[0], args[1:]
		path = strings.Replace(path, "{explanationId}", fmt.Sprintf("%v", explanationId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
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
var updateCmd = &cobra.Command{
	Use:   "update [agentId] [explanationId]",
	Short: "Update an adherence explanation",
	Long:  "Update an adherence explanation",
	Args:  utils.DetermineArgs([]string{ "agentId", "explanationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Updateadherenceexplanationstatusrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}"
		agentId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)
		explanationId, args := args[0], args[1:]
		path = strings.Replace(path, "{explanationId}", fmt.Sprintf("%v", explanationId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
