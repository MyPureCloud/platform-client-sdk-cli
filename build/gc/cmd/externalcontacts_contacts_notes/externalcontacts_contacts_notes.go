package externalcontacts_contacts_notes

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
	Description = utils.FormatUsageDescription("externalcontacts_contacts_notes", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/{contactId}/notes", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/{contactId}/notes", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/{contactId}/notes", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/{contactId}/notes", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/{contactId}/notes", )
	externalcontacts_contacts_notesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_contacts_notes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_contacts_notesCmd)
}

func Cmdexternalcontacts_contacts_notes() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/externalcontacts/contacts/{contactId}/notes", utils.FormatPermissions([]string{ "externalContacts:contact:edit",  }), utils.GenerateDevCentreLink("POST", "External Contacts", "/api/v2/externalcontacts/contacts/{contactId}/notes")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;ExternalContact&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Note&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Note&quot;
  }
}`)
	externalcontacts_contacts_notesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}", utils.FormatPermissions([]string{ "externalContacts:contact:edit",  }), utils.GenerateDevCentreLink("DELETE", "External Contacts", "/api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Empty&quot;
  }
}`)
	externalcontacts_contacts_notesCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "which fields, if any, to expand Valid values: author, externalDataSources")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}", utils.FormatPermissions([]string{ "externalContacts:contact:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Note&quot;
  }
}`)
	externalcontacts_contacts_notesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "20", "Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be &lt;= 1,000)")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be &lt;= 1,000)")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "", "Sort order")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "which fields, if any, to expand Valid values: author, externalDataSources")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/contacts/{contactId}/notes", utils.FormatPermissions([]string{ "externalContacts:contact:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/contacts/{contactId}/notes")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	externalcontacts_contacts_notesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}", utils.FormatPermissions([]string{ "externalContacts:contact:edit",  }), utils.GenerateDevCentreLink("PUT", "External Contacts", "/api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Note&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Note&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Note&quot;
  }
}`)
	externalcontacts_contacts_notesCmd.AddCommand(updateCmd)
	
	return externalcontacts_contacts_notesCmd
}

var createCmd = &cobra.Command{
	Use:   "create [contactId]",
	Short: "Create a note for an external contact",
	Long:  "Create a note for an external contact",
	Args:  utils.DetermineArgs([]string{ "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}/notes"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)

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
var deleteCmd = &cobra.Command{
	Use:   "delete [contactId] [noteId]",
	Short: "Delete a note for an external contact",
	Long:  "Delete a note for an external contact",
	Args:  utils.DetermineArgs([]string{ "contactId", "noteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)
		noteId, args := args[0], args[1:]
		path = strings.Replace(path, "{noteId}", fmt.Sprintf("%v", noteId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
	Use:   "get [contactId] [noteId]",
	Short: "Fetch a note for an external contact",
	Long:  "Fetch a note for an external contact",
	Args:  utils.DetermineArgs([]string{ "contactId", "noteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)
		noteId, args := args[0], args[1:]
		path = strings.Replace(path, "{noteId}", fmt.Sprintf("%v", noteId), -1)

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
	Use:   "list [contactId]",
	Short: "List notes for an external contact",
	Long:  "List notes for an external contact",
	Args:  utils.DetermineArgs([]string{ "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}/notes"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
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
var updateCmd = &cobra.Command{
	Use:   "update [contactId] [noteId]",
	Short: "Update a note for an external contact",
	Long:  "Update a note for an external contact",
	Args:  utils.DetermineArgs([]string{ "contactId", "noteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)
		noteId, args := args[0], args[1:]
		path = strings.Replace(path, "{noteId}", fmt.Sprintf("%v", noteId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
