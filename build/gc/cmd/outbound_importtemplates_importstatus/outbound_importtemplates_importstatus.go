package outbound_importtemplates_importstatus

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
	Description = utils.FormatUsageDescription("outbound_importtemplates_importstatus", "SWAGGER_OVERRIDE_/api/v2/outbound/importtemplates/{importTemplateId}/importstatus", )
	outbound_importtemplates_importstatusCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_importtemplates_importstatus"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_importtemplates_importstatusCmd)
}

func Cmdoutbound_importtemplates_importstatus() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "listNamePrefix", "", "listNamePrefix")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/importtemplates/{importTemplateId}/importstatus", utils.FormatPermissions([]string{ "outbound:importTemplate:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/importtemplates/{importTemplateId}/importstatus")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ImportStatus"
      }
    }
  }
}`)
	outbound_importtemplates_importstatusCmd.AddCommand(getCmd)
	return outbound_importtemplates_importstatusCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [importTemplateId]",
	Short: "Get the import status for an import template.",
	Long:  "Get the import status for an import template.",
	Args:  utils.DetermineArgs([]string{ "importTemplateId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/importtemplates/{importTemplateId}/importstatus"
		importTemplateId, args := args[0], args[1:]
		path = strings.Replace(path, "{importTemplateId}", fmt.Sprintf("%v", importTemplateId), -1)

		listNamePrefix := utils.GetFlag(cmd.Flags(), "string", "listNamePrefix")
		if listNamePrefix != "" {
			queryParams["listNamePrefix"] = listNamePrefix
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
