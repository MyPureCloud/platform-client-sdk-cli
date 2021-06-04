package telephony_providers_edges_logs_jobs_upload

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_logs_jobs_upload", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload", )
	telephony_providers_edges_logs_jobs_uploadCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_logs_jobs_upload"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_logs_jobs_uploadCmd)
}

func Cmdtelephony_providers_edges_logs_jobs_upload() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Log upload request&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/EdgeLogsJobUploadRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Accepted - Files are being uploaded to the job.  Watch the uploadStatus property on the job files.&quot;
}`)
	telephony_providers_edges_logs_jobs_uploadCmd.AddCommand(createCmd)
	
	return telephony_providers_edges_logs_jobs_uploadCmd
}

var createCmd = &cobra.Command{
	Use:   "create [edgeId] [jobId]",
	Short: "Request that the specified fileIds be uploaded from the Edge.",
	Long:  "Request that the specified fileIds be uploaded from the Edge.",
	Args:  utils.DetermineArgs([]string{ "edgeId", "jobId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload"
		edgeId, args := args[0], args[1:]
		path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
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
