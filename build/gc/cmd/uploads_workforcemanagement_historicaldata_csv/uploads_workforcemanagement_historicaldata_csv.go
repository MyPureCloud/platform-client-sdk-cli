package uploads_workforcemanagement_historicaldata_csv

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
  "description" : "query",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UploadUrlRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Presigned url successfully created.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UploadUrlResponse"
      }
    }
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
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Uploadurlrequest{}
			utils.Render(reqModel.String())
			
			return
		}

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

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
