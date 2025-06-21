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


    // Outputs - The JSON output produced by this rule. Valid according to the execution output contract.
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

