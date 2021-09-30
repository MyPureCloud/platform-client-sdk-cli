package flows_datatables_import_jobs

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
	Description = utils.FormatUsageDescription("flows_datatables_import_jobs", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/import/jobs", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/import/jobs", "SWAGGER_OVERRIDE_/api/v2/flows/datatables/{datatableId}/import/jobs", )
	flows_datatables_import_jobsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_datatables_import_jobs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_datatables_import_jobsCmd)
}

func Cmdflows_datatables_import_jobs() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows/datatables/{datatableId}/import/jobs", utils.FormatPermissions([]string{ "architect:datatable:edit",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/flows/datatables/{datatableId}/import/jobs")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "import job information",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/DataTableImportJob"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DataTableImportJob"
  }
}`)
	flows_datatables_import_jobsCmd.AddCommand(createCmd)
	
	getstateinformationCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getstateinformationCmd.UsageTemplate(), "GET", "/api/v2/flows/datatables/{datatableId}/import/jobs/{importJobId}", utils.FormatPermissions([]string{ "architect:datatable:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/datatables/{datatableId}/import/jobs/{importJobId}")))
	utils.AddFileFlagIfUpsert(getstateinformationCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getstateinformationCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DataTableImportJob"
  }
}`)
	flows_datatables_import_jobsCmd.AddCommand(getstateinformationCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/flows/datatables/{datatableId}/import/jobs", utils.FormatPermissions([]string{ "architect:datatable:edit",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/datatables/{datatableId}/import/jobs")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	flows_datatables_import_jobsCmd.AddCommand(listCmd)
	
	return flows_datatables_import_jobsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [datatableId]",
	Short: "Begin an import process for importing rows into a datatable",
	Long:  "Begin an import process for importing rows into a datatable",
	Args:  utils.DetermineArgs([]string{ "datatableId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Datatableimportjob{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/{datatableId}/import/jobs"
		datatableId, args := args[0], args[1:]
		path = strings.Replace(path, "{datatableId}", fmt.Sprintf("%v", datatableId), -1)


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
var getstateinformationCmd = &cobra.Command{
	Use:   "getstateinformation [datatableId] [importJobId]",
	Short: "Returns the state information about an import job",
	Long:  "Returns the state information about an import job",
	Args:  utils.DetermineArgs([]string{ "datatableId", "importJobId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/{datatableId}/import/jobs/{importJobId}"
		datatableId, args := args[0], args[1:]
		path = strings.Replace(path, "{datatableId}", fmt.Sprintf("%v", datatableId), -1)
		importJobId, args := args[0], args[1:]
		path = strings.Replace(path, "{importJobId}", fmt.Sprintf("%v", importJobId), -1)


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
var listCmd = &cobra.Command{
	Use:   "list [datatableId]",
	Short: "Get all recent import jobs",
	Long:  "Get all recent import jobs",
	Args:  utils.DetermineArgs([]string{ "datatableId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/datatables/{datatableId}/import/jobs"
		datatableId, args := args[0], args[1:]
		path = strings.Replace(path, "{datatableId}", fmt.Sprintf("%v", datatableId), -1)


		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
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
