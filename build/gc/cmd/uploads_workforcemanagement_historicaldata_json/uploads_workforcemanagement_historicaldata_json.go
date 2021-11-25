package uploads_workforcemanagement_historicaldata_json

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
	Description = utils.FormatUsageDescription("uploads_workforcemanagement_historicaldata_json", "SWAGGER_OVERRIDE_/api/v2/uploads/workforcemanagement/historicaldata/json", )
	uploads_workforcemanagement_historicaldata_jsonCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("uploads_workforcemanagement_historicaldata_json"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(uploads_workforcemanagement_historicaldata_jsonCmd)
}

func Cmduploads_workforcemanagement_historicaldata_json() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/uploads/workforcemanagement/historicaldata/json", utils.FormatPermissions([]string{ "wfm:historicalData:upload",  }), utils.GenerateDevCentreLink("POST", "Uploads", "/api/v2/uploads/workforcemanagement/historicaldata/json")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "query",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/UploadUrlRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Presigned url successfully created.",
  "schema" : {
    "$ref" : "#/definitions/UploadUrlResponse"
  }
}`)
	uploads_workforcemanagement_historicaldata_jsonCmd.AddCommand(createCmd)
	
	return uploads_workforcemanagement_historicaldata_jsonCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates presigned url for uploading WFM historical data file. Requires data in json format.",
	Long:  "Creates presigned url for uploading WFM historical data file. Requires data in json format.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Uploadurlrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/uploads/workforcemanagement/historicaldata/json"

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
