package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VariablevalidationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VariablevalidationDud struct { 
    

}

// Variablevalidation
type Variablevalidation struct { 
    // AdditionalProperties
    AdditionalProperties map[string]interface{} `json:"additionalProperties"`

}

// String returns a JSON representation of the model
func (o *Variablevalidation) String() string {
     o.AdditionalProperties = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Variablevalidation) MarshalJSON() ([]byte, error) {
    type Alias Variablevalidation

    if VariablevalidationMarshalled {
        return []byte("{}"), nil
    }
    VariablevalidationMarshalled = true

    return json.Marshal(&struct {
        
        AdditionalProperties map[string]interface{} `json:"additionalProperties"`
        *Alias
    }{

        
        AdditionalProperties: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

