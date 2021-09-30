package integrations_actions_draft

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
	Description = utils.FormatUsageDescription("integrations_actions_draft", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/{actionId}/draft", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/{actionId}/draft", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/{actionId}/draft", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/{actionId}/draft", )
	integrations_actions_draftCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_actions_draft"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_actions_draftCmd)
}

func Cmdintegrations_actions_draft() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/integrations/actions/{actionId}/draft", utils.FormatPermissions([]string{ "integrations:action:edit",  }), utils.GenerateDevCentreLink("POST", "Integrations", "/api/v2/integrations/actions/{actionId}/draft")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Action"
  }
}`)
	integrations_actions_draftCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/integrations/actions/{actionId}/draft", utils.FormatPermissions([]string{ "integrations:action:delete",  }), utils.GenerateDevCentreLink("DELETE", "Integrations", "/api/v2/integrations/actions/{actionId}/draft")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Delete was successful"
}`)
	integrations_actions_draftCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "string", "expand", "", "Indicates a field in the response which should be expanded. Valid values: contract")
	utils.AddFlag(getCmd.Flags(), "bool", "includeConfig", "false", "Return config in response.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/integrations/actions/{actionId}/draft", utils.FormatPermissions([]string{ "integrations:action:view", "bridge:actions:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/actions/{actionId}/draft")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Action"
  }
}`)
	integrations_actions_draftCmd.AddCommand(getCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/integrations/actions/{actionId}/draft", utils.FormatPermissions([]string{ "integrations:action:edit",  }), utils.GenerateDevCentreLink("PATCH", "Integrations", "/api/v2/integrations/actions/{actionId}/draft")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "description" : "Input used to patch the Action Draft.",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/UpdateDraftInput"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Action"
  }
}`)
	integrations_actions_draftCmd.AddCommand(updateCmd)
	
	return integrations_actions_draftCmd
}

var createCmd = &cobra.Command{
	Use:   "create [actionId]",
	Short: "Create a new Draft from existing Action",
	Long:  "Create a new Draft from existing Action",
	Args:  utils.DetermineArgs([]string{ "actionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/actions/{actionId}/draft"
		actionId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
var deleteCmd = &cobra.Command{
	Use:   "delete [actionId]",
	Short: "Delete a Draft",
	Long:  "Delete a Draft",
	Args:  utils.DetermineArgs([]string{ "actionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/actions/{actionId}/draft"
		actionId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
var getCmd = &cobra.Command{
	Use:   "get [actionId]",
	Short: "Retrieve a Draft",
	Long:  "Retrieve a Draft",
	Args:  utils.DetermineArgs([]string{ "actionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/actions/{actionId}/draft"
		actionId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)


		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		includeConfig := utils.GetFlag(cmd.Flags(), "bool", "includeConfig")
		if includeConfig != "" {
			queryParams["includeConfig"] = includeConfig
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
	Use:   "update [actionId]",
	Short: "Update an existing Draft",
	Long:  "Update an existing Draft",
	Args:  utils.DetermineArgs([]string{ "actionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Updatedraftinput{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/actions/{actionId}/draft"
		actionId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
