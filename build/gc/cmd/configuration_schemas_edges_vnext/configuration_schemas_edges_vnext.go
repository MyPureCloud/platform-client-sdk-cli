package configuration_schemas_edges_vnext

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("configuration_schemas_edges_vnext", "SWAGGER_OVERRIDE_/api/v2/configuration/schemas/edges/vnext/{schemaCategory}", "SWAGGER_OVERRIDE_/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}", "SWAGGER_OVERRIDE_/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}", "SWAGGER_OVERRIDE_/api/v2/configuration/schemas/edges/vnext", "SWAGGER_OVERRIDE_/api/v2/configuration/schemas/edges/vnext", )
	configuration_schemas_edges_vnextCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("configuration_schemas_edges_vnext"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(configuration_schemas_edges_vnextCmd)
}

func Cmdconfiguration_schemas_edges_vnext() *cobra.Command { 
	utils.AddFlag(getcategoryschemasCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(getcategoryschemasCmd.Flags(), "int", "pageNumber", "1", "Page number")
	getcategoryschemasCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getcategoryschemasCmd.UsageTemplate(), "GET", "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getcategoryschemasCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getcategoryschemasCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	configuration_schemas_edges_vnextCmd.AddCommand(getcategoryschemasCmd)
	
	getjsonschemaCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getjsonschemaCmd.UsageTemplate(), "GET", "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getjsonschemaCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getjsonschemaCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Organization&quot;
  }
}`)
	configuration_schemas_edges_vnextCmd.AddCommand(getjsonschemaCmd)
	
	utils.AddFlag(getschemametadataCmd.Flags(), "string", "varType", "", "Type")
	getschemametadataCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getschemametadataCmd.UsageTemplate(), "GET", "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getschemametadataCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getschemametadataCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Organization&quot;
  }
}`)
	configuration_schemas_edges_vnextCmd.AddCommand(getschemametadataCmd)
	
	utils.AddFlag(listcategoryschemasCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listcategoryschemasCmd.Flags(), "int", "pageNumber", "1", "Page number")
	listcategoryschemasCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listcategoryschemasCmd.UsageTemplate(), "GET", "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(listcategoryschemasCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listcategoryschemasCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	configuration_schemas_edges_vnextCmd.AddCommand(listcategoryschemasCmd)
	
	utils.AddFlag(listschemacategoriesCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listschemacategoriesCmd.Flags(), "int", "pageNumber", "1", "Page number")
	listschemacategoriesCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listschemacategoriesCmd.UsageTemplate(), "GET", "/api/v2/configuration/schemas/edges/vnext", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(listschemacategoriesCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listschemacategoriesCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	configuration_schemas_edges_vnextCmd.AddCommand(listschemacategoriesCmd)
	
	return configuration_schemas_edges_vnextCmd
}

var getcategoryschemasCmd = &cobra.Command{
	Use:   "getcategoryschemas [schemaCategory] [schemaType]",
	Short: "List schemas of a specific category (Deprecated)",
	Long:  "List schemas of a specific category (Deprecated)",
	Args:  utils.DetermineArgs([]string{ "schemaCategory", "schemaType", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}"
		schemaCategory, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaCategory}", fmt.Sprintf("%v", schemaCategory), -1)
		schemaType, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaType}", fmt.Sprintf("%v", schemaType), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
var getjsonschemaCmd = &cobra.Command{
	Use:   "getjsonschema [schemaCategory] [schemaType] [schemaId]",
	Short: "Get a json schema (Deprecated)",
	Long:  "Get a json schema (Deprecated)",
	Args:  utils.DetermineArgs([]string{ "schemaCategory", "schemaType", "schemaId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}"
		schemaCategory, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaCategory}", fmt.Sprintf("%v", schemaCategory), -1)
		schemaType, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaType}", fmt.Sprintf("%v", schemaType), -1)
		schemaId, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaId}", fmt.Sprintf("%v", schemaId), -1)

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
var getschemametadataCmd = &cobra.Command{
	Use:   "getschemametadata [schemaCategory] [schemaType] [schemaId] [extensionType] [metadataId]",
	Short: "Get metadata for a schema (Deprecated)",
	Long:  "Get metadata for a schema (Deprecated)",
	Args:  utils.DetermineArgs([]string{ "schemaCategory", "schemaType", "schemaId", "extensionType", "metadataId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}"
		schemaCategory, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaCategory}", fmt.Sprintf("%v", schemaCategory), -1)
		schemaType, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaType}", fmt.Sprintf("%v", schemaType), -1)
		schemaId, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaId}", fmt.Sprintf("%v", schemaId), -1)
		extensionType, args := args[0], args[1:]
		path = strings.Replace(path, "{extensionType}", fmt.Sprintf("%v", extensionType), -1)
		metadataId, args := args[0], args[1:]
		path = strings.Replace(path, "{metadataId}", fmt.Sprintf("%v", metadataId), -1)

		varType := utils.GetFlag(cmd.Flags(), "string", "varType")
		if varType != "" {
			queryParams["varType"] = varType
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
var listcategoryschemasCmd = &cobra.Command{
	Use:   "listcategoryschemas [schemaCategory]",
	Short: "List schemas of a specific category (Deprecated)",
	Long:  "List schemas of a specific category (Deprecated)",
	Args:  utils.DetermineArgs([]string{ "schemaCategory", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}"
		schemaCategory, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaCategory}", fmt.Sprintf("%v", schemaCategory), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
var listschemacategoriesCmd = &cobra.Command{
	Use:   "listschemacategories",
	Short: "Lists available schema categories (Deprecated)",
	Long:  "Lists available schema categories (Deprecated)",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/configuration/schemas/edges/vnext"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
