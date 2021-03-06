package uploads_workforcemanagement_historicaldata_csv

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
	Description = utils.FormatUsageDescription("uploads_workforcemanagement_historicaldata_csv", "SWAGGER_OVERRIDE_/api/v2/uploads/workforcemanagement/historicaldata/csv", )
	uploads_workforcemanagement_historicaldata_csvCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("uploads_workforcemanagement_historicaldata_csv"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(uploads_workforcemanagement_historicaldata_csvCmd)
}

func Cmduploads_workforcemanagement_historicaldata_csv() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/uploads/workforcemanagement/historicaldata/csv", utils.FormatPermissions([]string{ "wfm:historicalData:upload",  }), utils.GenerateDevCentreLink("POST", "Uploads", "/api/v2/uploads/workforcemanagement/historicaldata/csv")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;query&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UploadUrlRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Presigned url successfully created.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UploadUrlResponse&quot;
  }
}`)
	uploads_workforcemanagement_historicaldata_csvCmd.AddCommand(createCmd)
	
	return uploads_workforcemanagement_historicaldata_csvCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates presigned url for uploading WFM historical data file. Requires data in csv format.",
	Long:  "Creates presigned url for uploading WFM historical data file. Requires data in csv format.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/uploads/workforcemanagement/historicaldata/csv"

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
