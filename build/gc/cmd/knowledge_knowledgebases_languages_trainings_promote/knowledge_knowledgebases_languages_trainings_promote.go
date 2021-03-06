package knowledge_knowledgebases_languages_trainings_promote

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
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_languages_trainings_promote", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}/promote", )
	knowledge_knowledgebases_languages_trainings_promoteCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_languages_trainings_promote"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_languages_trainings_promoteCmd)
}

func Cmdknowledge_knowledgebases_languages_trainings_promote() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}/promote", utils.FormatPermissions([]string{ "knowledge:training:edit",  }), utils.GenerateDevCentreLink("POST", "Knowledge", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}/promote")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/KnowledgeTraining&quot;
  }
}`)
	knowledge_knowledgebases_languages_trainings_promoteCmd.AddCommand(createCmd)
	
	return knowledge_knowledgebases_languages_trainings_promoteCmd
}

var createCmd = &cobra.Command{
	Use:   "create [knowledgeBaseId] [languageCode] [trainingId]",
	Short: "Promote trained documents from draft state to active.",
	Long:  "Promote trained documents from draft state to active.",
	Args:  utils.DetermineArgs([]string{ "knowledgeBaseId", "languageCode", "trainingId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}/promote"
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
		trainingId, args := args[0], args[1:]
		path = strings.Replace(path, "{trainingId}", fmt.Sprintf("%v", trainingId), -1)

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
