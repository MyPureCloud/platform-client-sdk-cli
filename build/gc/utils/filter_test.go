package utils

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"testing"
)

type OperatorTestCaseStruct struct {
	condition         string
	expectedValues    []string
	notExpectedValues []string
}

func TestFilterByConditionMatchOperator(t *testing.T) {
	var (
		object1Id = uuid.NewString()
		object2Id = uuid.NewString()
		object3Id = uuid.NewString()
		object4Id = uuid.NewString()
		object5Id = uuid.NewString()
		data      = fmt.Sprintf(`
[
	{
		"id": "%s",
		"name": "Kyle"
	},
	{
		"id": "%s",
		"name": "Adam"
	},
	{
		"id": "%s",
		"name": "Marki"
	},
	{
		"id": "%s",
		"name": "Zarbi"
	},
	{
		"id": "%s",
		"carInfo": {
			"name": "Ford Fiesta"
		}
	}
]
`, object1Id, object2Id, object3Id, object4Id, object5Id)

		testCase1 = OperatorTestCaseStruct{"name match K([a-z]+)le", []string{object1Id}, []string{object2Id}}
		testCase2 = OperatorTestCaseStruct{"name match A([a-d]+)a(.+)", []string{object2Id}, []string{object1Id}}
		testCase3 = OperatorTestCaseStruct{"name match ([M-Z]+)ar(.+)i", []string{object3Id, object4Id}, []string{object1Id, object2Id}}
		testCase4 = OperatorTestCaseStruct{"name match ^(.{4})$", []string{object1Id, object2Id}, []string{object3Id, object4Id}}
		testCase5 = OperatorTestCaseStruct{"name match ^Gerry$", []string{"null"}, []string{object1Id, object2Id, object3Id, object4Id}}
		testCase6 = OperatorTestCaseStruct{"carInfo.name match ^Ford ([A-Z]{1})(.+)$", []string{object5Id}, []string{object1Id, object2Id, object3Id, object4Id}}
		testCases = []OperatorTestCaseStruct{testCase1, testCase2, testCase3, testCase4, testCase5, testCase6}
	)

	for _, test := range testCases {
		if err := verifyValueReturnedWithCondition(data, test.condition, test.expectedValues, test.notExpectedValues); err != nil {
			t.Errorf("ERROR: %v", err)
		}
	}
}

func TestFilterByConditionInvalidCondition(t *testing.T) {
	type ErrorTestCaseStruct struct {
		condition     string
		expectedError error
	}

	var (
		data = fmt.Sprintf(`
{
	"entities": [
		{
			"name": "Steve",
			"age": 21
		},
		{
			"name": "Jake",
			"age": 32,
			"dead": false
		},
		{
			"carInfo": {
				"carReg": "10-D-09"
			}
		}
	]
}`)
		testCase1 = ErrorTestCaseStruct{"name=Steve", unrecognizedConditionError("name=Steve")}
		testCase2 = ErrorTestCaseStruct{"name!Steve", unrecognizedConditionError("name!Steve")}
		testCase3 = ErrorTestCaseStruct{"name=!Steve", unrecognizedConditionError("name=!Steve")}
		testCase4 = ErrorTestCaseStruct{"name = Steve", unrecognizedConditionError("name = Steve")}
		testCase5 = ErrorTestCaseStruct{"name $ Steve", unrecognizedConditionError("name $ Steve")}
		testCase6 = ErrorTestCaseStruct{"carInfo==foobar", incomparableFieldError("carInfo", "foobar")}
		testCase7 = ErrorTestCaseStruct{"carInfo.carReg.mileage==100", invalidKeyOrderingError("carReg")}
		testCase8 = ErrorTestCaseStruct{"dead==untrue", invalidBooleanValue("untrue")}
		testCases = []ErrorTestCaseStruct{testCase1, testCase2, testCase3, testCase4, testCase5, testCase6, testCase7, testCase8}
	)

	for _, test := range testCases {
		err := verifyInvalidOperationError(data, test.condition, test.expectedError)
		if err != nil {
			t.Error(err)
		}
	}

	unsuitableData := fmt.Sprintf(`
{
	"name": "hello"
}
`)
	err := verifyInvalidOperationError(unsuitableData, "name==hello", entitiesArrayNotFoundError())
	if err != nil {
		t.Error(err)
	}
}

func TestFilterByConditionInvalidOperatorError(t *testing.T) {
	type OperationErrorStruct struct {
		condition string
		fieldType string
		operator  string
	}

	var (
		jsonData = fmt.Sprintf(`
[
	{
		"name": "Kyle",
		"version": 3,
		"a": {
			"foo": ["bar"]
		}
	}
]
`)
		testCase1  = OperationErrorStruct{"name<5", TypeString, lessThanOperator}
		testCase2  = OperationErrorStruct{"name>5", TypeString, greaterThanOperator}
		testCase3  = OperationErrorStruct{"name<=5", TypeString, lessThanEqualsOperator}
		testCase4  = OperationErrorStruct{"name>=5", TypeString, greaterThanEqualsOperator}
		testCase5  = OperationErrorStruct{"version contains 5", TypeInteger, containsOperator}
		testCase6  = OperationErrorStruct{"a.foo==bar", TypeArray, equalsOperator}
		testCase7  = OperationErrorStruct{"a.foo!=bar", TypeArray, notEqualsOperator}
		testCase8  = OperationErrorStruct{"a.foo<bar", TypeArray, lessThanOperator}
		testCase9  = OperationErrorStruct{"a.foo>bar", TypeArray, greaterThanOperator}
		testCase10 = OperationErrorStruct{"a.foo>=bar", TypeArray, greaterThanEqualsOperator}
		testCase11 = OperationErrorStruct{"a.foo<=bar", TypeArray, lessThanEqualsOperator}
		testCases  = []OperationErrorStruct{testCase1, testCase2, testCase3, testCase4, testCase5, testCase6, testCase7, testCase8, testCase9, testCase10, testCase11}
	)

	for _, test := range testCases {
		expectedError := invalidOperatorError(test.fieldType, test.operator)
		if err := verifyInvalidOperationError(jsonData, test.condition, expectedError); err != nil {
			t.Errorf("ERROR: %v", err)
		}
	}
}

func TestFilterByConditionContains(t *testing.T) {
	var (
		objId1 = uuid.NewString()
		objId2 = uuid.NewString()
		objId3 = uuid.NewString()

		jsonData = fmt.Sprintf(`
{
	"entities": [
		{
			"id": "%s",
			"address": "Eindhoven, The Netherlands.",
			"letters": ["x", "y", "z"]
		},
		{
			"id": "%s",
			"address": "La Paz, Bolivia.",
			"outerKey": {
				"letters": ["f", "g"],
				"carInfo": {
					"carReg": "1206-DN-63"
				}
			}
		},
		{
			"id": "%s",
			"address": "Rotterdam, The Netherlands.",
			"outerKey": { 
				"innerKey": {
					"letters": ["a", "b", "c"]
				},
				"carInfo": {
					"carReg": "2912-D-98"
				}
			}
		}
	]
}
`, objId1, objId2, objId3)

		testCase1 = OperatorTestCaseStruct{"address contains Eindhoven", []string{objId1}, []string{objId2, objId3}}
		testCase2 = OperatorTestCaseStruct{"address contains La Paz", []string{objId2}, []string{objId1, objId3}}
		testCase3 = OperatorTestCaseStruct{"address contains Nether", []string{objId1, objId3}, []string{objId2}}
		testCase4 = OperatorTestCaseStruct{"letters contains z", []string{objId1}, []string{objId3, objId2}}
		testCase5 = OperatorTestCaseStruct{"outerKey.letters contains g", []string{objId2}, []string{objId3, objId1}}
		testCase6 = OperatorTestCaseStruct{"outerKey.carInfo.carReg contains DN", []string{objId2}, []string{objId3, objId1}}
		testCase7 = OperatorTestCaseStruct{"outerKey.innerKey.letters contains b", []string{objId3}, []string{objId1, objId2}}
		testCase8 = OperatorTestCaseStruct{"outerKey.carInfo.carReg contains 98", []string{objId3}, []string{objId1, objId2}}
		testCases = []OperatorTestCaseStruct{
			testCase1,
			testCase2,
			testCase3,
			testCase4,
			testCase5,
			testCase6,
			testCase7,
			testCase8,
		}
	)

	for _, test := range testCases {
		err := verifyValueReturnedWithCondition(jsonData, test.condition, test.expectedValues, test.notExpectedValues)
		if err != nil {
			t.Errorf("ERROR: %v", err)
		}
	}
}

func TestFilterByConditionGreaterAndLessThanOperations(t *testing.T) {
	var (
		objId1 = uuid.NewString()
		objId2 = uuid.NewString()
		objId3 = uuid.NewString()
		objId4 = uuid.NewString()

		jsonData = fmt.Sprintf(`
[
	{
		"id": "%s",
		"version": 1
	},
	{
		"id": "%s",
		"version": 2
	},
	{
		"id": "%s",
		"version": 3
	},
	{
		"id": "%s",
		"version": 0
	}
]
`, objId1, objId2, objId3, objId4)

		testCase1 = OperatorTestCaseStruct{"version<2", []string{objId1, objId4}, []string{objId2, objId3}}
		testCase2 = OperatorTestCaseStruct{"version <= 2", []string{objId1, objId2, objId4}, []string{objId3}}
		testCase3 = OperatorTestCaseStruct{"version>2", []string{objId3}, []string{objId1, objId2}}
		testCase4 = OperatorTestCaseStruct{"version>=2", []string{objId2, objId3}, []string{objId1}}
		testCase5 = OperatorTestCaseStruct{"version<10", []string{objId1, objId2, objId3, objId4}, []string{}}
		testCases = []OperatorTestCaseStruct{testCase1, testCase2, testCase3, testCase4, testCase5}
	)

	for _, test := range testCases {
		err := verifyValueReturnedWithCondition(jsonData, test.condition, test.expectedValues, test.notExpectedValues)
		if err != nil {
			t.Errorf("ERROR: %v", err)
		}
	}
}

func TestFilterByConditionEquals(t *testing.T) {
	var (
		object1Id = uuid.NewString()
		object2Id = uuid.NewString()
		jsonData  = fmt.Sprintf(`
[
	{
		"id": "%s",
		"name": "Charlie C",
		"active": true,
		"division": {
			"name": "Home"
		},
		"version": 3,
		"skills": ["Sonic Speed", "Invisibility"]
	},
	{
		"id": "%s",
		"name": "Darragh McD",
		"active": false,
		"division": {
			"name": "New"
		},
		"version": 6,
		"skills": ["Sonic Speed", "Invisibility"]
	}
]`, object1Id, object2Id)

		testCase1 = OperatorTestCaseStruct{"active==true", []string{object1Id}, []string{object2Id}}
		testCase2 = OperatorTestCaseStruct{"name==Charlie C", []string{object1Id}, []string{object2Id}}
		testCase3 = OperatorTestCaseStruct{"division.name==Home", []string{object1Id}, []string{object2Id}}
		testCase4 = OperatorTestCaseStruct{"version==3", []string{object1Id}, []string{object2Id}}
		testCase5 = OperatorTestCaseStruct{"active!=true", []string{object2Id}, []string{object1Id}}
		testCase6 = OperatorTestCaseStruct{"name!=Charlie C", []string{object2Id}, []string{object1Id}}
		testCase7 = OperatorTestCaseStruct{"division.name!=Home", []string{object2Id}, []string{object1Id}}
		testCase8 = OperatorTestCaseStruct{"version!=3", []string{object2Id}, []string{object1Id}}
		testCase9 = OperatorTestCaseStruct{"name==Darragh McD", []string{object2Id}, []string{object1Id}}
		testCases = []OperatorTestCaseStruct{testCase1, testCase2, testCase3, testCase4, testCase5, testCase6, testCase7, testCase8, testCase9}
	)

	for _, test := range testCases {
		if err := verifyValueReturnedWithCondition(jsonData, test.condition, test.expectedValues, test.notExpectedValues); err != nil {
			t.Errorf("ERROR: %v", err)
		}
	}
}

func TestFilterByConditionNestedJson(t *testing.T) {
	var (
		object1Id = uuid.NewString()
		object2Id = uuid.NewString()
		object3Id = uuid.NewString()
		data      = fmt.Sprintf(`
{
	"entities": [
		{
			"id": "%s",
			"name": "Alpha",
			"a": {
				"b": "Beta",
				"c": {
					"d": {
						"f": "Gamma",
						"g": ["Kappa"]
					},
					"e": 0
				}
			}
		}, 
		{
			"id": "%s",
			"name": "Zeta",
			"a": {
				"b": "Theta",
				"c": {
					"d": {
						"f": "Epsilon",
						"g": ["Iota"]
					},
					"e": 101
				}
			}
		},
		{
			"id": "%s",
			"name": "Beta",
			"a": {
				"b": "Iota",
				"c": {
					"d": {
						"f": "Theta",
						"g": ["Sigma", "Iota"]
					},
					"e": 43
				}
			}
		}
	]
}
`, object1Id, object2Id, object3Id)
		testCase1 = OperatorTestCaseStruct{"a.c.d.f==Gamma", []string{object1Id}, []string{object2Id}}
		testCase2 = OperatorTestCaseStruct{"a.c.d.f==Epsilon", []string{object2Id}, []string{object1Id}}
		testCase3 = OperatorTestCaseStruct{"a.c.d.g contains Kappa", []string{object1Id}, []string{object2Id}}
		testCase4 = OperatorTestCaseStruct{"a.c.d.g contains Iota", []string{object2Id, object3Id}, []string{object1Id}}
		testCase5 = OperatorTestCaseStruct{"a.c.e<=0", []string{object1Id}, []string{object2Id}}
		testCase6 = OperatorTestCaseStruct{"a.c.e>100", []string{object2Id}, []string{object1Id}}
		testCase7 = OperatorTestCaseStruct{"a.c.e<100", []string{object3Id, object1Id}, []string{object2Id}}
		testCase8 = OperatorTestCaseStruct{"a.b contains eta", []string{object1Id, object2Id}, []string{object3Id}}
		testCases = []OperatorTestCaseStruct{testCase1, testCase2, testCase3, testCase4, testCase5, testCase6, testCase7, testCase8}
	)

	for _, test := range testCases {
		if err := verifyValueReturnedWithCondition(data, test.condition, test.expectedValues, test.notExpectedValues); err != nil {
			t.Error(err)
		}
	}
}

func verifyInvalidOperationError(jsonData string, condition string, expectedError error) error {
	expectedErrorStr := fmt.Sprintf("%v", expectedError)
	err := verifyValueReturnedWithCondition(jsonData, condition, []string{}, []string{})
	if err != nil {
		if fmt.Sprintf("%v", err) != expectedErrorStr {
			return fmt.Errorf("expected error: %s, got: %v", expectedErrorStr, err)
		} else {
			return nil
		}
	} else {
		return fmt.Errorf("expected to receive error %s, got %v", expectedErrorStr, err)
	}
}

func verifyValueReturnedWithCondition(data string, condition string, expectedValues []string, notExpectedValues []string) error {
	result, err := FilterByCondition(data, condition)
	if err != nil {
		return err
	}
	for _, expectedValue := range expectedValues {
		if !strings.Contains(result, expectedValue) {
			return fmt.Errorf("wrong json object returned. Expected to find value: '%s' with condition '%s'\nReceived:\n%s\n", expectedValue, condition, result)
		}
	}
	for _, notExpectedValue := range notExpectedValues {
		if strings.Contains(result, notExpectedValue) {
			return fmt.Errorf("wrong json object returned. Did not expect to find value: '%s' with condition '%s'\nReceived:\n%s\n", notExpectedValue, condition, result)
		}
	}
	return nil
}
