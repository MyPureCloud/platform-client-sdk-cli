package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontablerowexecutionoutputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontablerowexecutionoutputDud struct { 
    


    


    

}

// Decisiontablerowexecutionoutput
type Decisiontablerowexecutionoutput struct { 
    // RowId - Unique rule identifier.
    RowId string `json:"rowId"`


    // RowIndex - Unique rule identifier.
    RowIndex int `json:"rowIndex"`


    // Outputs - The JSON output produced by this rule. Valid according to the execution output contract. In the case of enum decision table output columns, the enum options key will be provided as the value, not the enum options label as this can be changed. For business rules queue columns both “queue” and “id” keys will always be returned  regardless of the business rules queue attribute key and these do not change.
    Outputs map[string]interface{} `json:"outputs"`

}

// String returns a JSON representation of the model
func (o *Decisiontablerowexecutionoutput) String() string {
    
    
     o.Outputs = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontablerowexecutionoutput) MarshalJSON() ([]byte, error) {
    type Alias Decisiontablerowexecutionoutput

    if DecisiontablerowexecutionoutputMarshalled {
        return []byte("{}"), nil
    }
    DecisiontablerowexecutionoutputMarshalled = true

    return json.Marshal(&struct {
        
        RowId string `json:"rowId"`
        
        RowIndex int `json:"rowIndex"`
        
        Outputs map[string]interface{} `json:"outputs"`
        *Alias
    }{

        


        


        
        Outputs: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

