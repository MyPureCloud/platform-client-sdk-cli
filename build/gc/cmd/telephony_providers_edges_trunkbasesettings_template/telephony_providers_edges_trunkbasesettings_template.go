package telephony_providers_edges_trunkbasesettings_template

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_trunkbasesettings_template", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/trunkbasesettings/template", )
	telephony_providers_edges_trunkbasesettings_templateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_trunkbasesettings_template"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_trunkbasesettings_templateCmd)
}

func Cmdtelephony_providers_edges_trunkbasesettings_template() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "trunkMetabaseId", "", "The id of a metabase object upon which to base this Trunk Base Settings - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/trunkbasesettings/template", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/trunkbasesettings/template")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("trunkMetabaseId")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TrunkBase"
  }
}`)
	telephony_providers_edges_trunkbasesettings_templateCmd.AddCommand(getCmd)
	
	return telephony_providers_edges_trunkbasesettings_templateCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a Trunk Base Settings instance template from a given make and model. This object can then be modified and saved as a new Trunk Base Settings instance",
	Long:  "Get a Trunk Base Settings instance template from a given make and model. This object can then be modified and saved as a new Trunk Base Settings instance",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/trunkbasesettings/template"


		trunkMetabaseId := utils.GetFlag(cmd.Flags(), "string", "trunkMetabaseId")
		if trunkMetabaseId != "" {
			queryParams["trunkMetabaseId"] = trunkMetabaseId
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
