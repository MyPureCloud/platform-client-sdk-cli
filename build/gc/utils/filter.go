package utils

import (
	"encoding/json"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"regexp"
	"strconv"
	"strings"
)

const (
	TypeInteger string = "int"
	TypeString         = "string"
	TypeBoolean        = "bool"
	TypeArray          = "array"
)

const (
	equalsOperator            string = "=="
	notEqualsOperator                = "!="
	lessThanEqualsOperator           = "<="
	greaterThanEqualsOperator        = ">="
	lessThanOperator                 = "<"
	greaterThanOperator              = ">"
	containsOperator                 = "contains"
	matchOperator                    = "match"
)

var (
	supportedOperators = []string{
		equalsOperator,
		notEqualsOperator,
		lessThanEqualsOperator,
		greaterThanEqualsOperator,
		lessThanOperator,
		greaterThanOperator,
		containsOperator,
		matchOperator,
	}
)

func FilterByCondition(data string, condition string) (string, error) {
	objects, err := getJsonObjectsFromString(data)
	if err != nil {
		return "", err
	}

	// Find operator in condition
	operator := findOperatorInString(condition)
	if operator == "" {
		return "", unrecognizedConditionError(condition)
	}

	path, value := getFieldKeyAndValueFromConditionString(condition, operator)
	keys := getKeysFromJsonFieldPath(path)

	allMatchedObjects, err := findObjectsMatchingCondition(objects, keys, value, operator)
	if err != nil {
		return "", err
	}

	// Convert go data to json byte array
	jsonBytes, err := json.MarshalIndent(allMatchedObjects, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func findObjectsMatchingCondition(returnedObjects []interface{}, keys []string, value string, operator string) ([]interface{}, error) {
	var (
		allMatchedObjects []interface{}
		match             bool
	)

	for _, object := range returnedObjects {
		currentObjectsValue, err := getReferencedValueFromMap(keys, object.(map[string]interface{}), value)
		if err != nil {
			return nil, err
		}
		// Field not found in current object
		if currentObjectsValue == nil {
			continue
		}
		if currentObjectArray, ok := currentObjectsValue.([]interface{}); ok {
			match, err = fieldMatchesValueInArray(currentObjectArray, value, operator)
		} else {
			match, err = fieldMatchesValue(currentObjectsValue, value, operator)
		}
		if err != nil {
			return nil, err
		}
		if match {
			allMatchedObjects = append(allMatchedObjects, object.(map[string]interface{}))
		}
	}

	return allMatchedObjects, nil
}

func getReferencedValueFromMap(keys []string, object map[string]interface{}, value string) (interface{}, error) {
	var (
		lastItem     interface{}
		currentValue interface{}
		currentItem  = object
	)
	for i, k := range keys {
		currentValue = currentItem[k]
		if currentValue == nil {
			logger.Infof("field '%s' not found in current json object", k)
			break
		}
		if i == len(keys)-1 {
			lastItem = currentValue
			break
		}
		if itemMap, ok := currentValue.(map[string]interface{}); ok {
			currentItem = itemMap
			continue
		} else {
			return nil, invalidKeyOrderingError(k)
		}
	}
	if _, ok := lastItem.(map[string]interface{}); ok {
		return nil, incomparableFieldError(keys[len(keys)-1], value)
	}
	return lastItem, nil
}

func fieldMatchesValueInArray(jsonArray []interface{}, cliInputValue string, operator string) (bool, error) {
	if operator != containsOperator {
		return false, invalidOperatorError(TypeArray, operator)
	}
	for _, c := range jsonArray {
		match, err := fieldMatchesValue(c, cliInputValue, equalsOperator)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}

func fieldMatchesValue(jsonValue interface{}, cliInputValue string, operator string) (bool, error) {
	switch t := jsonValue.(type) {
	case string:
		return compareStrings(t, cliInputValue, operator)
	case bool:
		return compareBooleans(t, cliInputValue, operator)
	default:
		return compareAsIntegers(jsonValue, cliInputValue, operator)
	}
}

func compareStrings(jsonValue string, cliInputValue string, operator string) (bool, error) {
	switch operator {
	case equalsOperator:
		return jsonValue == cliInputValue, nil
	case notEqualsOperator:
		return jsonValue != cliInputValue, nil
	case containsOperator:
		return strings.Contains(jsonValue, cliInputValue), nil
	case matchOperator:
		matched, err := regexp.Match(cliInputValue, []byte(jsonValue))
		if err != nil {
			return false, err
		}
		return matched, nil
	default:
		return false, invalidOperatorError(TypeString, operator)
	}
}

func compareAsIntegers(jsonValue interface{}, cliInputValue string, operator string) (bool, error) {
	value1Int, err := strconv.Atoi(fmt.Sprintf("%v", jsonValue))
	if err != nil {
		return false, err
	}
	value2Int, err := strconv.Atoi(cliInputValue)
	if err != nil {
		return false, err
	}
	switch operator {
	case equalsOperator:
		return value1Int == value2Int, nil
	case notEqualsOperator:
		return value1Int != value2Int, nil
	case greaterThanOperator:
		return value1Int > value2Int, nil
	case lessThanOperator:
		return value1Int < value2Int, nil
	case greaterThanEqualsOperator:
		return value1Int >= value2Int, nil
	case lessThanEqualsOperator:
		return value1Int <= value2Int, nil
	default:
		return false, invalidOperatorError(TypeInteger, operator)
	}
}

func compareBooleans(jsonValue bool, cliInputValue string, operator string) (bool, error) {
	if cliInputValue != "true" && cliInputValue != "false" {
		return false, invalidBooleanValue(cliInputValue)
	}
	switch operator {
	case equalsOperator:
		return fmt.Sprintf("%v", jsonValue) == fmt.Sprintf("%v", cliInputValue), nil
	case notEqualsOperator:
		return fmt.Sprintf("%v", jsonValue) != fmt.Sprintf("%v", cliInputValue), nil
	default:
		return false, invalidOperatorError(TypeBoolean, operator)
	}
}

func findOperatorInString(text string) string {
	for _, operator := range supportedOperators {
		if strings.Contains(text, operator) {
			return operator
		}
	}
	return ""
}

func getJsonObjectsFromString(data string) ([]interface{}, error) {
	var (
		jsonData interface{}
		objects  []interface{}
	)
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json data: %v", err)
	}
	if jsonMap, ok := jsonData.(map[string]interface{}); ok {
		if entitiesArray, ok := jsonMap["entities"].([]interface{}); ok {
			objects = entitiesArray
		} else {
			return nil, entitiesArrayNotFoundError()
		}
	} else if jsonArray, ok := jsonData.([]interface{}); ok {
		objects = jsonArray
	}
	return objects, nil
}

func getFieldKeyAndValueFromConditionString(condition string, operator string) (string, string) {
	expressionSplit := strings.Split(condition, operator)
	return expressionSplit[0], strings.TrimSpace(expressionSplit[1])
}

func getKeysFromJsonFieldPath(path string) []string {
	var trimmedKeys []string
	keys := strings.Split(path, ".")
	for _, k := range keys {
		if k != "" && k != " " {
			trimmedKeys = append(trimmedKeys, strings.TrimSpace(k))
		}
	}
	return trimmedKeys
}

func invalidOperatorError(valueType string, operator string) error {
	return fmt.Errorf("invalid operator for %s comparison: %s\n", valueType, operator)
}

func entitiesArrayNotFoundError() error {
	return fmt.Errorf("could not locate entites array inside response object")
}

func unrecognizedConditionError(condition string) error {
	return fmt.Errorf("could not understand condition: %s\n", condition)
}

func incomparableFieldError(key string, value string) error {
	return fmt.Errorf("key '%s' is incomparable to value: '%s'", key, value)
}

func invalidKeyOrderingError(key string) error {
	return fmt.Errorf("the value inside field '%s' is a non-map, but it is not the last key in the provided path", key)
}

func invalidBooleanValue(value string) error {
	return fmt.Errorf("'%s' is not boolean. Valid options: true, false", value)
}
