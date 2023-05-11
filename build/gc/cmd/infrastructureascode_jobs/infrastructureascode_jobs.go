package infrastructureascode_jobs

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
	Description = utils.FormatUsageDescription("infrastructureascode_jobs", "SWAGGER_OVERRIDE_/api/v2/infrastructureascode/jobs", "SWAGGER_OVERRIDE_/api/v2/infrastructureascode/jobs", "SWAGGER_OVERRIDE_/api/v2/infrastructureascode/jobs", )
	infrastructureascode_jobsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("infrastructureascode_jobs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(infrastructureascode_jobsCmd)
}

func Cmdinfrastructureascode_jobs() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/infrastructureascode/jobs", utils.FormatPermissions([]string{ "infrastructureascode:job:add",  }), utils.GenerateDevCentreLink("POST", "Infrastructure as Code", "/api/v2/infrastructureascode/jobs")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AcceleratorInput"
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
        "$ref" : "#/components/schemas/InfrastructureascodeJob"
      }
    }
  }
}`)
	infrastructureascode_jobsCmd.AddCommand(createCmd)

	utils.AddFlag(getCmd.Flags(), "int", "maxResults", "1", "Number of jobs to show")
	utils.AddFlag(getCmd.Flags(), "bool", "includeErrors", "false", "Include error messages")
	utils.AddFlag(getCmd.Flags(), "string", "sortBy", "dateSubmitted", "Sort by Valid values: id, dateSubmitted, submittedBy, acceleratorId, status")
	utils.AddFlag(getCmd.Flags(), "string", "sortOrder", "desc", "Sort order Valid values: asc, desc")
	utils.AddFlag(getCmd.Flags(), "string", "acceleratorId", "", "Find only jobs associated with this accelerator")
	utils.AddFlag(getCmd.Flags(), "string", "submittedBy", "", "Find only jobs submitted by this user")
	utils.AddFlag(getCmd.Flags(), "string", "status", "", "Find only jobs in this state Valid values: Created, Queued, Running, Complete, Failed, Incomplete")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/infrastructureascode/jobs", utils.FormatPermissions([]string{ "infrastructureascode:job:view",  }), utils.GenerateDevCentreLink("GET", "Infrastructure as Code", "/api/v2/infrastructureascode/jobs")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/InfrastructureascodeJob"
      }
    }
  }
}`)
	infrastructureascode_jobsCmd.AddCommand(getCmd)

	utils.AddFlag(getjobCmd.Flags(), "bool", "details", "false", "Include details of execution, including job results or error information")
	getjobCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getjobCmd.UsageTemplate(), "GET", "/api/v2/infrastructureascode/jobs/{jobId}", utils.FormatPermissions([]string{ "infrastructureascode:job:view",  }), utils.GenerateDevCentreLink("GET", "Infrastructure as Code", "/api/v2/infrastructureascode/jobs/{jobId}")))
	utils.AddFileFlagIfUpsert(getjobCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getjobCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/InfrastructureascodeJob"
      }
    }
  }
}`)
	infrastructureascode_jobsCmd.AddCommand(getjobCmd)
	return infrastructureascode_jobsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Job",
	Long:  "Create a Job",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Acceleratorinput{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/infrastructureascode/jobs"

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
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get job history",
	Long:  "Get job history",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/infrastructureascode/jobs"

		maxResults := utils.GetFlag(cmd.Flags(), "int", "maxResults")
		if maxResults != "" {
			queryParams["maxResults"] = maxResults
		}
		includeErrors := utils.GetFlag(cmd.Flags(), "bool", "includeErrors")
		if includeErrors != "" {
			queryParams["includeErrors"] = includeErrors
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		acceleratorId := utils.GetFlag(cmd.Flags(), "string", "acceleratorId")
		if acceleratorId != "" {
			queryParams["acceleratorId"] = acceleratorId
		}
		submittedBy := utils.GetFlag(cmd.Flags(), "string", "submittedBy")
		if submittedBy != "" {
			queryParams["submittedBy"] = submittedBy
		}
		status := utils.GetFlag(cmd.Flags(), "string", "status")
		if status != "" {
			queryParams["status"] = status
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
var getjobCmd = &cobra.Command{
	Use:   "getjob [jobId]",
	Short: "Get job status and results",
	Long:  "Get job status and results",
	Args:  utils.DetermineArgs([]string{ "jobId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/infrastructureascode/jobs/{jobId}"
		jobId, args := args[0], args[1:]
		path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)

		details := utils.GetFlag(cmd.Flags(), "bool", "details")
		if details != "" {
			queryParams["details"] = details
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

		const opId = "getjob"
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
