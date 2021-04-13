package contacts_contactlists

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
	contacts_contactlistsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contacts_contactlists"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud contacts_contactlists", "SWAGGER_OVERRIDE_outbound contacts", "", "", "SWAGGER_OVERRIDE_outbound contacts", "", ),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud contacts_contactlists`, "SWAGGER_OVERRIDE_outbound contacts", "", "", "SWAGGER_OVERRIDE_outbound contacts", "", ),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contacts_contactlistsCmd)
}

func Cmdcontacts_contactlists() *cobra.Command { 
	utils.AddFlag(addCmd.Flags(), "bool", "priority", "", "Contact priority. True means the contact(s) will be dialed next; false means the contact will go to the end of the contact queue.")
	utils.AddFlag(addCmd.Flags(), "bool", "clearSystemData", "", "Clear system data. True means the system columns (attempts, callable status, etc) stored on the contact will be cleared if the contact already exists; false means they won`t.")
	utils.AddFlag(addCmd.Flags(), "bool", "doNotQueue", "", "Do not queue. True means that updated contacts will not have their positions in the queue altered, so contacts that have already been dialed will not be redialed. For new contacts, this parameter has no effect; False means that updated contacts will be re-queued, according to the `priority` parameter.")
	addCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", addCmd.UsageTemplate(), "POST", "/api/v2/outbound/contactlists/{contactListId}/contacts", utils.FormatPermissions([]string{ "outbound:contact:add",  })))
	utils.AddFileFlagIfUpsert(addCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Contact&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/WritableDialerContact&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(addCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/DialerContact&quot;
    }
  }
}`)
	contacts_contactlistsCmd.AddCommand(addCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}", utils.FormatPermissions([]string{ "outbound:contact:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Operation was successful.&quot;
}`)
	contacts_contactlistsCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}", utils.FormatPermissions([]string{ "outbound:contact:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DialerContact&quot;
  }
}`)
	contacts_contactlistsCmd.AddCommand(getCmd)
	
	utils.AddFlag(removeCmd.Flags(), "[]string", "contactIds", "", "ContactIds to delete. - REQUIRED")
	removeCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", removeCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/contactlists/{contactListId}/contacts", utils.FormatPermissions([]string{ "outbound:contact:delete",  })))
	utils.AddFileFlagIfUpsert(removeCmd.Flags(), "DELETE", ``)
	removeCmd.MarkFlagRequired("contactIds")
	utils.AddPaginateFlagsIfListingResponse(removeCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Contacts Deleted.&quot;
}`)
	contacts_contactlistsCmd.AddCommand(removeCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}", utils.FormatPermissions([]string{ "outbound:contact:edit",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Contact&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DialerContact&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DialerContact&quot;
  }
}`)
	contacts_contactlistsCmd.AddCommand(updateCmd)
	
	return contacts_contactlistsCmd
}

var addCmd = &cobra.Command{
	Use:   "add [contactListId]",
	Short: "Add contacts to a contact list.",
	Long:  `Add contacts to a contact list.`,
	Args:  utils.DetermineArgs([]string{ "contactListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/{contactListId}/contacts"
		contactListId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactListId}", fmt.Sprintf("%v", contactListId), -1)

		priority := utils.GetFlag(cmd.Flags(), "bool", "priority")
		if priority != "" {
			queryParams["priority"] = priority
		}
		clearSystemData := utils.GetFlag(cmd.Flags(), "bool", "clearSystemData")
		if clearSystemData != "" {
			queryParams["clearSystemData"] = clearSystemData
		}
		doNotQueue := utils.GetFlag(cmd.Flags(), "bool", "doNotQueue")
		if doNotQueue != "" {
			queryParams["doNotQueue"] = doNotQueue
		}
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
	Use:   "delete [contactListId] [contactId]",
	Short: "Delete a contact.",
	Long:  `Delete a contact.`,
	Args:  utils.DetermineArgs([]string{ "contactListId", "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}"
		contactListId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactListId}", fmt.Sprintf("%v", contactListId), -1)
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
	Use:   "get [contactListId] [contactId]",
	Short: "Get a contact.",
	Long:  `Get a contact.`,
	Args:  utils.DetermineArgs([]string{ "contactListId", "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}"
		contactListId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactListId}", fmt.Sprintf("%v", contactListId), -1)
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
var removeCmd = &cobra.Command{
	Use:   "remove [contactListId]",
	Short: "Delete contacts from a contact list.",
	Long:  `Delete contacts from a contact list.`,
	Args:  utils.DetermineArgs([]string{ "contactListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/{contactListId}/contacts"
		contactListId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactListId}", fmt.Sprintf("%v", contactListId), -1)

		contactIds := utils.GetFlag(cmd.Flags(), "[]string", "contactIds")
		if contactIds != "" {
			queryParams["contactIds"] = contactIds
		}
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
var updateCmd = &cobra.Command{
	Use:   "update [contactListId] [contactId]",
	Short: "Update a contact.",
	Long:  `Update a contact.`,
	Args:  utils.DetermineArgs([]string{ "contactListId", "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}"
		contactListId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactListId}", fmt.Sprintf("%v", contactListId), -1)
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
