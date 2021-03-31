package dids

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
	didsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("dids"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud dids"),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud dids`),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(didsCmd)
}

func Cmddids() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/dids/{didId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DID&quot;
  }
}`)
	didsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "number", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ASC", "Sort order")
	utils.AddFlag(listCmd.Flags(), "string", "phoneNumber", "", "Filter by phoneNumber")
	utils.AddFlag(listCmd.Flags(), "string", "ownerId", "", "Filter by the owner of a phone number")
	utils.AddFlag(listCmd.Flags(), "string", "didPoolId", "", "Filter by the DID Pool assignment")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "Filter by a specific list of ID`s")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/dids", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DIDEntityListing&quot;
  }
}`)
	didsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/telephony/providers/edges/dids/{didId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;DID&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DID&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DID&quot;
  }
}`)
	didsCmd.AddCommand(updateCmd)
	
	return didsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [didId]",
	Short: "Get a DID by ID.",
	Long:  `Get a DID by ID.`,
	Args:  utils.DetermineArgs([]string{ "didId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/dids/{didId}"
		didId, args := args[0], args[1:]
		path = strings.Replace(path, "{didId}", fmt.Sprintf("%v", didId), -1)

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
	Short: "Get a listing of DIDs",
	Long:  `Get a listing of DIDs`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/dids"

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
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		phoneNumber := utils.GetFlag(cmd.Flags(), "string", "phoneNumber")
		if phoneNumber != "" {
			queryParams["phoneNumber"] = phoneNumber
		}
		ownerId := utils.GetFlag(cmd.Flags(), "string", "ownerId")
		if ownerId != "" {
			queryParams["ownerId"] = ownerId
		}
		didPoolId := utils.GetFlag(cmd.Flags(), "string", "didPoolId")
		if didPoolId != "" {
			queryParams["didPoolId"] = didPoolId
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
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
	Use:   "update [didId]",
	Short: "Update a DID by ID.",
	Long:  `Update a DID by ID.`,
	Args:  utils.DetermineArgs([]string{ "didId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/dids/{didId}"
		didId, args := args[0], args[1:]
		path = strings.Replace(path, "{didId}", fmt.Sprintf("%v", didId), -1)

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
