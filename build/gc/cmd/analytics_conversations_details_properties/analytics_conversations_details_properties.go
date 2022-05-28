package analytics_conversations_details_properties

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
	Description = utils.FormatUsageDescription("analytics_conversations_details_properties", "SWAGGER_OVERRIDE_/api/v2/analytics/conversations/{conversationId}/details/properties", )
	analytics_conversations_details_propertiesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_conversations_details_properties"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_conversations_details_propertiesCmd)
}

func Cmdanalytics_conversations_details_properties() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/analytics/conversations/{conversationId}/details/properties", utils.FormatPermissions([]string{ "analytics:conversationProperties:index",  }), utils.GenerateDevCentreLink("POST", "Analytics", "/api/v2/analytics/conversations/{conversationId}/details/properties")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "request",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PropertyIndexRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Accepted - Indexing properties",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PropertyIndexRequest"
      }
    }
  }
}`)
	analytics_conversations_details_propertiesCmd.AddCommand(createCmd)
	return analytics_conversations_details_propertiesCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId]",
	Short: "Index conversation properties",
	Long:  "Index conversation properties",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Propertyindexrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/conversations/{conversationId}/details/properties"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

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

		utils.Render(results)
	},
}
