package businessrules_decisiontables_search

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
	Description = utils.FormatUsageDescription("businessrules_decisiontables_search", "SWAGGER_OVERRIDE_/api/v2/businessrules/decisiontables/search", )
	businessrules_decisiontables_searchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("businessrules_decisiontables_search"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(businessrules_decisiontables_searchCmd)
}

func Cmdbusinessrules_decisiontables_search() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "pageSize", "", "Number of entities to return. Maximum of 100.")
	utils.AddFlag(listCmd.Flags(), "string", "schemaId", "", "Search for decision tables that use the schema with this ID. Cannot be combined with name search. Search results will not be paginated if used.")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Search for decision tables with a name that contains the given search string. Search is case insensitive and will match any table that contains this string in any part of the name. Cannot be combined with schema search. Search results will not be paginated if used.")
	utils.AddFlag(listCmd.Flags(), "bool", "withPublishedVersion", "", "Filters results to only decision tables that have at least one version in Published status")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Fields to expand in response Valid values: ExecutionInputSchema, ExecutionOutputSchema")
	utils.AddFlag(listCmd.Flags(), "[]string", "ids", "", "Decision table IDs to search for")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/businessrules/decisiontables/search", utils.FormatPermissions([]string{ "businessrules:decisionTable:search",  }), utils.GenerateDevCentreLink("GET", "Business Rules", "/api/v2/businessrules/decisiontables/search")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	businessrules_decisiontables_searchCmd.AddCommand(listCmd)
	return businessrules_decisiontables_searchCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Search for decision tables.",
	Long:  "Search for decision tables.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/businessrules/decisiontables/search"

		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		pageSize := utils.GetFlag(cmd.Flags(), "string", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		schemaId := utils.GetFlag(cmd.Flags(), "string", "schemaId")
		if schemaId != "" {
			queryParams["schemaId"] = schemaId
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		withPublishedVersion := utils.GetFlag(cmd.Flags(), "bool", "withPublishedVersion")
		if withPublishedVersion != "" {
			queryParams["withPublishedVersion"] = withPublishedVersion
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		ids := utils.GetFlag(cmd.Flags(), "[]string", "ids")
		if ids != "" {
			queryParams["ids"] = ids
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

		const opId = "list"
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
