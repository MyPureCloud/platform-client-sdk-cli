package recording_mediaretentionpolicies

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
	Description = utils.FormatUsageDescription("recording_mediaretentionpolicies", "SWAGGER_OVERRIDE_/api/v2/recording/mediaretentionpolicies", "SWAGGER_OVERRIDE_/api/v2/recording/mediaretentionpolicies", "SWAGGER_OVERRIDE_/api/v2/recording/mediaretentionpolicies", "SWAGGER_OVERRIDE_/api/v2/recording/mediaretentionpolicies", "SWAGGER_OVERRIDE_/api/v2/recording/mediaretentionpolicies", "SWAGGER_OVERRIDE_/api/v2/recording/mediaretentionpolicies", "SWAGGER_OVERRIDE_/api/v2/recording/mediaretentionpolicies", )
	recording_mediaretentionpoliciesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recording_mediaretentionpolicies"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recording_mediaretentionpoliciesCmd)
}

func Cmdrecording_mediaretentionpolicies() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/recording/mediaretentionpolicies", utils.FormatPermissions([]string{ "recording:retentionPolicy:add",  }), utils.GenerateDevCentreLink("POST", "Recording", "/api/v2/recording/mediaretentionpolicies")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Policy",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/PolicyCreate"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Policy"
  }
}`)
	recording_mediaretentionpoliciesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/recording/mediaretentionpolicies/{policyId}", utils.FormatPermissions([]string{ "recording:retentionPolicy:delete",  }), utils.GenerateDevCentreLink("DELETE", "Recording", "/api/v2/recording/mediaretentionpolicies/{policyId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Operation was successful."
}`)
	recording_mediaretentionpoliciesCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(deletepoliciesCmd.Flags(), "string", "ids", "", " - REQUIRED")
	deletepoliciesCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deletepoliciesCmd.UsageTemplate(), "DELETE", "/api/v2/recording/mediaretentionpolicies", utils.FormatPermissions([]string{ "recording:retentionPolicy:delete",  }), utils.GenerateDevCentreLink("DELETE", "Recording", "/api/v2/recording/mediaretentionpolicies")))
	utils.AddFileFlagIfUpsert(deletepoliciesCmd.Flags(), "DELETE", ``)
	deletepoliciesCmd.MarkFlagRequired("ids")
	
	utils.AddPaginateFlagsIfListingResponse(deletepoliciesCmd.Flags(), "DELETE", `{
  "description" : "Operation was successful."
}`)
	recording_mediaretentionpoliciesCmd.AddCommand(deletepoliciesCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/recording/mediaretentionpolicies/{policyId}", utils.FormatPermissions([]string{ "recording:retentionPolicy:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/recording/mediaretentionpolicies/{policyId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Policy"
  }
}`)
	recording_mediaretentionpoliciesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "The total page size requested")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The page number requested")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "variable name requested to sort by")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "variable name requested by expand list")
	utils.AddFlag(listCmd.Flags(), "string", "nextPage", "", "next page token")
	utils.AddFlag(listCmd.Flags(), "string", "previousPage", "", "Previous page token")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "the policy name - used for filtering results in searches.")
	utils.AddFlag(listCmd.Flags(), "bool", "enabled", "", "checks to see if policy is enabled - use enabled = true or enabled = false")
	utils.AddFlag(listCmd.Flags(), "bool", "summary", "false", "provides a less verbose response of policy lists.")
	utils.AddFlag(listCmd.Flags(), "bool", "hasErrors", "", "provides a way to fetch all policies with errors or policies that do not have errors")
	utils.AddFlag(listCmd.Flags(), "int", "deleteDaysThreshold", "", "provides a way to fetch all policies with any actions having deleteDays exceeding the provided value")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/recording/mediaretentionpolicies", utils.FormatPermissions([]string{ "recording:retentionPolicy:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/recording/mediaretentionpolicies")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	recording_mediaretentionpoliciesCmd.AddCommand(listCmd)
	
	patchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", patchCmd.UsageTemplate(), "PATCH", "/api/v2/recording/mediaretentionpolicies/{policyId}", utils.FormatPermissions([]string{ "recording:retentionPolicy:edit",  }), utils.GenerateDevCentreLink("PATCH", "Recording", "/api/v2/recording/mediaretentionpolicies/{policyId}")))
	utils.AddFileFlagIfUpsert(patchCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "description" : "Policy",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/PolicyUpdate"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(patchCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Policy"
  }
}`)
	recording_mediaretentionpoliciesCmd.AddCommand(patchCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/recording/mediaretentionpolicies/{policyId}", utils.FormatPermissions([]string{ "recording:retentionPolicy:edit",  }), utils.GenerateDevCentreLink("PUT", "Recording", "/api/v2/recording/mediaretentionpolicies/{policyId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "Policy",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/Policy"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Policy"
  }
}`)
	recording_mediaretentionpoliciesCmd.AddCommand(updateCmd)
	
	return recording_mediaretentionpoliciesCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create media retention policy",
	Long:  "Create media retention policy",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Policycreate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/mediaretentionpolicies"

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
var deleteCmd = &cobra.Command{
	Use:   "delete [policyId]",
	Short: "Delete a media retention policy",
	Long:  "Delete a media retention policy",
	Args:  utils.DetermineArgs([]string{ "policyId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/mediaretentionpolicies/{policyId}"
		policyId, args := args[0], args[1:]
		path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var deletepoliciesCmd = &cobra.Command{
	Use:   "deletepolicies",
	Short: "Delete media retention policies",
	Long:  "Delete media retention policies",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/mediaretentionpolicies"

		ids := utils.GetFlag(cmd.Flags(), "string", "ids")
		if ids != "" {
			queryParams["ids"] = ids
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "deletepolicies"
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [policyId]",
	Short: "Get a media retention policy",
	Long:  "Get a media retention policy",
	Args:  utils.DetermineArgs([]string{ "policyId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/mediaretentionpolicies/{policyId}"
		policyId, args := args[0], args[1:]
		path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Gets media retention policy list with query options to filter on name and enabled.",
	Long:  "Gets media retention policy list with query options to filter on name and enabled.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/mediaretentionpolicies"

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
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		enabled := utils.GetFlag(cmd.Flags(), "bool", "enabled")
		if enabled != "" {
			queryParams["enabled"] = enabled
		}
		summary := utils.GetFlag(cmd.Flags(), "bool", "summary")
		if summary != "" {
			queryParams["summary"] = summary
		}
		hasErrors := utils.GetFlag(cmd.Flags(), "bool", "hasErrors")
		if hasErrors != "" {
			queryParams["hasErrors"] = hasErrors
		}
		deleteDaysThreshold := utils.GetFlag(cmd.Flags(), "int", "deleteDaysThreshold")
		if deleteDaysThreshold != "" {
			queryParams["deleteDaysThreshold"] = deleteDaysThreshold
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var patchCmd = &cobra.Command{
	Use:   "patch [policyId]",
	Short: "Patch a media retention policy",
	Long:  "Patch a media retention policy",
	Args:  utils.DetermineArgs([]string{ "policyId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Policyupdate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/mediaretentionpolicies/{policyId}"
		policyId, args := args[0], args[1:]
		path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "patch"
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [policyId]",
	Short: "Update a media retention policy",
	Long:  "Update a media retention policy",
	Args:  utils.DetermineArgs([]string{ "policyId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Policy{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/mediaretentionpolicies/{policyId}"
		policyId, args := args[0], args[1:]
		path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyId), -1)

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
