package orgauthorization_pairings

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
	Description = utils.FormatUsageDescription("orgauthorization_pairings", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/pairings", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/pairings", )
	orgauthorization_pairingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_pairings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_pairingsCmd)
}

func Cmdorgauthorization_pairings() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/orgauthorization/pairings", utils.FormatPermissions([]string{ "authorization:orgTrustee:add",  }), utils.GenerateDevCentreLink("POST", "Organization Authorization", "/api/v2/orgauthorization/pairings")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Pairing Info",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/TrustRequestCreate"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TrustRequest"
  }
}`)
	orgauthorization_pairingsCmd.AddCommand(createCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/orgauthorization/pairings/{pairingId}", utils.FormatPermissions([]string{ "authorization:orgTrustee:view", "authorization:orgTrustor:view",  }), utils.GenerateDevCentreLink("GET", "Organization Authorization", "/api/v2/orgauthorization/pairings/{pairingId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TrustRequest"
  }
}`)
	orgauthorization_pairingsCmd.AddCommand(getCmd)
	
	return orgauthorization_pairingsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A pairing id is created by the trustee and given to the trustor to create a trust.",
	Long:  "A pairing id is created by the trustee and given to the trustor to create a trust.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Trustrequestcreate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/pairings"

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
var getCmd = &cobra.Command{
	Use:   "get [pairingId]",
	Short: "Get Pairing Info",
	Long:  "Get Pairing Info",
	Args:  utils.DetermineArgs([]string{ "pairingId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/pairings/{pairingId}"
		pairingId, args := args[0], args[1:]
		path = strings.Replace(path, "{pairingId}", fmt.Sprintf("%v", pairingId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "get"
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
