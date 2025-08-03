package externalcontacts_contacts_journey_segments

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
	Description = utils.FormatUsageDescription("externalcontacts_contacts_journey_segments", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/{contactId}/journey/segments", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/{contactId}/journey/segments", )
	externalcontacts_contacts_journey_segmentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_contacts_journey_segments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_contacts_journey_segmentsCmd)
}

func Cmdexternalcontacts_contacts_journey_segments() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/externalcontacts/contacts/{contactId}/journey/segments", utils.FormatPermissions([]string{ "externalContacts:segmentAssignment:add", "externalContacts:segmentAssignment:delete",  }), utils.GenerateDevCentreLink("POST", "External Contacts", "/api/v2/externalcontacts/contacts/{contactId}/journey/segments")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UpdateSegmentAssignmentRequest"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Request completed successfully or was partially successful.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UpdateSegmentAssignmentResponse"
      }
    }
  }
}`)
	externalcontacts_contacts_journey_segmentsCmd.AddCommand(createCmd)

	utils.AddFlag(listCmd.Flags(), "bool", "includeMerged", "", "Indicates whether to return segment assignments from all external contacts in the merge-set of the given one.")
	utils.AddFlag(listCmd.Flags(), "int", "limit", "", "Number of entities to return. Default of 25, maximum of 500.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/contacts/{contactId}/journey/segments", utils.FormatPermissions([]string{ "externalContacts:segmentAssignment:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/contacts/{contactId}/journey/segments")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SegmentAssignmentListing"
      }
    }
  }
}`)
	externalcontacts_contacts_journey_segmentsCmd.AddCommand(listCmd)
	return externalcontacts_contacts_journey_segmentsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [contactId]",
	Short: "Assign/Unassign up to 10 segments to/from an external contact or, if a segment is already assigned, update the expiry date of the segment assignment. Any unprocessed segment assignments are returned in the body for the client to retry, in the event of a partial success.",
	Long:  "Assign/Unassign up to 10 segments to/from an external contact or, if a segment is already assigned, update the expiry date of the segment assignment. Any unprocessed segment assignments are returned in the body for the client to retry, in the event of a partial success.",
	Args:  utils.DetermineArgs([]string{ "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Updatesegmentassignmentrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}/journey/segments"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list [contactId]",
	Short: "Retrieve segment assignments by external contact ID.",
	Long:  "Retrieve segment assignments by external contact ID.",
	Args:  utils.DetermineArgs([]string{ "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}/journey/segments"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)

		includeMerged := utils.GetFlag(cmd.Flags(), "bool", "includeMerged")
		if includeMerged != "" {
			queryParams["includeMerged"] = includeMerged
		}
		limit := utils.GetFlag(cmd.Flags(), "int", "limit")
		if limit != "" {
			queryParams["limit"] = limit
		}
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
