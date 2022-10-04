package organizations_features

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
	Description = utils.FormatUsageDescription("organizations_features", "SWAGGER_OVERRIDE_/api/v2/organizations/features", )
	organizations_featuresCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("organizations_features"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(organizations_featuresCmd)
}

func Cmdorganizations_features() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/organizations/features/{featureName}", utils.FormatPermissions([]string{ "directory:organization:admin",  }), utils.GenerateDevCentreLink("PATCH", "Organization", "/api/v2/organizations/features/{featureName}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "New state of feature",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/FeatureState"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/OrganizationFeatures"
      }
    }
  }
}`)
	organizations_featuresCmd.AddCommand(updateCmd)
	return organizations_featuresCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [featureName]",
	Short: "Update organization",
	Long:  "Update organization",
	Args:  utils.DetermineArgs([]string{ "featureName", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Featurestate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/organizations/features/{featureName}"
		featureName, args := args[0], args[1:]
		path = strings.Replace(path, "{featureName}", fmt.Sprintf("%v", featureName), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PATCH"
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
