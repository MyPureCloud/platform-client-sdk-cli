package journey_actiontemplates

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("journey_actiontemplates", "SWAGGER_OVERRIDE_/api/v2/journey/actiontemplates", "SWAGGER_OVERRIDE_/api/v2/journey/actiontemplates", "SWAGGER_OVERRIDE_/api/v2/journey/actiontemplates", "SWAGGER_OVERRIDE_/api/v2/journey/actiontemplates", "SWAGGER_OVERRIDE_/api/v2/journey/actiontemplates", )
	journey_actiontemplatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_actiontemplates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_actiontemplatesCmd)
}

func Cmdjourney_actiontemplates() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/journey/actiontemplates", utils.FormatPermissions([]string{ "journey:actiontemplate:add",  }), utils.GenerateDevCentreLink("POST", "Journey", "/api/v2/journey/actiontemplates")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;required&quot; : false,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ActionTemplate&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ActionTemplate&quot;
  }
}`)
	journey_actiontemplatesCmd.AddCommand(createCmd)
	
	utils.AddFlag(deleteCmd.Flags(), "bool", "hardDelete", "", "Determines whether Action Template should be soft-deleted (have it`s state set to deleted) or hard-deleted (permanently removed). Set to false (soft-delete) by default.")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/journey/actiontemplates/{actionTemplateId}", utils.FormatPermissions([]string{ "journey:actiontemplate:delete",  }), utils.GenerateDevCentreLink("DELETE", "Journey", "/api/v2/journey/actiontemplates/{actionTemplateId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Action template deleted.&quot;
}`)
	journey_actiontemplatesCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/journey/actiontemplates/{actionTemplateId}", utils.FormatPermissions([]string{ "journey:actiontemplate:view",  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/actiontemplates/{actionTemplateId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ActionTemplate&quot;
  }
}`)
	journey_actiontemplatesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "Field(s) to sort by. Prefix with `-` for descending (e.g. sortBy=name,-createdDate).")
	utils.AddFlag(listCmd.Flags(), "string", "mediaType", "", "Media type Valid values: webchat, webMessagingOffer, contentOffer, integrationAction, architectFlow")
	utils.AddFlag(listCmd.Flags(), "string", "state", "", "Action template state. Valid values: Active, Inactive, Deleted")
	utils.AddFlag(listCmd.Flags(), "[]string", "queryFields", "", "ActionTemplate field(s) to query on. Requires `queryValue` to also be set.")
	utils.AddFlag(listCmd.Flags(), "string", "queryValue", "", "Value to query on. Requires `queryFields` to also be set.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/journey/actiontemplates", utils.FormatPermissions([]string{ "journey:actiontemplate:view",  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/actiontemplates")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	journey_actiontemplatesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/journey/actiontemplates/{actionTemplateId}", utils.FormatPermissions([]string{ "journey:actiontemplate:edit",  }), utils.GenerateDevCentreLink("PATCH", "Journey", "/api/v2/journey/actiontemplates/{actionTemplateId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;required&quot; : false,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/PatchActionTemplate&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ActionTemplate&quot;
  }
}`)
	journey_actiontemplatesCmd.AddCommand(updateCmd)
	
	return journey_actiontemplatesCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a single action template.",
	Long:  "Create a single action template.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/journey/actiontemplates"

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
	Use:   "delete [actionTemplateId]",
	Short: "Delete a single action template.",
	Long:  "Delete a single action template.",
	Args:  utils.DetermineArgs([]string{ "actionTemplateId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/journey/actiontemplates/{actionTemplateId}"
		actionTemplateId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionTemplateId}", fmt.Sprintf("%v", actionTemplateId), -1)

		hardDelete := utils.GetFlag(cmd.Flags(), "bool", "hardDelete")
		if hardDelete != "" {
			queryParams["hardDelete"] = hardDelete
		}
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
	Use:   "get [actionTemplateId]",
	Short: "Retrieve a single action template.",
	Long:  "Retrieve a single action template.",
	Args:  utils.DetermineArgs([]string{ "actionTemplateId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/journey/actiontemplates/{actionTemplateId}"
		actionTemplateId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionTemplateId}", fmt.Sprintf("%v", actionTemplateId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve all action templates.",
	Long:  "Retrieve all action templates.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/journey/actiontemplates"

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		mediaType := utils.GetFlag(cmd.Flags(), "string", "mediaType")
		if mediaType != "" {
			queryParams["mediaType"] = mediaType
		}
		state := utils.GetFlag(cmd.Flags(), "string", "state")
		if state != "" {
			queryParams["state"] = state
		}
		queryFields := utils.GetFlag(cmd.Flags(), "[]string", "queryFields")
		if queryFields != "" {
			queryParams["queryFields"] = queryFields
		}
		queryValue := utils.GetFlag(cmd.Flags(), "string", "queryValue")
		if queryValue != "" {
			queryParams["queryValue"] = queryValue
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
	Use:   "update [actionTemplateId]",
	Short: "Update a single action template.",
	Long:  "Update a single action template.",
	Args:  utils.DetermineArgs([]string{ "actionTemplateId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/journey/actiontemplates/{actionTemplateId}"
		actionTemplateId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionTemplateId}", fmt.Sprintf("%v", actionTemplateId), -1)

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
