package recording_batchrequests

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
	Description = utils.FormatUsageDescription("recording_batchrequests", "SWAGGER_OVERRIDE_/api/v2/recording/batchrequests", "SWAGGER_OVERRIDE_/api/v2/recording/batchrequests", )
	recording_batchrequestsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recording_batchrequests"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recording_batchrequestsCmd)
}

func Cmdrecording_batchrequests() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/recording/batchrequests", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Recording", "/api/v2/recording/batchrequests")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Job submission criteria",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/BatchDownloadJobSubmission"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/BatchDownloadJobSubmissionResult"
  }
}`)
	recording_batchrequestsCmd.AddCommand(createCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/recording/batchrequests/{jobId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/recording/batchrequests/{jobId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/BatchDownloadJobStatusResult"
  }
}`)
	recording_batchrequestsCmd.AddCommand(getCmd)
	
	return recording_batchrequestsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Submit a batch download request for recordings. Recordings in response will be in their original format/codec - configured in the Trunk configuration.",
	Long:  "Submit a batch download request for recordings. Recordings in response will be in their original format/codec - configured in the Trunk configuration.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Batchdownloadjobsubmission{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/batchrequests"


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
var getCmd = &cobra.Command{
	Use:   "get [jobId]",
	Short: "Get the status and results for a batch request job, only the user that submitted the job may retrieve results",
	Long:  "Get the status and results for a batch request job, only the user that submitted the job may retrieve results",
	Args:  utils.DetermineArgs([]string{ "jobId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/recording/batchrequests/{jobId}"
		jobId, args := args[0], args[1:]
		path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)


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
