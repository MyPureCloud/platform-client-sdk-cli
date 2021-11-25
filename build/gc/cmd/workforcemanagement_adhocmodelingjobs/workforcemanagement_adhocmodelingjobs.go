package workforcemanagement_adhocmodelingjobs

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
	Description = utils.FormatUsageDescription("workforcemanagement_adhocmodelingjobs", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/adhocmodelingjobs", )
	workforcemanagement_adhocmodelingjobsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_adhocmodelingjobs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_adhocmodelingjobsCmd)
}

func Cmdworkforcemanagement_adhocmodelingjobs() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/adhocmodelingjobs/{jobId}", utils.FormatPermissions([]string{ "wfm:adhocModel:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/adhocmodelingjobs/{jobId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/ModelingStatusResponse"
  }
}`)
	workforcemanagement_adhocmodelingjobsCmd.AddCommand(getCmd)
	
	return workforcemanagement_adhocmodelingjobsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [jobId]",
	Short: "Get status of the modeling job",
	Long:  "Get status of the modeling job",
	Args:  utils.DetermineArgs([]string{ "jobId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/adhocmodelingjobs/{jobId}"
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
