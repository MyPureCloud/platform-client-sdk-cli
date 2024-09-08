package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RuleconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RuleconfigDud struct { 
    


    


    

}

// Ruleconfig
type Ruleconfig struct { 
    // Id - Rule ID.
    Id string `json:"id"`


    // Enabled - Rule enabled.
    Enabled bool `json:"enabled"`


    // Rule - Rule configuration.
    Rule Copilotrule `json:"rule"`

}

// String returns a JSON representation of the model
func (o *Ruleconfig) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ruleconfig) MarshalJSON() ([]byte, error) {
    type Alias Ruleconfig

    if RuleconfigMarshalled {
        return []byte("{}"), nil
    }
    RuleconfigMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Enabled bool `json:"enabled"`
        
        Rule Copilotrule `json:"rule"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

