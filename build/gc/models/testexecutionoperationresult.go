package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TestexecutionoperationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TestexecutionoperationresultDud struct { 
    


    


    


    


    

}

// Testexecutionoperationresult
type Testexecutionoperationresult struct { 
    // Step - The step number to indicate the order in which the operation was performed
    Step int `json:"step"`


    // Name - Name of the operation performed
    Name string `json:"name"`


    // Success - Indicated whether or not the operation was successful
    Success bool `json:"success"`


    // Result - The result of the operation
    Result interface{} `json:"result"`


    // VarError - Error that occurred during the operation
    VarError Errorbody `json:"error"`

}

// String returns a JSON representation of the model
func (o *Testexecutionoperationresult) String() string {
    
    
    
     o.Result = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testexecutionoperationresult) MarshalJSON() ([]byte, error) {
    type Alias Testexecutionoperationresult

    if TestexecutionoperationresultMarshalled {
        return []byte("{}"), nil
    }
    TestexecutionoperationresultMarshalled = true

    return json.Marshal(&struct {
        
        Step int `json:"step"`
        
        Name string `json:"name"`
        
        Success bool `json:"success"`
        
        Result interface{} `json:"result"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        


        


        


        
        Result: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}

