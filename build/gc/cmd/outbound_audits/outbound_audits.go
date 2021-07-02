package outbound_audits

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
	Description = utils.FormatUsageDescription("outbound_audits", "SWAGGER_OVERRIDE_/api/v2/outbound/audits", )
	outbound_auditsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_audits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_auditsCmd)
}

func Cmdoutbound_audits() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(createCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(createCmd.Flags(), "string", "sortBy", "entity.name", "Sort by")
	utils.AddFlag(createCmd.Flags(), "string", "sortOrder", "ascending", "Sort order")
	utils.AddFlag(createCmd.Flags(), "bool", "facetsOnly", "false", "Facets only")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/audits", utils.FormatPermissions([]string{ "outbound:audit:view",  }), utils.GenerateDevCentreLink("POST", "Outbound", "/api/v2/outbound/audits")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;AuditSearch&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DialerAuditRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AuditSearchResult&quot;
  }
}`)
	outbound_auditsCmd.AddCommand(createCmd)
	
	return outbound_auditsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Retrieves audits for dialer.",
	Long:  "Retrieves audits for dialer.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/audits"

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
		facetsOnly := utils.GetFlag(cmd.Flags(), "bool", "facetsOnly")
		if facetsOnly != "" {
			queryParams["facetsOnly"] = facetsOnly
		}
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
