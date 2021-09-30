package outbound_dnclists_phonenumbers

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
  "in" : "body",
  "name" : "body",
  "description" : "DNC Phone Numbers",
  "required" : true,
  "schema" : {
    "type" : "array",
    "items" : {
      "type" : "string"
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "The request could not be understood by the server due to malformed syntax.",
  "schema" : {
    "$ref" : "#/definitions/ErrorBody"
  },
  "x-inin-error-codes" : {
    "dnc.source.operation.not.supported" : "An attempt was made to append numbers to a DNC list that is not of type Internal",
    "dnc.phone.numbers.per.list.limit.exceeded" : "The DNC list has reached the limit on total records. See details",
    "bad.request" : "The request could not be understood by the server due to malformed syntax.",
    "response.entity.too.large" : "The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable",
    "invalid.date" : "Dates must be specified as ISO-8601 strings. For example: yyyy-MM-ddTHH:mm:ss.SSSZ",
    "invalid.query.param.value" : "Value [%s] is not valid for parameter [%s]. Allowable values are: %s",
    "invalid.property" : "Value [%s] is not a valid property for object [%s]",
    "dnc.records.per.organization.limit.exceeded" : "The organization has reached the limit on total DNC records. See details",
    "dnc.records.per.list.limit.exceeded" : "The DNC list has reached the limit on total records. See details",
    "constraint.validation" : "%s",
    "invalid.value" : "Value [%s] is not valid for field type [%s]. Allowable values are: %s",
    "dnc.phone.numbers.per.organization.limit.exceeded" : "The organization has reached the limit on total DNC records. See details"
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
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

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
