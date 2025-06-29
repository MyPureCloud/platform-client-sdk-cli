package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableexecutionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableexecutionrequestDud struct { 
    

}

// Decisiontableexecutionrequest
type Decisiontableexecutionrequest struct { 
    // Inputs - The JSON input data for executing the decision table. Must be valid according to the execution input contract defined on the table. In the case of enum decision table columns the enum options key should be provided as the value, not the enum options label as this can be changed. For business rules queue columns both “queue” and “id” keys will be required regardless of the queue attribute key and these do not change.
    Inputs map[string]interface{} `json:"inputs"`

}

// String returns a JSON representation of the model
func (o *Decisiontableexecutionrequest) String() string {
     o.Inputs = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableexecutionrequest) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableexecutionrequest

    if DecisiontableexecutionrequestMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableexecutionrequestMarshalled = true

    return json.Marshal(&struct {
        
        Inputs map[string]interface{} `json:"inputs"`
        *Alias
    }{

        
        Inputs: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

