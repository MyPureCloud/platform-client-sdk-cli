package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionturninvocationparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionturninvocationparametersDud struct { 
    


    


    

}

// Guidesessionturninvocationparameters
type Guidesessionturninvocationparameters struct { 
    // Name - The name of the parameter.
    Name string `json:"name"`


    // VarType - The type of the parameter.
    VarType string `json:"type"`


    // Value - The value of the parameter.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Guidesessionturninvocationparameters) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionturninvocationparameters) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionturninvocationparameters

    if GuidesessionturninvocationparametersMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionturninvocationparametersMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

