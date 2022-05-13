package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TestexecutionresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TestexecutionresultDud struct { 
    


    


    


    

}

// Testexecutionresult
type Testexecutionresult struct { 
    // Operations - Execution operations performed as part of the test
    Operations []Testexecutionoperationresult `json:"operations"`


    // VarError - The final error encountered during the test that resulted in test failure
    VarError Errorbody `json:"error"`


    // FinalResult - The final result of the test. This is the response that would be returned during normal action execution
    FinalResult interface{} `json:"finalResult"`


    // Success - Indicates whether or not the test was a success
    Success bool `json:"success"`

}

// String returns a JSON representation of the model
func (o *Testexecutionresult) String() string {
     o.Operations = []Testexecutionoperationresult{{}} 
    
     o.FinalResult = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testexecutionresult) MarshalJSON() ([]byte, error) {
    type Alias Testexecutionresult

    if TestexecutionresultMarshalled {
        return []byte("{}"), nil
    }
    TestexecutionresultMarshalled = true

    return json.Marshal(&struct {
        
        Operations []Testexecutionoperationresult `json:"operations"`
        
        VarError Errorbody `json:"error"`
        
        FinalResult interface{} `json:"finalResult"`
        
        Success bool `json:"success"`
        *Alias
    }{

        
        Operations: []Testexecutionoperationresult{{}},
        


        


        
        FinalResult: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}

