package scim_groups

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
	Description = utils.FormatUsageDescription("scim_groups", "SWAGGER_OVERRIDE_/api/v2/scim/groups", "SWAGGER_OVERRIDE_/api/v2/scim/groups", "SWAGGER_OVERRIDE_/api/v2/scim/groups", "SWAGGER_OVERRIDE_/api/v2/scim/groups", )
	scim_groupsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scim_groups"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scim_groupsCmd)
}

func Cmdscim_groups() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "[]string", "attributes", "", "Indicates which attributes to include. Returns these attributes and the id, active, and meta attributes. Use attributes to avoid expensive secondary calls for the default attributes. Valid values: id, displayName, members, externalId, meta, meta.version, meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:id, urn:ietf:params:scim:schemas:core:2.0:Group:meta, urn:ietf:params:scim:schemas:core:2.0:Group:meta.version, urn:ietf:params:scim:schemas:core:2.0:Group:meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:displayName, urn:ietf:params:scim:schemas:core:2.0:Group:members, urn:ietf:params:scim:schemas:core:2.0:Group:externalId")
	utils.AddFlag(getCmd.Flags(), "[]string", "excludedAttributes", "", "Indicates which attributes to exclude. Returns the default attributes minus excludedAttributes. Always returns id, active, and meta attributes. Use excludedAttributes to avoid expensive secondary calls for the default attributes. Valid values: id, displayName, members, externalId, meta, meta.version, meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:id, urn:ietf:params:scim:schemas:core:2.0:Group:meta, urn:ietf:params:scim:schemas:core:2.0:Group:meta.version, urn:ietf:params:scim:schemas:core:2.0:Group:meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:displayName, urn:ietf:params:scim:schemas:core:2.0:Group:members, urn:ietf:params:scim:schemas:core:2.0:Group:externalId")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/scim/groups/{groupId}", utils.FormatPermissions([]string{ "directory:group:edit",  }), utils.GenerateDevCentreLink("GET", "SCIM", "/api/v2/scim/groups/{groupId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/scim+json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2Group"
      }
    },
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2Group"
      }
    }
  }
}`)
	scim_groupsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "startIndex", "1", "The 1-based index of the first query result.")
	utils.AddFlag(listCmd.Flags(), "int", "count", "25", "The requested number of items per page. A value of 0 returns totalResults. A page size over 25 may exceed internal resource limits and return a 429 error. For a page size over 25, use the excludedAttributes or attributes query parameters to exclude or only include secondary lookup values such as externalId,  roles, urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User:routingLanguages, or urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User:routingSkills.")
	utils.AddFlag(listCmd.Flags(), "[]string", "attributes", "", "Indicates which attributes to include. Returns these attributes and the id, active, and meta attributes. Use attributes to avoid expensive secondary calls for the default attributes. Valid values: id, displayName, members, externalId, meta, meta.version, meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:id, urn:ietf:params:scim:schemas:core:2.0:Group:meta, urn:ietf:params:scim:schemas:core:2.0:Group:meta.version, urn:ietf:params:scim:schemas:core:2.0:Group:meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:displayName, urn:ietf:params:scim:schemas:core:2.0:Group:members, urn:ietf:params:scim:schemas:core:2.0:Group:externalId")
	utils.AddFlag(listCmd.Flags(), "[]string", "excludedAttributes", "", "Indicates which attributes to exclude. Returns the default attributes minus excludedAttributes. Always returns id, active, and meta attributes. Use excludedAttributes to avoid expensive secondary calls for the default attributes. Valid values: id, displayName, members, externalId, meta, meta.version, meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:id, urn:ietf:params:scim:schemas:core:2.0:Group:meta, urn:ietf:params:scim:schemas:core:2.0:Group:meta.version, urn:ietf:params:scim:schemas:core:2.0:Group:meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:displayName, urn:ietf:params:scim:schemas:core:2.0:Group:members, urn:ietf:params:scim:schemas:core:2.0:Group:externalId")
	utils.AddFlag(listCmd.Flags(), "string", "filter", "", "Filters results. If nothing is specified, returns all groups. Examples of valid values: id eq 5f4bc742-a019-4e38-8e2a-d39d5bc0b0f3, displayname eq Sales.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/scim/groups", utils.FormatPermissions([]string{ "directory:group:edit",  }), utils.GenerateDevCentreLink("GET", "SCIM", "/api/v2/scim/groups")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/scim+json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    },
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	scim_groupsCmd.AddCommand(listCmd)

	replaceCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", replaceCmd.UsageTemplate(), "PUT", "/api/v2/scim/groups/{groupId}", utils.FormatPermissions([]string{ "directory:group:edit",  }), utils.GenerateDevCentreLink("PUT", "SCIM", "/api/v2/scim/groups/{groupId}")))
	utils.AddFileFlagIfUpsert(replaceCmd.Flags(), "PUT", `{
  "description" : "The information used to replace a group.",
  "content" : {
    "application/scim+json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2Group"
      }
    },
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2Group"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(replaceCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/scim+json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2Group"
      }
    },
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2Group"
      }
    }
  }
}`)
	scim_groupsCmd.AddCommand(replaceCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/scim/groups/{groupId}", utils.FormatPermissions([]string{ "directory:group:edit",  }), utils.GenerateDevCentreLink("PATCH", "SCIM", "/api/v2/scim/groups/{groupId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "The information used to modify a group.",
  "content" : {
    "application/scim+json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2PatchRequest"
      }
    },
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2PatchRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/scim+json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2Group"
      }
    },
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ScimV2Group"
      }
    }
  }
}`)
	scim_groupsCmd.AddCommand(updateCmd)
	return scim_groupsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [groupId]",
	Short: "Get a group",
	Long:  "Get a group",
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/scim/groups/{groupId}"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

		attributes := utils.GetFlag(cmd.Flags(), "[]string", "attributes")
		if attributes != "" {
			queryParams["attributes"] = attributes
		}
		excludedAttributes := utils.GetFlag(cmd.Flags(), "[]string", "excludedAttributes")
		if excludedAttributes != "" {
			queryParams["excludedAttributes"] = excludedAttributes
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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of groups",
	Long:  "Get a list of groups",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/scim/groups"

		startIndex := utils.GetFlag(cmd.Flags(), "int", "startIndex")
		if startIndex != "" {
			queryParams["startIndex"] = startIndex
		}
		count := utils.GetFlag(cmd.Flags(), "int", "count")
		if count != "" {
			queryParams["count"] = count
		}
		attributes := utils.GetFlag(cmd.Flags(), "[]string", "attributes")
		if attributes != "" {
			queryParams["attributes"] = attributes
		}
		excludedAttributes := utils.GetFlag(cmd.Flags(), "[]string", "excludedAttributes")
		if excludedAttributes != "" {
			queryParams["excludedAttributes"] = excludedAttributes
		}
		filter := utils.GetFlag(cmd.Flags(), "string", "filter")
		if filter != "" {
			queryParams["filter"] = filter
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
var replaceCmd = &cobra.Command{
	Use:   "replace [groupId]",
	Short: "Replace a group",
	Long:  "Replace a group",
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Scimv2group{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/scim/groups/{groupId}"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

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

		const opId = "replace"
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
	Use:   "update [groupId]",
	Short: "Modify a group",
	Long:  "Modify a group",
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Scimv2patchrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/scim/groups/{groupId}"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

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
