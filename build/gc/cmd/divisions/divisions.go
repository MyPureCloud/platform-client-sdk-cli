package divisions

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
	divisionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("divisions"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud divisions"),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud divisions`),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(divisionsCmd)
}

func Cmddivisions() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/authorization/divisions", utils.FormatPermissions([]string{ "authorization:division:add", "authorization:grant:add",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AuthzDivision&quot;
  }
}`)
	divisionsCmd.AddCommand(createCmd)
	
	utils.AddFlag(deleteCmd.Flags(), "bool", "force", "false", "Force delete this division as well as the grants and objects associated with it")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/authorization/divisions/{divisionId}", utils.FormatPermissions([]string{ "authorization:division:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Deleted&quot;
}`)
	divisionsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "bool", "objectCount", "false", "Get count of objects in this division, grouped by type")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/authorization/divisions/{divisionId}", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AuthzDivision&quot;
  }
}`)
	divisionsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "The total page size requested")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The page number requested")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "variable name requested to sort by")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "variable name requested by expand list")
	utils.AddFlag(listCmd.Flags(), "string", "nextPage", "", "next page token")
	utils.AddFlag(listCmd.Flags(), "string", "previousPage", "", "Previous page token")
	utils.AddFlag(listCmd.Flags(), "bool", "objectCount", "false", "Include the count of objects contained in the division")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "Optionally request specific divisions by their IDs")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Search term to filter by division name")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/authorization/divisions", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AuthzDivisionEntityListing&quot;
  }
}`)
	divisionsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/authorization/divisions/{divisionId}", utils.FormatPermissions([]string{ "authorization:division:edit",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT")
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AuthzDivision&quot;
  }
}`)
	divisionsCmd.AddCommand(updateCmd)
	
	return divisionsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a division.",
	Long:  `Create a division.`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions"

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
	Use:   "delete [divisionId]",
	Short: "Delete a division.",
	Long:  `Delete a division.`,
	Args:  utils.DetermineArgs([]string{ "divisionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions/{divisionId}"
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)

		force := utils.GetFlag(cmd.Flags(), "bool", "force")
		if force != "" {
			queryParams["force"] = force
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
	Use:   "get [divisionId]",
	Short: "Returns an authorization division.",
	Long:  `Returns an authorization division.`,
	Args:  utils.DetermineArgs([]string{ "divisionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions/{divisionId}"
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)

		objectCount := utils.GetFlag(cmd.Flags(), "bool", "objectCount")
		if objectCount != "" {
			queryParams["objectCount"] = objectCount
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
	Use:   "list",
	Short: "Retrieve a list of all divisions defined for the organization",
	Long:  `Retrieve a list of all divisions defined for the organization`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions"

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
		objectCount := utils.GetFlag(cmd.Flags(), "bool", "objectCount")
		if objectCount != "" {
			queryParams["objectCount"] = objectCount
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
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
	Use:   "update [divisionId]",
	Short: "Update a division.",
	Long:  `Update a division.`,
	Args:  utils.DetermineArgs([]string{ "divisionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions/{divisionId}"
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)

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
