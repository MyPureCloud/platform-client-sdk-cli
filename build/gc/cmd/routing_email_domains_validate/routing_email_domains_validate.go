package routing_email_domains_validate

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
	Description = utils.FormatUsageDescription("routing_email_domains_validate", "SWAGGER_OVERRIDE_/api/v2/routing/email/domains/{domainId}/validate", )
	routing_email_domains_validateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_email_domains_validate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_email_domains_validateCmd)
}

func Cmdrouting_email_domains_validate() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/routing/email/domains/{domainId}/validate", utils.FormatPermissions([]string{ "routing:email:manage",  }), utils.GenerateDevCentreLink("PATCH", "Routing", "/api/v2/routing/email/domains/{domainId}/validate")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "description" : "Domain settings",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/InboundDomainPatchRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/InboundDomain"
  }
}`)
	routing_email_domains_validateCmd.AddCommand(updateCmd)
	
	return routing_email_domains_validateCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [domainId]",
	Short: "Validate domain settings",
	Long:  "Validate domain settings",
	Args:  utils.DetermineArgs([]string{ "domainId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Inbounddomainpatchrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/email/domains/{domainId}/validate"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
