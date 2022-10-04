package quality_calibrations

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
	Description = utils.FormatUsageDescription("quality_calibrations", "SWAGGER_OVERRIDE_/api/v2/quality/calibrations", "SWAGGER_OVERRIDE_/api/v2/quality/calibrations", "SWAGGER_OVERRIDE_/api/v2/quality/calibrations", "SWAGGER_OVERRIDE_/api/v2/quality/calibrations", "SWAGGER_OVERRIDE_/api/v2/quality/calibrations", )
	quality_calibrationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_calibrations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_calibrationsCmd)
}

func Cmdquality_calibrations() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "string", "expand", "", "calibratorId")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/quality/calibrations", utils.FormatPermissions([]string{ "quality:calibration:add",  }), utils.GenerateDevCentreLink("POST", "Quality", "/api/v2/quality/calibrations")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "calibration",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CalibrationCreate"
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
        "$ref" : "#/components/schemas/Calibration"
      }
    }
  }
}`)
	quality_calibrationsCmd.AddCommand(createCmd)

	utils.AddFlag(deleteCmd.Flags(), "string", "calibratorId", "", "calibratorId - REQUIRED")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/quality/calibrations/{calibrationId}", utils.FormatPermissions([]string{ "quality:calibration:delete",  }), utils.GenerateDevCentreLink("DELETE", "Quality", "/api/v2/quality/calibrations/{calibrationId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	deleteCmd.MarkFlagRequired("calibratorId")
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Calibration"
      }
    }
  }
}`)
	quality_calibrationsCmd.AddCommand(deleteCmd)

	utils.AddFlag(getCmd.Flags(), "string", "calibratorId", "", "calibratorId")
	utils.AddFlag(getCmd.Flags(), "string", "conversationId", "", "conversationId")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/quality/calibrations/{calibrationId}", utils.FormatPermissions([]string{ "quality:calibration:view",  }), utils.GenerateDevCentreLink("GET", "Quality", "/api/v2/quality/calibrations/{calibrationId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Calibration"
      }
    }
  }
}`)
	quality_calibrationsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "The total page size requested")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The page number requested")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "variable name requested to sort by")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "variable name requested by expand list")
	utils.AddFlag(listCmd.Flags(), "string", "nextPage", "", "next page token")
	utils.AddFlag(listCmd.Flags(), "string", "previousPage", "", "Previous page token")
	utils.AddFlag(listCmd.Flags(), "string", "conversationId", "", "conversation id")
	utils.AddFlag(listCmd.Flags(), "time.Time", "startTime", "", "Beginning of the calibration query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z")
	utils.AddFlag(listCmd.Flags(), "time.Time", "endTime", "", "end of the calibration query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z")
	utils.AddFlag(listCmd.Flags(), "string", "calibratorId", "", "user id of calibrator - REQUIRED")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/quality/calibrations", utils.FormatPermissions([]string{ "quality:calibration:view",  }), utils.GenerateDevCentreLink("GET", "Quality", "/api/v2/quality/calibrations")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("calibratorId")
	
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
	quality_calibrationsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/quality/calibrations/{calibrationId}", utils.FormatPermissions([]string{ "quality:calibration:edit",  }), utils.GenerateDevCentreLink("PUT", "Quality", "/api/v2/quality/calibrations/{calibrationId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "Calibration",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Calibration"
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
        "$ref" : "#/components/schemas/Calibration"
      }
    }
  }
}`)
	quality_calibrationsCmd.AddCommand(updateCmd)
	return quality_calibrationsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a calibration",
	Long:  "Create a calibration",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Calibrationcreate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/calibrations"

		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
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
	Use:   "delete [calibrationId]",
	Short: "Delete a calibration by id.",
	Long:  "Delete a calibration by id.",
	Args:  utils.DetermineArgs([]string{ "calibrationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/calibrations/{calibrationId}"
		calibrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{calibrationId}", fmt.Sprintf("%v", calibrationId), -1)

		calibratorId := utils.GetFlag(cmd.Flags(), "string", "calibratorId")
		if calibratorId != "" {
			queryParams["calibratorId"] = calibratorId
		}
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
var getCmd = &cobra.Command{
	Use:   "get [calibrationId]",
	Short: "Get a calibration by id.  Requires either calibrator id or conversation id",
	Long:  "Get a calibration by id.  Requires either calibrator id or conversation id",
	Args:  utils.DetermineArgs([]string{ "calibrationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/calibrations/{calibrationId}"
		calibrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{calibrationId}", fmt.Sprintf("%v", calibrationId), -1)

		calibratorId := utils.GetFlag(cmd.Flags(), "string", "calibratorId")
		if calibratorId != "" {
			queryParams["calibratorId"] = calibratorId
		}
		conversationId := utils.GetFlag(cmd.Flags(), "string", "conversationId")
		if conversationId != "" {
			queryParams["conversationId"] = conversationId
		}
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
	Use:   "list",
	Short: "Get the list of calibrations",
	Long:  "Get the list of calibrations",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/calibrations"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		nextPage := utils.GetFlag(cmd.Flags(), "string", "nextPage")
		if nextPage != "" {
			queryParams["nextPage"] = nextPage
		}
		previousPage := utils.GetFlag(cmd.Flags(), "string", "previousPage")
		if previousPage != "" {
			queryParams["previousPage"] = previousPage
		}
		conversationId := utils.GetFlag(cmd.Flags(), "string", "conversationId")
		if conversationId != "" {
			queryParams["conversationId"] = conversationId
		}
		startTime := utils.GetFlag(cmd.Flags(), "time.Time", "startTime")
		if startTime != "" {
			queryParams["startTime"] = startTime
		}
		endTime := utils.GetFlag(cmd.Flags(), "time.Time", "endTime")
		if endTime != "" {
			queryParams["endTime"] = endTime
		}
		calibratorId := utils.GetFlag(cmd.Flags(), "string", "calibratorId")
		if calibratorId != "" {
			queryParams["calibratorId"] = calibratorId
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
var updateCmd = &cobra.Command{
	Use:   "update [calibrationId]",
	Short: "Update a calibration to the specified calibration via PUT.  Editable fields include: evaluators, expertEvaluator, and scoringIndex",
	Long:  "Update a calibration to the specified calibration via PUT.  Editable fields include: evaluators, expertEvaluator, and scoringIndex",
	Args:  utils.DetermineArgs([]string{ "calibrationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Calibration{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/calibrations/{calibrationId}"
		calibrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{calibrationId}", fmt.Sprintf("%v", calibrationId), -1)

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
