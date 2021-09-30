package speechandtextanalytics_topics_general

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
	Description = utils.FormatUsageDescription("speechandtextanalytics_topics_general", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/topics/general", )
	speechandtextanalytics_topics_generalCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_topics_general"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_topics_generalCmd)
}

func Cmdspeechandtextanalytics_topics_general() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "dialect", "", "The dialect of the general topics, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard Valid values: en-US, es-US, en-AU, en-GB, en-ZA, es-ES, en-IN, fr-FR, fr-CA, it-IT, de-DE, pt-BR")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/speechandtextanalytics/topics/general", utils.FormatPermissions([]string{ "speechAndTextAnalytics:topic:view",  }), utils.GenerateDevCentreLink("GET", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/topics/general")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/GeneralTopicsEntityListing"
  }
}`)
	speechandtextanalytics_topics_generalCmd.AddCommand(getCmd)
	
	return speechandtextanalytics_topics_generalCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the Speech and Text Analytics general topics for a given dialect",
	Long:  "Get the Speech and Text Analytics general topics for a given dialect",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/topics/general"


		dialect := utils.GetFlag(cmd.Flags(), "string", "dialect")
		if dialect != "" {
			queryParams["dialect"] = dialect
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
