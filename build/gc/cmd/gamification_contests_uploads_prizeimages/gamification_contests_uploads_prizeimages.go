package gamification_contests_uploads_prizeimages

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
	Description = utils.FormatUsageDescription("gamification_contests_uploads_prizeimages", "SWAGGER_OVERRIDE_/api/v2/gamification/contests/uploads/prizeimages", )
	gamification_contests_uploads_prizeimagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_contests_uploads_prizeimages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_contests_uploads_prizeimagesCmd)
}

func Cmdgamification_contests_uploads_prizeimages() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/gamification/contests/uploads/prizeimages", utils.FormatPermissions([]string{ "gamification:contestPrizeImage:upload",  }), utils.GenerateDevCentreLink("POST", "Gamification", "/api/v2/gamification/contests/uploads/prizeimages")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "query",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/GamificationContestPrizeImageUploadUrlRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Pre-signed urls successfully created.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UploadUrlResponse"
      }
    }
  }
}`)
	gamification_contests_uploads_prizeimagesCmd.AddCommand(createCmd)
	return gamification_contests_uploads_prizeimagesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Generates pre-signed URL to upload a prize image for gamification contests",
	Long:  "Generates pre-signed URL to upload a prize image for gamification contests",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Gamificationcontestprizeimageuploadurlrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/contests/uploads/prizeimages"

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
