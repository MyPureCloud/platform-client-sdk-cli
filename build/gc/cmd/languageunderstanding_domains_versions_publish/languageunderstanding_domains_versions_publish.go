package languageunderstanding_domains_versions_publish

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
	Description = utils.FormatUsageDescription("languageunderstanding_domains_versions_publish", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/publish", )
	languageunderstanding_domains_versions_publishCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding_domains_versions_publish"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstanding_domains_versions_publishCmd)
}

func Cmdlanguageunderstanding_domains_versions_publish() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/publish", utils.FormatPermissions([]string{ "languageUnderstanding:nluDomainVersion:add", "dialog:botVersion:add",  }), utils.GenerateDevCentreLink("POST", "Language Understanding", "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/publish")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/NluDomainVersion"
      }
    }
  }
}`)
	languageunderstanding_domains_versions_publishCmd.AddCommand(createCmd)
	return languageunderstanding_domains_versions_publishCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [domainId] [domainVersionId]",
	Short: "Publish the draft NLU Domain Version.",
	Long:  "Publish the draft NLU Domain Version.",
	Args:  utils.DetermineArgs([]string{ "domainId", "domainVersionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/publish"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
		domainVersionId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)

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
