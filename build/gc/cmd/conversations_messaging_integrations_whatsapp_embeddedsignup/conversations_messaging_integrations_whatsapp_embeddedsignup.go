package conversations_messaging_integrations_whatsapp_embeddedsignup

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
	Description = utils.FormatUsageDescription("conversations_messaging_integrations_whatsapp_embeddedsignup", "SWAGGER_OVERRIDE_/api/v2/conversations/messaging/integrations/whatsapp/embeddedsignup", "SWAGGER_OVERRIDE_/api/v2/conversations/messaging/integrations/whatsapp/embeddedsignup", )
	conversations_messaging_integrations_whatsapp_embeddedsignupCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messaging_integrations_whatsapp_embeddedsignup"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messaging_integrations_whatsapp_embeddedsignupCmd)
}

func Cmdconversations_messaging_integrations_whatsapp_embeddedsignup() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/messaging/integrations/whatsapp/embeddedsignup", utils.FormatPermissions([]string{ "messaging:whatsappIntegration:add",  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/messaging/integrations/whatsapp/embeddedsignup")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "WhatsAppEmbeddedSignupIntegrationRequest",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WhatsAppEmbeddedSignupIntegrationRequest"
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
        "$ref" : "#/components/schemas/WhatsAppIntegration"
      }
    }
  }
}`)
	conversations_messaging_integrations_whatsapp_embeddedsignupCmd.AddCommand(createCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/conversations/messaging/integrations/whatsapp/embeddedsignup/{integrationId}", utils.FormatPermissions([]string{ "messaging:integration:edit",  }), utils.GenerateDevCentreLink("PATCH", "Conversations", "/api/v2/conversations/messaging/integrations/whatsapp/embeddedsignup/{integrationId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "WhatsAppEmbeddedSignupIntegrationActivationRequest",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WhatsAppEmbeddedSignupIntegrationActivationRequest"
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
        "$ref" : "#/components/schemas/WhatsAppIntegration"
      }
    }
  }
}`)
	conversations_messaging_integrations_whatsapp_embeddedsignupCmd.AddCommand(updateCmd)
	return conversations_messaging_integrations_whatsapp_embeddedsignupCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a WhatsApp Integration using the WhatsApp embedded signup flow",
	Long:  "Create a WhatsApp Integration using the WhatsApp embedded signup flow",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Whatsappembeddedsignupintegrationrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/messaging/integrations/whatsapp/embeddedsignup"

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
var updateCmd = &cobra.Command{
	Use:   "update [integrationId]",
	Short: "Activate a WhatsApp messaging integration created using the WhatsApp embedded signup flow",
	Long:  "Activate a WhatsApp messaging integration created using the WhatsApp embedded signup flow",
	Args:  utils.DetermineArgs([]string{ "integrationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Whatsappembeddedsignupintegrationactivationrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/messaging/integrations/whatsapp/embeddedsignup/{integrationId}"
		integrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{integrationId}", fmt.Sprintf("%v", integrationId), -1)

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
