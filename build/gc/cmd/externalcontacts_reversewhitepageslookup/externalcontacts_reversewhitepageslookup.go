package externalcontacts_reversewhitepageslookup

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
	Description = utils.FormatUsageDescription("externalcontacts_reversewhitepageslookup", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/reversewhitepageslookup", )
	externalcontacts_reversewhitepageslookupCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_reversewhitepageslookup"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_reversewhitepageslookupCmd)
}

func Cmdexternalcontacts_reversewhitepageslookup() *cobra.Command { 
	utils.AddFlag(searchCmd.Flags(), "string", "lookupVal", "", "User supplied value to lookup contacts/externalOrganizations (supports email addresses, e164 phone numbers, Twitter screen names) - REQUIRED")
	utils.AddFlag(searchCmd.Flags(), "[]string", "expand", "", "which field, if any, to expand Valid values: contacts.externalOrganization, externalDataSources")
	searchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", searchCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/reversewhitepageslookup", utils.FormatPermissions([]string{ "externalContacts:contact:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/reversewhitepageslookup")))
	utils.AddFileFlagIfUpsert(searchCmd.Flags(), "GET", ``)
	searchCmd.MarkFlagRequired("lookupVal")
	
	utils.AddPaginateFlagsIfListingResponse(searchCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ReverseWhitepagesLookupResult"
      }
    }
  }
}`)
	externalcontacts_reversewhitepageslookupCmd.AddCommand(searchCmd)
	return externalcontacts_reversewhitepageslookupCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Look up contacts and externalOrganizations based on an attribute. Maximum of 25 values returned.",
	Long:  "Look up contacts and externalOrganizations based on an attribute. Maximum of 25 values returned.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/reversewhitepageslookup"

		lookupVal := utils.GetFlag(cmd.Flags(), "string", "lookupVal")
		if lookupVal != "" {
			queryParams["lookupVal"] = lookupVal
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
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

		const opId = "search"
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
