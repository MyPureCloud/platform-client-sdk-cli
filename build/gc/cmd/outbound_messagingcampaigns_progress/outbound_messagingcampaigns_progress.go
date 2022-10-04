package outbound_messagingcampaigns_progress

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
	Description = utils.FormatUsageDescription("outbound_messagingcampaigns_progress", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns/progress", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/progress", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/progress", )
	outbound_messagingcampaigns_progressCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_messagingcampaigns_progress"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_messagingcampaigns_progressCmd)
}

func Cmdoutbound_messagingcampaigns_progress() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/messagingcampaigns/progress", utils.FormatPermissions([]string{ "outbound:messagingCampaign:view", "outbound:emailCampaign:view",  }), utils.GenerateDevCentreLink("POST", "Outbound", "/api/v2/outbound/messagingcampaigns/progress")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Messaging Campaign IDs",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "type" : "string"
        }
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
        "type" : "array",
        "items" : {
          "$ref" : "#/components/schemas/CampaignProgress"
        }
      }
    }
  }
}`)
	outbound_messagingcampaigns_progressCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/progress", utils.FormatPermissions([]string{ "outbound:messagingCampaign:edit", "outbound:emailCampaign:edit",  }), utils.GenerateDevCentreLink("DELETE", "Outbound", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/progress")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Accepted - the messaging campaign will be recycled momentarily",
  "content" : { }
}`)
	outbound_messagingcampaigns_progressCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/progress", utils.FormatPermissions([]string{ "outbound:messagingCampaign:view", "outbound:emailCampaign:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/progress")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CampaignProgress"
      }
    }
  }
}`)
	outbound_messagingcampaigns_progressCmd.AddCommand(getCmd)
	return outbound_messagingcampaigns_progressCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Get progress for a list of messaging campaigns",
	Long:  "Get progress for a list of messaging campaigns",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.String{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns/progress"

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
	Use:   "delete [messagingCampaignId]",
	Short: "Reset messaging campaign progress and recycle the messaging campaign",
	Long:  "Reset messaging campaign progress and recycle the messaging campaign",
	Args:  utils.DetermineArgs([]string{ "messagingCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/progress"
		messagingCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{messagingCampaignId}", fmt.Sprintf("%v", messagingCampaignId), -1)

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
	Use:   "get [messagingCampaignId]",
	Short: "Get messaging campaign`s progress",
	Long:  "Get messaging campaign`s progress",
	Args:  utils.DetermineArgs([]string{ "messagingCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/progress"
		messagingCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{messagingCampaignId}", fmt.Sprintf("%v", messagingCampaignId), -1)

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
