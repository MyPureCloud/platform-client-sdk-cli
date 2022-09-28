package routing_predictors_models_features

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
	Description = utils.FormatUsageDescription("routing_predictors_models_features", "SWAGGER_OVERRIDE_/api/v2/routing/predictors/{predictorId}/models/{modelId}/features", )
	routing_predictors_models_featuresCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_predictors_models_features"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_predictors_models_featuresCmd)
}

func Cmdrouting_predictors_models_features() *cobra.Command { 
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/routing/predictors/{predictorId}/models/{modelId}/features", utils.FormatPermissions([]string{ "routing:predictorModelFeature:view",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/predictors/{predictorId}/models/{modelId}/features")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PredictorModelFeatureListing"
      }
    }
  }
}`)
	routing_predictors_models_featuresCmd.AddCommand(listCmd)
	return routing_predictors_models_featuresCmd
}

var listCmd = &cobra.Command{
	Use:   "list [predictorId] [modelId]",
	Short: "Retrieve Predictor Model Features.",
	Long:  "Retrieve Predictor Model Features.",
	Args:  utils.DetermineArgs([]string{ "predictorId", "modelId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/predictors/{predictorId}/models/{modelId}/features"
		predictorId, args := args[0], args[1:]
		path = strings.Replace(path, "{predictorId}", fmt.Sprintf("%v", predictorId), -1)
		modelId, args := args[0], args[1:]
		path = strings.Replace(path, "{modelId}", fmt.Sprintf("%v", modelId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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

		utils.Render(results)
	},
}
