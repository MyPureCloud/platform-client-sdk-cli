package recordings_deletionprotection

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
	Description = utils.FormatUsageDescription("recordings_deletionprotection", "SWAGGER_OVERRIDE_/api/v2/recordings/deletionprotection", "SWAGGER_OVERRIDE_/api/v2/recordings/deletionprotection", )
	recordings_deletionprotectionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recordings_deletionprotection"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recordings_deletionprotectionCmd)
}

func Cmdrecordings_deletionprotection() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/recordings/deletionprotection", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Recording", "/api/v2/recordings/deletionprotection")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "conversationIds",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/ConversationDeletionProtectionQuery"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "array",
    "items" : {
      "$ref" : "#/definitions/AddressableEntityRef"
    }
  }
}`)
	recordings_deletionprotectionCmd.AddCommand(createCmd)
	
	utils.AddFlag(updateCmd.Flags(), "bool", "protect", "true", "Check for apply, uncheck for revoke (each action requires the respective permission)")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/recordings/deletionprotection", utils.FormatPermissions([]string{ "recording:deletionProtection:apply", "recording:deletionProtection:revoke",  }), utils.GenerateDevCentreLink("PUT", "Recording", "/api/v2/recordings/deletionprotection")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/ConversationDeletionProtectionQuery"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "Operation was successful."
}`)
	recordings_deletionprotectionCmd.AddCommand(updateCmd)
	
	return recordings_deletionprotectionCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Get a list of conversations with protected recordings",
	Long:  "Get a list of conversations with protected recordings",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Conversationdeletionprotectionquery{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recordings/deletionprotection"

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Apply or revoke recording protection for conversations",
	Long:  "Apply or revoke recording protection for conversations",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Conversationdeletionprotectionquery{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recordings/deletionprotection"

		protect := utils.GetFlag(cmd.Flags(), "bool", "protect")
		if protect != "" {
			queryParams["protect"] = protect
		}
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
