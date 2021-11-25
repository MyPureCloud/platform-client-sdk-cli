package groups_members

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
)

func init() {
	note := "Note: The 'version' value from the command input will be ignored and the latest version value will be retrieved from the API instead"
	AddCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n\n%s\n", AddCmd.UsageTemplate(), "POST", "/api/v2/groups/{groupId}/members", utils.FormatPermissions([]string{}), note))

	addCmd = AddCmd
}

type group struct {
	Version int `json:"version"`
}

type addGroupMembersBody struct {
	MemberIds []string `json:"memberIds"`
	Version   int      `json:"version"`
}

var (
	addMembersOperation = models.HandWrittenOperation{
		Path:   "/api/v2/groups/{groupId}/members",
		Method: http.MethodPost,
	}
	getMembersOperation = models.HandWrittenOperation{
		Path:   "/api/v2/groups/{groupId}",
		Method: http.MethodGet,
	}
)

var AddCmd = &cobra.Command{
	Use:   "add [groupId]",
	Short: "Add members",
	Long:  `Add members`,
	Args:  utils.DetermineArgs([]string{"groupId"}),

	Run: func(cmd *cobra.Command, args []string) {
		groupId, args := args[0], args[1:]
		path := strings.Replace(getMembersOperation.Path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

		currentVersion := getGroupVersion(path, cmd)

		inputData := utils.ResolveInputData(cmd)

		// convert inputData to []byte for unmarshalling
		buffer := &bytes.Buffer{}
		gob.NewEncoder(buffer).Encode(inputData)
		byteSlice := buffer.Bytes()

		body := &addGroupMembersBody{}
		err := json.Unmarshal([]byte(byteSlice), body)
		if err != nil {
			logger.Fatal(err)
		}

		body.Version = currentVersion
		bodyString, _ := json.Marshal(body)

		// request body
		data := []string{string(bodyString)}

		path = strings.Replace(addMembersOperation.Path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

		retryFunc := retry.RetryWithData(path, data, CommandService.Post)
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

func getGroupVersion(path string, cmd *cobra.Command) int {
	retryFunc := CommandService.DetermineAction(getMembersOperation.Method, path, cmd, "")
	retryConfig := &retry.RetryConfiguration{
		RetryWaitMin: 5 * time.Second,
		RetryWaitMax: 60 * time.Second,
		RetryMax:     20,
	}
	results, err := retryFunc(retryConfig)
	if err != nil {
		logger.Fatal(err)
	}

	groupResult := group{}
	err = json.Unmarshal([]byte(results), &groupResult)
	if err != nil {
		logger.Fatal(err)
	}

	return groupResult.Version
}
