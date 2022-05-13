package orgauthorization_trustor_audits

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
	Description = utils.FormatUsageDescription("orgauthorization_trustor_audits", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustor/audits", )
	orgauthorization_trustor_auditsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_trustor_audits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_trustor_auditsCmd)
}

func Cmdorgauthorization_trustor_audits() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(createCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(createCmd.Flags(), "string", "sortBy", "timestamp", "Sort by")
	utils.AddFlag(createCmd.Flags(), "string", "sortOrder", "descending", "Sort order")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/orgauthorization/trustor/audits", utils.FormatPermissions([]string{ "authorization:audit:view",  }), utils.GenerateDevCentreLink("POST", "Organization Authorization", "/api/v2/orgauthorization/trustor/audits")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Values to scope the request.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TrustorAuditQueryRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AuditQueryResponse"
      }
    }
  }
}`)
	orgauthorization_trustor_auditsCmd.AddCommand(createCmd)
	return orgauthorization_trustor_auditsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Get Org Trustor Audits",
	Long:  "Get Org Trustor Audits",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Trustorauditqueryrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustor/audits"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
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
