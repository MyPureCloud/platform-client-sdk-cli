package outbound_dnclists_phonenumbers

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
	Description = utils.FormatUsageDescription("outbound_dnclists_phonenumbers", "SWAGGER_OVERRIDE_/api/v2/outbound/dnclists/{dncListId}/phonenumbers", )
	outbound_dnclists_phonenumbersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_dnclists_phonenumbers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_dnclists_phonenumbersCmd)
}

func Cmdoutbound_dnclists_phonenumbers() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/dnclists/{dncListId}/phonenumbers", utils.FormatPermissions([]string{ "outbound:dnc:add",  }), utils.GenerateDevCentreLink("POST", "Outbound", "/api/v2/outbound/dnclists/{dncListId}/phonenumbers")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;DNC Phone Numbers&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;type&quot; : &quot;string&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;The request could not be understood by the server due to malformed syntax.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ErrorBody&quot;
  },
  &quot;x-inin-error-codes&quot; : {
    &quot;dnc.source.operation.not.supported&quot; : &quot;An attempt was made to append numbers to a DNC list that is not of type Internal&quot;,
    &quot;dnc.phone.numbers.per.list.limit.exceeded&quot; : &quot;The DNC list has reached the limit on total records. See details&quot;,
    &quot;bad.request&quot; : &quot;The request could not be understood by the server due to malformed syntax.&quot;,
    &quot;response.entity.too.large&quot; : &quot;The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable&quot;,
    &quot;invalid.date&quot; : &quot;Dates must be specified as ISO-8601 strings. For example: yyyy-MM-ddTHH:mm:ss.SSSZ&quot;,
    &quot;dnc.records.per.organization.limit.exceeded&quot; : &quot;The organization has reached the limit on total DNC records. See details&quot;,
    &quot;dnc.records.per.list.limit.exceeded&quot; : &quot;The DNC list has reached the limit on total records. See details&quot;,
    &quot;invalid.value&quot; : &quot;Value [%s] is not valid for field type [%s]. Allowable values are: %s&quot;,
    &quot;dnc.phone.numbers.per.organization.limit.exceeded&quot; : &quot;The organization has reached the limit on total DNC records. See details&quot;
  }
}`)
	outbound_dnclists_phonenumbersCmd.AddCommand(createCmd)
	
	return outbound_dnclists_phonenumbersCmd
}

var createCmd = &cobra.Command{
	Use:   "create [dncListId]",
	Short: "Add phone numbers to a DNC list.",
	Long:  "Add phone numbers to a DNC list.",
	Args:  utils.DetermineArgs([]string{ "dncListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/dnclists/{dncListId}/phonenumbers"
		dncListId, args := args[0], args[1:]
		path = strings.Replace(path, "{dncListId}", fmt.Sprintf("%v", dncListId), -1)

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
