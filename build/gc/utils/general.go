package utils

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"sigs.k8s.io/yaml"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/data_format"
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
	case "time.Time":
		fallthrough
	case "string":
		flags.String(name, "", usage)
		break
	case "int":
		flags.String(name, value, usage)
		break
	case "float32":
		flags.String(name, value, usage)
		break
	case "float64":
		flags.String(name, value, usage)
		break
	default:
		logger.Fatal("Unknown parameter type. Support must be added for it: ", paramType)
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
		flags.StringP("file", "f", "", "File name containing the JSON body")
		flags.BoolP("printrequestbody", "b", false, "Print the request body format of the API.")
		flags.StringP("directory", "d", "", "Directory path with files containing request bodies")
	}
}

func AddPaginateFlagsIfListingResponse(flags *pflag.FlagSet, method, jsonSchema string) {
	if method == http.MethodGet && strings.Contains(jsonSchema, "SWAGGER_OVERRIDE_list") {
		flags.BoolP("autopaginate", "a", false, "Automatically paginate through the results stripping page information")
		flags.BoolP("stream", "s", false, "Paginate and stream data as it is being processed leaving page information intact")
		flags.String("filtercondition", "", "Filter list command output based on a given condition or regular expression")
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
	case "time.Time":
		fallthrough
	case "int":
		fallthrough
	case "float32":
		fallthrough
	case "float64":
		fallthrough
	case "string":
		flag, _ = flags.GetString(name)
		break
	default:
		logger.Fatal("Unknown parameter type. Support must be added for it.", paramType)
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

	// Some command names are separated by underscores. We only want the last name
	name := strings.Split(messages[0], "_")
	message := name[len(name)-1]
	message = strings.Replace(message, "testfile", "test", -1)
	message = strings.Replace(message, "documentationfile", "documentation", -1)
	if len(messages) == 1 {
		return message
	}

	// Add a description if it was specified in the resource definition
	const SwaggerOverride = "SWAGGER_OVERRIDE_"
	var description string
	var descriptions = make([]string, 0)
	for _, description = range messages {
		if strings.Contains(description, SwaggerOverride) {
			alreadyIncluded := false
			for _, existingDescription := range descriptions {
				if existingDescription == description {
					alreadyIncluded = true
				}
			}
			if !alreadyIncluded {
				descriptions = append(descriptions, description)
			}
		}
	}
	return strings.Replace(strings.Join(descriptions, " "), SwaggerOverride, "", -1)
}

// GenerateCustomDescription determines the description given to a command if its subcommands lead to separate paths
// For example, `gc telephony providers edges trunks` leads to
// /api/v2/telephony/providers/edges/trunks and /api/v2/telephony/providers/edges/{edgeId}/trunks
// so it will return the description "/api/v2/telephony/providers/edges/trunks /api/v2/telephony/providers/edges/{edgeId}/trunks"
func GenerateCustomDescription(description string, subcommandDescriptions ...string) string {
	// Don't do anything if there is only one subcommand
	if len(subcommandDescriptions) == 1 {
		return description
	}

	newSubcommandDescriptions := make([]string, 0)
	for _, subcommandDescription := range subcommandDescriptions {
		if !strings.Contains(subcommandDescription, " ") {
			newSubcommandDescriptions = append(newSubcommandDescriptions, subcommandDescription)
		} else {
			// If one of the subcommands leads to separate paths they must be split up
			for _, path := range strings.Split(subcommandDescription, " ") {
				newSubcommandDescriptions = append(newSubcommandDescriptions, path)
			}
		}
	}

	paths := make([]string, 0)
	trailingPathRegex := regexp.MustCompile(`\/.[A-Za-z0-9]{0,}(\/*)$`)
	trailingPathParamRegex := regexp.MustCompile(`\/{[A-Za-z0-9]{0,}}$`)
	for _, subcommandDescription := range newSubcommandDescriptions {
		subcommandDescription = trailingPathRegex.ReplaceAllString(subcommandDescription, "")
		for ok := true; ok; ok = strings.HasSuffix(subcommandDescription, "}") {
			subcommandDescription = trailingPathParamRegex.ReplaceAllString(subcommandDescription, "")
		}
		if len(strings.Split(subcommandDescription, "/")) == 4 {
			return subcommandDescription
		}
		alreadyIncluded := false
		for _, existingString := range paths {
			if existingString == subcommandDescription {
				alreadyIncluded = true
			}
		}
		if !alreadyIncluded {
			paths = append(paths, subcommandDescription)
		}
	}

	if len(paths) == 1 {
		return description
	}

	return strings.Join(paths, " ")
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

func GenerateDevCentreLink(method, category, path string) string {
	category = strings.ToLower(strings.Replace(category, " ", "", -1))
	path = strings.Replace(path, "/", "-", -1)
	path = strings.Replace(path, "{", "-", -1)
	path = strings.Replace(path, "}", "-", -1)

	linkString := "Documentation:\n"
	linkString += fmt.Sprintf("  https://developer.genesys.cloud/api/rest/v2/%v/#%v%v\n",
		category,
		strings.ToLower(method),
		path)

	return linkString
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

	return convertToJSON(string(inputBuffer.Bytes()))
}

func ConvertFile(fileName string) string {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Unable to open file %s.", fileName), err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	fileContent, _ := ioutil.ReadAll(jsonFile)
	return convertToJSON(string(fileContent))
}

func readDirectory(dirName string) []string {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error reading %s: ", dirName), err)
	}
	if len(files) == 0 {
		logger.Fatal(fmt.Sprintf("Error reading %s: no files in directory\n", dirName))
	}

	var data []string

	for _, file := range files {
		fileName := dirName + file.Name()
		if dirName[len(dirName)-1] != '/' {
			fileName = dirName + "/" + file.Name()
		}
		data = append(data, ConvertFile(fileName))
	}

	return data
}

// ResolveInputData is used to determine where the Put, Patch and Delete Post data should be read from
func ResolveInputData(cmd *cobra.Command) []string {
	fileName, _ := cmd.Flags().GetString("file")
	dirName, _ := cmd.Flags().GetString("directory")
	if fileName != "" {
		return []string{ConvertFile(fileName)}
	}
	if dirName != "" {
		return readDirectory(dirName)
	}
	for _, command := range cmd.Commands() {
		fileName, _ := command.Flags().GetString("file")
		dirName, _ := command.Flags().GetString("directory")
		if fileName != "" {
			return []string{ConvertFile(fileName)}
		}
		if dirName != "" {
			return readDirectory(dirName)
		}
	}

	return []string{ConvertStdInString()}
}

// Generate a random string used as PKCE Code Verifier - length = 43 to 128
func GeneratePKCECodeVerifier(n int) (string, error) {
	if n < 43 || n > 128 {
		return "", fmt.Errorf("Error: PKCE Code Verifier (length) must be between 43 and 128 characters")
	}
	const unreservedCharacters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._~"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(unreservedCharacters))))
		if err != nil {
			return "", err
		}
		ret[i] = unreservedCharacters[num.Int64()]
	}

	return string(ret), nil
}

// Compute Base64Url PKCE Code Challenge from Code Verifier
func ComputePKCECodeChallenge(codeVerifier string) (string, error) {
	if len(codeVerifier) < 43 || len(codeVerifier) > 128 {
		return "", fmt.Errorf("Error: PKCE Code Verifier (length) must be between 43 and 128 characters")
	}
	h := sha256.New()
	h.Write([]byte(codeVerifier))
	codeVerifierHash := h.Sum(nil)
	codeChallenge := base64.StdEncoding.EncodeToString([]byte(codeVerifierHash))
	codeChallenge = strings.Replace(codeChallenge, "+", "-", -1) // 62nd char of encoding
	codeChallenge = strings.Replace(codeChallenge, "/", "_", -1) // 63rd char of encoding
	codeChallenge = (strings.Split(codeChallenge, "="))[0]
	return codeChallenge, nil
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

func convertToJSON(data string) string {
	if strings.EqualFold("yaml", data_format.InputFormat) {
		result, err := yaml.YAMLToJSON([]byte(data))
		if err != nil {
			logger.Fatalf("Error converting YAML to JSON: %v\n", err)
		}
		return string(result)
	}
	return data
}

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
