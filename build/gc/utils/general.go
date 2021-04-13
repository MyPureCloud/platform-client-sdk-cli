package utils

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func AddFlag(flags *pflag.FlagSet, paramType string, name string, value string, usage string) {
	switch paramType {
	case "[]string":
		flags.StringSlice(name, []string{}, usage)
		break
	case "bool":
		usage = fmt.Sprintf("%v %v", usage, "Valid values: true, false")
		fallthrough
	case "string":
		flags.String(name, "", usage)
		break
	case "int":
		intValue, _ := strconv.Atoi(value)
		flags.Int(name, intValue, usage)
		break
	}
}

func AddFileFlagIfUpsert(flags *pflag.FlagSet, method string, jsonSchema string) {
	if jsonSchema == "" {
		return
	}
	switch method {
	case http.MethodPatch:
		fallthrough
	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		flags.StringP("file", "f", "", "File name containing the JSON for creating an object")
	}
}

func AddPaginateFlagsIfListingResponse(flags *pflag.FlagSet, method, jsonSchema string) {
	if method == http.MethodGet && strings.Contains(jsonSchema, "Listing") {
		flags.BoolP("autopaginate", "a", false, "Automatically paginate through the results stripping page information")
		flags.BoolP("stream", "s", false, "Paginate and stream data as it is being processed leaving page information intact")
	}
}

func GetFlag(flags *pflag.FlagSet, paramType string, name string) string {
	flag := ""
	switch paramType {
	case "[]string":
		flags, _ := flags.GetStringSlice(name)
		flag = strings.Join(flags, ",")
		break
	case "bool":
		fallthrough
	case "string":
		flag, _ = flags.GetString(name)
		break
	case "int":
		flagInt, _ := flags.GetInt(name)
		flag = strconv.Itoa(flagInt)
		break
	}
	return flag
}

func FormatUsageDescription(inputs ...string) string {
	messages := make([]string, 0)
	// Only add non-empty strings to the messages slice
	for _, input := range inputs {
		if len(input) != 0 {
			messages = append(messages, input)
		}
	}

	// Some command names are separated by underscores. We only want the first name
	message := strings.Split(messages[0], "_")[0]
	if len(messages) == 1 {
		return message
	}

	// Add a description if it was specified in the resource definition
	const SwaggerOverride = "SWAGGER_OVERRIDE_"
	var description string
	for _, description = range messages {
		if strings.Contains(description, SwaggerOverride) {
			break
		}
	}
	return fmt.Sprintf("Manages Genesys Cloud %s", strings.Replace(description, SwaggerOverride, "", 1))
}

func FormatPermissions(permissions []string) string {
	if len(permissions) == 0 {
		return ""
	}

	permissionString := "\nPermissions:\n"
	for _, permission := range permissions {
		permissionString = fmt.Sprintf("%s  %s\n", permissionString, permission)
	}

	return permissionString
}

func DetermineArgs(args []string) cobra.PositionalArgs {
	validArgs := 0
	for _, arg := range args {
		if arg != "body" {
			validArgs++
		}
	}
	return cobra.ExactArgs(validArgs)
}

func AliasOperationId(operationId string, classVarName string) string {
	return strings.ReplaceAll(operationId, classVarName, "")
}

func ConvertStdInString() string {
	consolescanner := bufio.NewScanner(os.Stdin)
	var inputBuffer bytes.Buffer

	done := make(chan struct{})
	gotText := false

	// by default, bufio.Scanner scans newline-separated lines
	go func() {
		for consolescanner.Scan() {
			input := consolescanner.Text()
			gotText = true
			inputBuffer.WriteString(input)
		}
		close(done)
	}()

	// if no input is read after 1 second then the command will error out
	select {
	case <-done:
	case <-time.After(1000 * time.Millisecond):
		if !gotText {
			logger.Fatal("This command requires a body to be piped to it, or provided from a file via the \"--file\" flag\n")
		}
	}

	return string(inputBuffer.Bytes())
}

func ConvertFileJSON(fileName string) string {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Unable to open file %s.", fileName), err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	fileContent, _ := ioutil.ReadAll(jsonFile)
	return string(fileContent)
}

// ResolveInputData is used to determine where the Put, Patch and Delete Post data should be read from
func ResolveInputData(cmd *cobra.Command) string {
	fileName, _ := cmd.Flags().GetString("file")
	if fileName != "" {
		return ConvertFileJSON(fileName)
	}
	for _, command := range cmd.Commands() {
		fileName, _ := command.Flags().GetString("file")
		if fileName != "" {
			return ConvertFileJSON(fileName)
		}
	}

	return ConvertStdInString()
}

func GenerateGuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		logger.Fatal(err)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func MilliSecondsToNanoSeconds(milliSeconds int64) time.Duration {
	return time.Duration(milliSeconds * 1000 * 1000)
}

func SecondsToNanoSeconds(seconds int) time.Duration {
	return MilliSecondsToNanoSeconds(int64(seconds)) * 1000
}
