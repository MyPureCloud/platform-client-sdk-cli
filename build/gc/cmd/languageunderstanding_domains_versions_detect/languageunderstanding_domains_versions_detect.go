package languageunderstanding_domains_versions_detect

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
	Description = utils.FormatUsageDescription("languageunderstanding_domains_versions_detect", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/detect", )
	languageunderstanding_domains_versions_detectCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding_domains_versions_detect"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstanding_domains_versions_detectCmd)
}

func Cmdlanguageunderstanding_domains_versions_detect() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/detect", utils.FormatPermissions([]string{ "languageUnderstanding:nluDomainVersion:view", "dialog:botVersion:view",  }), utils.GenerateDevCentreLink("POST", "Language Understanding", "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/detect")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "The input data to perform detection on.",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/NluDetectionRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Requested NLU detection for the given input using the specified NLU domain version",
  "schema" : {
    "$ref" : "#/definitions/NluDetectionResponse"
  }
}`)
	languageunderstanding_domains_versions_detectCmd.AddCommand(createCmd)
	
	return languageunderstanding_domains_versions_detectCmd
}

var createCmd = &cobra.Command{
	Use:   "create [domainId] [domainVersionId]",
	Short: "Detect intent, entities, etc. in the submitted text using the specified NLU domain version.",
	Long:  "Detect intent, entities, etc. in the submitted text using the specified NLU domain version.",
	Args:  utils.DetermineArgs([]string{ "domainId", "domainVersionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Nludetectionrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/detect"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
		domainVersionId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)


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
