package skills
import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
)

var skillsCmd = &cobra.Command{
	Use:   utils.FormatUsageDescription("skills"),
	Short: utils.FormatUsageDescription("Manages Genesys Cloud skills"),
	Long:  utils.FormatUsageDescription(`Manages Genesys Cloud skills`),
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(skillsCmd)
}

func Cmdskills() *cobra.Command { 
	createCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/routing/skills", utils.FormatPermissions([]string{ "routing:skill:manage",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	skillsCmd.AddCommand(createCmd)
	
	deleteCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/routing/skills/{skillId}", utils.FormatPermissions([]string{ "routing:skill:manage",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	skillsCmd.AddCommand(deleteCmd)
	
	getCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/routing/skills/{skillId}", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	skillsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Filter for results that start with this value")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "id")
	listCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/routing/skills", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	skillsCmd.AddCommand(listCmd)
	
	return skillsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Skill",
	Long:  `Create Skill`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/skills"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", "create", urlString, "/api/v2/routing/skills")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [skillId]",
	Short: "Delete Routing Skill",
	Long:  `Delete Routing Skill`,
	Args:  utils.DetermineArgs([]string{ "skillId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/skills/{skillId}"
		skillId, args := args[0], args[1:]
		path = strings.Replace(path, "{skillId}", fmt.Sprintf("%v", skillId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", "delete", urlString, "/api/v2/routing/skills/{skillId}")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [skillId]",
	Short: "Get Routing Skill",
	Long:  `Get Routing Skill`,
	Args:  utils.DetermineArgs([]string{ "skillId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/skills/{skillId}"
		skillId, args := args[0], args[1:]
		path = strings.Replace(path, "{skillId}", fmt.Sprintf("%v", skillId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", "get", urlString, "/api/v2/routing/skills/{skillId}")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
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
	Short: "Get the list of routing skills.",
	Long:  `Get the list of routing skills.`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/skills"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
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

		retryFunc := CommandService.DetermineAction("GET", "list", urlString, "/api/v2/routing/skills")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
