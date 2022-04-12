package responsemanagement_responseassets_status

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
	Description = utils.FormatUsageDescription("responsemanagement_responseassets_status", "SWAGGER_OVERRIDE_/api/v2/responsemanagement/responseassets/status", )
	responsemanagement_responseassets_statusCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("responsemanagement_responseassets_status"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(responsemanagement_responseassets_statusCmd)
}

func Cmdresponsemanagement_responseassets_status() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/responsemanagement/responseassets/status/{statusId}", utils.FormatPermissions([]string{ "responseAssets:asset:view",  }), utils.GenerateDevCentreLink("GET", "Response Management", "/api/v2/responsemanagement/responseassets/status/{statusId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/ResponseAssetStatus"
  }
}`)
	responsemanagement_responseassets_statusCmd.AddCommand(getCmd)
	
	return responsemanagement_responseassets_statusCmd
}

var getCmd = &cobra.Command{
	Use:   "get [statusId]",
	Short: "Get response asset upload status",
	Long:  "Get response asset upload status",
	Args:  utils.DetermineArgs([]string{ "statusId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/responsemanagement/responseassets/status/{statusId}"
		statusId, args := args[0], args[1:]
		path = strings.Replace(path, "{statusId}", fmt.Sprintf("%v", statusId), -1)

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