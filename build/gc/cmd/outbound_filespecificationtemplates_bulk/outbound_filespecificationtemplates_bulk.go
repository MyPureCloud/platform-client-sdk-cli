package outbound_filespecificationtemplates_bulk

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
	Description = utils.FormatUsageDescription("outbound_filespecificationtemplates_bulk", "SWAGGER_OVERRIDE_/api/v2/outbound/filespecificationtemplates/bulk", )
	outbound_filespecificationtemplates_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_filespecificationtemplates_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_filespecificationtemplates_bulkCmd)
}

func Cmdoutbound_filespecificationtemplates_bulk() *cobra.Command { 
	utils.AddFlag(deleteCmd.Flags(), "[]string", "id", "", "File Specification template id(s) to delete - REQUIRED")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/filespecificationtemplates/bulk", utils.FormatPermissions([]string{ "outbound:fileSpecificationTemplate:delete",  }), utils.GenerateDevCentreLink("DELETE", "Outbound", "/api/v2/outbound/filespecificationtemplates/bulk")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	deleteCmd.MarkFlagRequired("id")
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "File specification templates accepted for delete.",
  "content" : { }
}`)
	outbound_filespecificationtemplates_bulkCmd.AddCommand(deleteCmd)
	return outbound_filespecificationtemplates_bulkCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete multiple file specification templates.",
	Long:  "Delete multiple file specification templates.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/filespecificationtemplates/bulk"

		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
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
