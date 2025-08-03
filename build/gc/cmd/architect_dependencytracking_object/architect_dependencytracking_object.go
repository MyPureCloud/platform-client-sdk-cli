package architect_dependencytracking_object

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
	Description = utils.FormatUsageDescription("architect_dependencytracking_object", "SWAGGER_OVERRIDE_/api/v2/architect/dependencytracking/object", )
	architect_dependencytracking_objectCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("architect_dependencytracking_object"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(architect_dependencytracking_objectCmd)
}

func Cmdarchitect_dependencytracking_object() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "id", "", "Object ID - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "string", "version", "", "Object version")
	utils.AddFlag(getCmd.Flags(), "string", "objectType", "", "Object type Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, AUDIOCONNECTORBOT, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, CONVERSATIONCUSTOMATTRIBUTESCHEMA, DATAACTION, DATATABLE, DECISIONTABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, DIGITALBOTCONNECTOR, DIGITALBOTCONNECTORINTEGRATION, DIGITALBOTFLOW, DIVISION, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GRAMMAR, GROUP, GUIDE, GUIDEVERSION, IMAGE, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, KNOWLEDGEBASEDOCUMENT, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OAUTHCLIENT, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SMSPHONENUMBER, STTENGINE, SURVEYFORM, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, UTILIZATIONLABEL, VOICEFLOW, VOICEMAILFLOW, VOICESURVEYFLOW, WIDGET, WORKFLOW, WORKITEMFLOW, WORKTYPE")
	utils.AddFlag(getCmd.Flags(), "bool", "consumedResources", "", "Include resources this item consumes")
	utils.AddFlag(getCmd.Flags(), "bool", "consumingResources", "", "Include resources that consume this item")
	utils.AddFlag(getCmd.Flags(), "[]string", "consumedResourceType", "", "Types of consumed resources to return, if consumed resources are requested Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, AUDIOCONNECTORBOT, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, CONVERSATIONCUSTOMATTRIBUTESCHEMA, DATAACTION, DATATABLE, DECISIONTABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, DIGITALBOTCONNECTOR, DIGITALBOTCONNECTORINTEGRATION, DIGITALBOTFLOW, DIVISION, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GRAMMAR, GROUP, GUIDE, GUIDEVERSION, IMAGE, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, KNOWLEDGEBASEDOCUMENT, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OAUTHCLIENT, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SMSPHONENUMBER, STTENGINE, SURVEYFORM, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, UTILIZATIONLABEL, VOICEFLOW, VOICEMAILFLOW, VOICESURVEYFLOW, WIDGET, WORKFLOW, WORKITEMFLOW, WORKTYPE")
	utils.AddFlag(getCmd.Flags(), "[]string", "consumingResourceType", "", "Types of consuming resources to return, if consuming resources are requested Valid values: ACDLANGUAGE, ACDSKILL, ACDWRAPUPCODE, AUDIOCONNECTORBOT, BOTCONNECTORBOT, BOTCONNECTORINTEGRATION, BOTFLOW, BRIDGEACTION, COMMONMODULEFLOW, COMPOSERSCRIPT, CONTACTLIST, CONVERSATIONCUSTOMATTRIBUTESCHEMA, DATAACTION, DATATABLE, DECISIONTABLE, DIALOGENGINEBOT, DIALOGENGINEBOTVERSION, DIALOGFLOWAGENT, DIALOGFLOWCXAGENT, DIGITALBOTCONNECTOR, DIGITALBOTCONNECTORINTEGRATION, DIGITALBOTFLOW, DIVISION, EMAILROUTE, EMERGENCYGROUP, FLOWACTION, FLOWDATATYPE, FLOWMILESTONE, FLOWOUTCOME, GRAMMAR, GROUP, GUIDE, GUIDEVERSION, IMAGE, INBOUNDCALLFLOW, INBOUNDCHATFLOW, INBOUNDEMAILFLOW, INBOUNDSHORTMESSAGEFLOW, INQUEUECALLFLOW, INQUEUEEMAILFLOW, INQUEUESHORTMESSAGEFLOW, IVRCONFIGURATION, KNOWLEDGEBASE, KNOWLEDGEBASEDOCUMENT, LANGUAGE, LEXBOT, LEXBOTALIAS, LEXV2BOT, LEXV2BOTALIAS, NLUDOMAIN, NUANCEMIXBOT, NUANCEMIXINTEGRATION, OAUTHCLIENT, OUTBOUNDCALLFLOW, QUEUE, RECORDINGPOLICY, RESPONSE, SCHEDULE, SCHEDULEGROUP, SECUREACTION, SECURECALLFLOW, SMSPHONENUMBER, STTENGINE, SURVEYFORM, SURVEYINVITEFLOW, SYSTEMPROMPT, TTSENGINE, TTSVOICE, USER, USERPROMPT, UTILIZATIONLABEL, VOICEFLOW, VOICEMAILFLOW, VOICESURVEYFLOW, WIDGET, WORKFLOW, WORKITEMFLOW, WORKTYPE")
	utils.AddFlag(getCmd.Flags(), "bool", "consumedResourceRequest", "", "Indicate that this is going to look up a consumed resource object")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/architect/dependencytracking/object", utils.FormatPermissions([]string{ "architect:dependencyTracking:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/architect/dependencytracking/object")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("id")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/DependencyObject"
      }
    }
  }
}`)
	architect_dependencytracking_objectCmd.AddCommand(getCmd)
	return architect_dependencytracking_objectCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a Dependency Tracking object",
	Long:  "Get a Dependency Tracking object",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/dependencytracking/object"

		id := utils.GetFlag(cmd.Flags(), "string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		version := utils.GetFlag(cmd.Flags(), "string", "version")
		if version != "" {
			queryParams["version"] = version
		}
		objectType := utils.GetFlag(cmd.Flags(), "string", "objectType")
		if objectType != "" {
			queryParams["objectType"] = objectType
		}
		consumedResources := utils.GetFlag(cmd.Flags(), "bool", "consumedResources")
		if consumedResources != "" {
			queryParams["consumedResources"] = consumedResources
		}
		consumingResources := utils.GetFlag(cmd.Flags(), "bool", "consumingResources")
		if consumingResources != "" {
			queryParams["consumingResources"] = consumingResources
		}
		consumedResourceType := utils.GetFlag(cmd.Flags(), "[]string", "consumedResourceType")
		if consumedResourceType != "" {
			queryParams["consumedResourceType"] = consumedResourceType
		}
		consumingResourceType := utils.GetFlag(cmd.Flags(), "[]string", "consumingResourceType")
		if consumingResourceType != "" {
			queryParams["consumingResourceType"] = consumingResourceType
		}
		consumedResourceRequest := utils.GetFlag(cmd.Flags(), "bool", "consumedResourceRequest")
		if consumedResourceRequest != "" {
			queryParams["consumedResourceRequest"] = consumedResourceRequest
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
