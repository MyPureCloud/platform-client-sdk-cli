package webmessaging_deployments_pushdevices

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
	Description = utils.FormatUsageDescription("webmessaging_deployments_pushdevices", "SWAGGER_OVERRIDE_/api/v2/webmessaging/deployments/{deploymentId}/pushdevices", "SWAGGER_OVERRIDE_/api/v2/webmessaging/deployments/{deploymentId}/pushdevices", "SWAGGER_OVERRIDE_/api/v2/webmessaging/deployments/{deploymentId}/pushdevices", )
	webmessaging_deployments_pushdevicesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webmessaging_deployments_pushdevices"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webmessaging_deployments_pushdevicesCmd)
}

func Cmdwebmessaging_deployments_pushdevices() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "WebMessaging", "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Request body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PushDeviceInsertRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Device created successfully",
  "content" : { }
}`)
	webmessaging_deployments_pushdevicesCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "WebMessaging", "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Device deleted successfully",
  "content" : { }
}`)
	webmessaging_deployments_pushdevicesCmd.AddCommand(deleteCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PATCH", "WebMessaging", "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Request body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PushDeviceUpdateRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "Device updated successfully",
  "content" : { }
}`)
	webmessaging_deployments_pushdevicesCmd.AddCommand(updateCmd)
	return webmessaging_deployments_pushdevicesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [deploymentId] [tokenId]",
	Short: "Add a new device information",
	Long:  "Add a new device information",
	Args:  utils.DetermineArgs([]string{ "deploymentId", "tokenId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Pushdeviceinsertrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}"
		deploymentId, args := args[0], args[1:]
		path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentId), -1)
		tokenId, args := args[0], args[1:]
		path = strings.Replace(path, "{tokenId}", fmt.Sprintf("%v", tokenId), -1)

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
	Use:   "delete [deploymentId] [tokenId]",
	Short: "Delete device information",
	Long:  "Delete device information",
	Args:  utils.DetermineArgs([]string{ "deploymentId", "tokenId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}"
		deploymentId, args := args[0], args[1:]
		path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentId), -1)
		tokenId, args := args[0], args[1:]
		path = strings.Replace(path, "{tokenId}", fmt.Sprintf("%v", tokenId), -1)

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
var updateCmd = &cobra.Command{
	Use:   "update [deploymentId] [tokenId]",
	Short: "Edit device information",
	Long:  "Edit device information",
	Args:  utils.DetermineArgs([]string{ "deploymentId", "tokenId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Pushdeviceupdaterequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}"
		deploymentId, args := args[0], args[1:]
		path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentId), -1)
		tokenId, args := args[0], args[1:]
		path = strings.Replace(path, "{tokenId}", fmt.Sprintf("%v", tokenId), -1)

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
