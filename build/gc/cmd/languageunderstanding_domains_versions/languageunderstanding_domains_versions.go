package languageunderstanding_domains_versions

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
	Description = utils.FormatUsageDescription("languageunderstanding_domains_versions", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/versions", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/versions", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/versions", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/versions", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/versions", )
	languageunderstanding_domains_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding_domains_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstanding_domains_versionsCmd)
}

func Cmdlanguageunderstanding_domains_versions() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/languageunderstanding/domains/{domainId}/versions", utils.FormatPermissions([]string{ "languageUnderstanding:nluDomainVersion:add", "dialog:botVersion:add",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;The NLU Domain Version to create.&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/NluDomainVersion&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/NluDomainVersion&quot;
  }
}`)
	languageunderstanding_domains_versionsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}", utils.FormatPermissions([]string{ "languageUnderstanding:nluDomainVersion:delete", "dialog:botVersion:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Deleted the specified NLU Domain Version&quot;
}`)
	languageunderstanding_domains_versionsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "bool", "includeUtterances", "", "Whether utterances for intent definition should be included when marshalling response.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}", utils.FormatPermissions([]string{ "languageUnderstanding:nluDomainVersion:view", "dialog:botVersion:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;Retrieved the specified NLU Domain Version&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/NluDomainVersion&quot;
  }
}`)
	languageunderstanding_domains_versionsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "bool", "includeUtterances", "", "Whether utterances for intent definition should be included when marshalling response.")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/languageunderstanding/domains/{domainId}/versions", utils.FormatPermissions([]string{ "languageUnderstanding:nluDomainVersion:view", "dialog:botVersion:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	languageunderstanding_domains_versionsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}", utils.FormatPermissions([]string{ "languageUnderstanding:nluDomainVersion:edit", "dialog:botVersion:edit",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;The updated NLU Domain Version.&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/NluDomainVersion&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;Updated the specified NLU Domain Version&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/NluDomainVersion&quot;
  }
}`)
	languageunderstanding_domains_versionsCmd.AddCommand(updateCmd)
	
	return languageunderstanding_domains_versionsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [domainId]",
	Short: "Create an NLU Domain Version.",
	Long:  "Create an NLU Domain Version.",
	Args:  utils.DetermineArgs([]string{ "domainId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/versions"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)

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
	Use:   "delete [domainId] [domainVersionId]",
	Short: "Delete an NLU Domain Version",
	Long:  "Delete an NLU Domain Version",
	Args:  utils.DetermineArgs([]string{ "domainId", "domainVersionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
		domainVersionId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)

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
	Use:   "get [domainId] [domainVersionId]",
	Short: "Find an NLU Domain Version.",
	Long:  "Find an NLU Domain Version.",
	Args:  utils.DetermineArgs([]string{ "domainId", "domainVersionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
		domainVersionId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)

		includeUtterances := utils.GetFlag(cmd.Flags(), "bool", "includeUtterances")
		if includeUtterances != "" {
			queryParams["includeUtterances"] = includeUtterances
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
var listCmd = &cobra.Command{
	Use:   "list [domainId]",
	Short: "Get all NLU Domain Versions for a given Domain.",
	Long:  "Get all NLU Domain Versions for a given Domain.",
	Args:  utils.DetermineArgs([]string{ "domainId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/versions"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)

		includeUtterances := utils.GetFlag(cmd.Flags(), "bool", "includeUtterances")
		if includeUtterances != "" {
			queryParams["includeUtterances"] = includeUtterances
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
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
	Use:   "update [domainId] [domainVersionId]",
	Short: "Update an NLU Domain Version.",
	Long:  "Update an NLU Domain Version.",
	Args:  utils.DetermineArgs([]string{ "domainId", "domainVersionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
		domainVersionId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
