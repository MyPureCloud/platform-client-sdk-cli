package telephony_siptraces

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
	Description = utils.FormatUsageDescription("telephony_siptraces", "SWAGGER_OVERRIDE_/api/v2/telephony/siptraces", )
	telephony_siptracesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_siptraces"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_siptracesCmd)
}

func Cmdtelephony_siptraces() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "callId", "", "unique identification of the placed call")
	utils.AddFlag(getCmd.Flags(), "string", "toUser", "", "User to who the call was placed")
	utils.AddFlag(getCmd.Flags(), "string", "fromUser", "", "user who placed the call")
	utils.AddFlag(getCmd.Flags(), "string", "conversationId", "", "Unique identification of the conversation")
	utils.AddFlag(getCmd.Flags(), "time.Time", "dateStart", "", "Start date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "dateEnd", "", "End date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/siptraces", utils.FormatPermissions([]string{ "telephony:pcap:view",  }), utils.GenerateDevCentreLink("GET", "Telephony", "/api/v2/telephony/siptraces")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("dateStart")
	getCmd.MarkFlagRequired("dateEnd")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SipSearchResult"
      }
    }
  }
}`)
	telephony_siptracesCmd.AddCommand(getCmd)
	return telephony_siptracesCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch SIP metadata",
	Long:  "Fetch SIP metadata",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/telephony/siptraces"

		callId := utils.GetFlag(cmd.Flags(), "string", "callId")
		if callId != "" {
			queryParams["callId"] = callId
		}
		toUser := utils.GetFlag(cmd.Flags(), "string", "toUser")
		if toUser != "" {
			queryParams["toUser"] = toUser
		}
		fromUser := utils.GetFlag(cmd.Flags(), "string", "fromUser")
		if fromUser != "" {
			queryParams["fromUser"] = fromUser
		}
		conversationId := utils.GetFlag(cmd.Flags(), "string", "conversationId")
		if conversationId != "" {
			queryParams["conversationId"] = conversationId
		}
		dateStart := utils.GetFlag(cmd.Flags(), "time.Time", "dateStart")
		if dateStart != "" {
			queryParams["dateStart"] = dateStart
		}
		dateEnd := utils.GetFlag(cmd.Flags(), "time.Time", "dateEnd")
		if dateEnd != "" {
			queryParams["dateEnd"] = dateEnd
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
