package externalcontacts_scan_relationships

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
	Description = utils.FormatUsageDescription("externalcontacts_scan_relationships", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/scan/relationships", )
	externalcontacts_scan_relationshipsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_scan_relationships"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_scan_relationshipsCmd)
}

func Cmdexternalcontacts_scan_relationships() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "limit", "", "The number of relationships per page; must be between 10 and 200, default is 100)")
	utils.AddFlag(listCmd.Flags(), "string", "cursor", "", "Indicates where to resume query results (not required for first page), each page returns a new cursor with a 24h TTL")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/scan/relationships", utils.FormatPermissions([]string{ "externalContacts:contact:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/scan/relationships")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	externalcontacts_scan_relationshipsCmd.AddCommand(listCmd)
	
	return externalcontacts_scan_relationshipsCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Scan for relationships",
	Long:  "Scan for relationships",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/scan/relationships"

		limit := utils.GetFlag(cmd.Flags(), "int", "limit")
		if limit != "" {
			queryParams["limit"] = limit
		}
		cursor := utils.GetFlag(cmd.Flags(), "string", "cursor")
		if cursor != "" {
			queryParams["cursor"] = cursor
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
