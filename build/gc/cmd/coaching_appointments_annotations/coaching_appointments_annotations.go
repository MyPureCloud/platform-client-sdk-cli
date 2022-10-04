package coaching_appointments_annotations

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
	Description = utils.FormatUsageDescription("coaching_appointments_annotations", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments/{appointmentId}/annotations", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments/{appointmentId}/annotations", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments/{appointmentId}/annotations", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments/{appointmentId}/annotations", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments/{appointmentId}/annotations", )
	coaching_appointments_annotationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("coaching_appointments_annotations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(coaching_appointments_annotationsCmd)
}

func Cmdcoaching_appointments_annotations() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/coaching/appointments/{appointmentId}/annotations", utils.FormatPermissions([]string{ "coaching:annotation:add", "coaching:privateAnnotation:add",  }), utils.GenerateDevCentreLink("POST", "Coaching", "/api/v2/coaching/appointments/{appointmentId}/annotations")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "The annotation to add",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CoachingAnnotationCreateRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Annotation created",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CoachingAnnotation"
      }
    }
  }
}`)
	coaching_appointments_annotationsCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}", utils.FormatPermissions([]string{ "coaching:annotation:delete", "coaching:privateAnnotation:delete",  }), utils.GenerateDevCentreLink("DELETE", "Coaching", "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Annotation deleted",
  "content" : { }
}`)
	coaching_appointments_annotationsCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}", utils.FormatPermissions([]string{ "coaching:annotation:view", "coaching:privateAnnotation:view",  }), utils.GenerateDevCentreLink("GET", "Coaching", "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "Annotation retrieved",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CoachingAnnotation"
      }
    }
  }
}`)
	coaching_appointments_annotationsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/coaching/appointments/{appointmentId}/annotations", utils.FormatPermissions([]string{ "coaching:annotation:view", "coaching:privateAnnotation:view",  }), utils.GenerateDevCentreLink("GET", "Coaching", "/api/v2/coaching/appointments/{appointmentId}/annotations")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "Annotations retrieved",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	coaching_appointments_annotationsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}", utils.FormatPermissions([]string{ "coaching:annotation:edit", "coaching:privateAnnotation:edit",  }), utils.GenerateDevCentreLink("PATCH", "Coaching", "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "The new version of the annotation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CoachingAnnotation"
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
        "$ref" : "#/components/schemas/CoachingAnnotation"
      }
    }
  }
}`)
	coaching_appointments_annotationsCmd.AddCommand(updateCmd)
	return coaching_appointments_annotationsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [appointmentId]",
	Short: "Create a new annotation.",
	Long:  "Create a new annotation.",
	Args:  utils.DetermineArgs([]string{ "appointmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Coachingannotationcreaterequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}/annotations"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)

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
var deleteCmd = &cobra.Command{
	Use:   "delete [appointmentId] [annotationId]",
	Short: "Delete an existing annotation",
	Long:  "Delete an existing annotation",
	Args:  utils.DetermineArgs([]string{ "appointmentId", "annotationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)
		annotationId, args := args[0], args[1:]
		path = strings.Replace(path, "{annotationId}", fmt.Sprintf("%v", annotationId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "delete"
		const httpMethod = "DELETE"
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
var getCmd = &cobra.Command{
	Use:   "get [appointmentId] [annotationId]",
	Short: "Retrieve an annotation.",
	Long:  "Retrieve an annotation.",
	Args:  utils.DetermineArgs([]string{ "appointmentId", "annotationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)
		annotationId, args := args[0], args[1:]
		path = strings.Replace(path, "{annotationId}", fmt.Sprintf("%v", annotationId), -1)

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
	Use:   "list [appointmentId]",
	Short: "Get a list of annotations.",
	Long:  "Get a list of annotations.",
	Args:  utils.DetermineArgs([]string{ "appointmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}/annotations"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)

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
var updateCmd = &cobra.Command{
	Use:   "update [appointmentId] [annotationId]",
	Short: "Update an existing annotation.",
	Long:  "Update an existing annotation.",
	Args:  utils.DetermineArgs([]string{ "appointmentId", "annotationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Coachingannotation{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}/annotations/{annotationId}"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)
		annotationId, args := args[0], args[1:]
		path = strings.Replace(path, "{annotationId}", fmt.Sprintf("%v", annotationId), -1)

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
