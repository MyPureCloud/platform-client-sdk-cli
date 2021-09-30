package integrations_workforcemanagement_vendorconnection

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
	Description = utils.FormatUsageDescription("integrations_workforcemanagement_vendorconnection", "SWAGGER_OVERRIDE_/api/v2/integrations/workforcemanagement/vendorconnection", )
	integrations_workforcemanagement_vendorconnectionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_workforcemanagement_vendorconnection"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_workforcemanagement_vendorconnectionCmd)
}

func Cmdintegrations_workforcemanagement_vendorconnection() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/integrations/workforcemanagement/vendorconnection", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Integrations", "/api/v2/integrations/workforcemanagement/vendorconnection")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/VendorConnectionRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/UserActionCategoryEntityListing"
  }
}`)
	integrations_workforcemanagement_vendorconnectionCmd.AddCommand(createCmd)
	
	return integrations_workforcemanagement_vendorconnectionCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a vendor connection",
	Long:  "Add a vendor connection",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Vendorconnectionrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/workforcemanagement/vendorconnection"


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
