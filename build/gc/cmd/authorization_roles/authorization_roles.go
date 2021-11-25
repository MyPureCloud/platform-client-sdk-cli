package authorization_roles

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
	Description = utils.FormatUsageDescription("authorization_roles", "SWAGGER_OVERRIDE_/api/v2/authorization/roles", "SWAGGER_OVERRIDE_/api/v2/authorization/roles", "SWAGGER_OVERRIDE_/api/v2/authorization/roles", "SWAGGER_OVERRIDE_/api/v2/authorization/roles", "SWAGGER_OVERRIDE_/api/v2/authorization/roles", "SWAGGER_OVERRIDE_/api/v2/authorization/roles", "SWAGGER_OVERRIDE_/api/v2/authorization/roles", )
	authorization_rolesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_roles"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_rolesCmd)
}

func Cmdauthorization_roles() *cobra.Command { 
	utils.AddFlag(bulkgrantCmd.Flags(), "string", "subjectType", "PC_USER", "what the type of the subjects are (PC_GROUP, PC_USER or PC_OAUTH_CLIENT)")
	bulkgrantCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", bulkgrantCmd.UsageTemplate(), "POST", "/api/v2/authorization/roles/{roleId}", utils.FormatPermissions([]string{ "authorization:grant:add",  }), utils.GenerateDevCentreLink("POST", "Authorization", "/api/v2/authorization/roles/{roleId}")))
	utils.AddFileFlagIfUpsert(bulkgrantCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Subjects and Divisions",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/SubjectDivisions"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(bulkgrantCmd.Flags(), "POST", `{
  "description" : "Bulk Grants Created"
}`)
	authorization_rolesCmd.AddCommand(bulkgrantCmd)
	
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/authorization/roles", utils.FormatPermissions([]string{ "authorization:role:add",  }), utils.GenerateDevCentreLink("POST", "Authorization", "/api/v2/authorization/roles")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Organization role",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/DomainOrganizationRoleCreate"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DomainOrganizationRole"
  }
}`)
	authorization_rolesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/authorization/roles/{roleId}", utils.FormatPermissions([]string{ "authorization:role:delete",  }), utils.GenerateDevCentreLink("DELETE", "Authorization", "/api/v2/authorization/roles/{roleId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The request could not be understood by the server due to malformed syntax.",
  "schema" : {
    "$ref" : "#/definitions/ErrorBody"
  },
  "x-inin-error-codes" : {
    "bad.request" : "The request could not be understood by the server due to malformed syntax.",
    "response.entity.too.large" : "The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable"
  }
}`)
	authorization_rolesCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. unusedPermissions returns the permissions not used for the role Valid values: unusedPermissions")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/authorization/roles/{roleId}", utils.FormatPermissions([]string{ "authorization:role:view",  }), utils.GenerateDevCentreLink("GET", "Authorization", "/api/v2/authorization/roles/{roleId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DomainOrganizationRole"
  }
}`)
	authorization_rolesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "The total page size requested")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The page number requested")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "variable name requested to sort by")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "variable name requested by expand list")
	utils.AddFlag(listCmd.Flags(), "string", "nextPage", "", "next page token")
	utils.AddFlag(listCmd.Flags(), "string", "previousPage", "", "Previous page token")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "")
	utils.AddFlag(listCmd.Flags(), "[]string", "permission", "", "")
	utils.AddFlag(listCmd.Flags(), "[]string", "defaultRoleId", "", "")
	utils.AddFlag(listCmd.Flags(), "bool", "userCount", "true", "")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "id")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/authorization/roles", utils.FormatPermissions([]string{ "authorization:role:view",  }), utils.GenerateDevCentreLink("GET", "Authorization", "/api/v2/authorization/roles")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	authorization_rolesCmd.AddCommand(listCmd)
	
	patchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", patchCmd.UsageTemplate(), "PATCH", "/api/v2/authorization/roles/{roleId}", utils.FormatPermissions([]string{ "authorization:role:edit",  }), utils.GenerateDevCentreLink("PATCH", "Authorization", "/api/v2/authorization/roles/{roleId}")))
	utils.AddFileFlagIfUpsert(patchCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "description" : "Organization role",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/DomainOrganizationRole"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(patchCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DomainOrganizationRole"
  }
}`)
	authorization_rolesCmd.AddCommand(patchCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/authorization/roles/{roleId}", utils.FormatPermissions([]string{ "authorization:role:edit",  }), utils.GenerateDevCentreLink("PUT", "Authorization", "/api/v2/authorization/roles/{roleId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "Organization role",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/DomainOrganizationRoleUpdate"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DomainOrganizationRole"
  }
}`)
	authorization_rolesCmd.AddCommand(updateCmd)
	
	return authorization_rolesCmd
}

var bulkgrantCmd = &cobra.Command{
	Use:   "bulkgrant [roleId]",
	Short: "Bulk-grant subjects and divisions with an organization role.",
	Long:  "Bulk-grant subjects and divisions with an organization role.",
	Args:  utils.DetermineArgs([]string{ "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Subjectdivisions{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{roleId}"
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)

		subjectType := utils.GetFlag(cmd.Flags(), "string", "subjectType")
		if subjectType != "" {
			queryParams["subjectType"] = subjectType
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "bulkgrant"
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
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an organization role.",
	Long:  "Create an organization role.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Domainorganizationrolecreate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles"

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
var deleteCmd = &cobra.Command{
	Use:   "delete [roleId]",
	Short: "Delete an organization role.",
	Long:  "Delete an organization role.",
	Args:  utils.DetermineArgs([]string{ "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{roleId}"
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [roleId]",
	Short: "Get a single organization role.",
	Long:  "Get a single organization role.",
	Args:  utils.DetermineArgs([]string{ "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{roleId}"
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)

		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of all roles defined for the organization",
	Long:  "Retrieve a list of all roles defined for the organization",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		nextPage := utils.GetFlag(cmd.Flags(), "string", "nextPage")
		if nextPage != "" {
			queryParams["nextPage"] = nextPage
		}
		previousPage := utils.GetFlag(cmd.Flags(), "string", "previousPage")
		if previousPage != "" {
			queryParams["previousPage"] = previousPage
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		permission := utils.GetFlag(cmd.Flags(), "[]string", "permission")
		if permission != "" {
			queryParams["permission"] = permission
		}
		defaultRoleId := utils.GetFlag(cmd.Flags(), "[]string", "defaultRoleId")
		if defaultRoleId != "" {
			queryParams["defaultRoleId"] = defaultRoleId
		}
		userCount := utils.GetFlag(cmd.Flags(), "bool", "userCount")
		if userCount != "" {
			queryParams["userCount"] = userCount
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var patchCmd = &cobra.Command{
	Use:   "patch [roleId]",
	Short: "Patch Organization Role for needsUpdate Field",
	Long:  "Patch Organization Role for needsUpdate Field",
	Args:  utils.DetermineArgs([]string{ "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Domainorganizationrole{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{roleId}"
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "patch"
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [roleId]",
	Short: "Update an organization role.",
	Long:  "Update an organization role.",
	Args:  utils.DetermineArgs([]string{ "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Domainorganizationroleupdate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{roleId}"
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PUT"
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
