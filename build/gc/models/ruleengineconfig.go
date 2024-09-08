package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RuleengineconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RuleengineconfigDud struct { 
    


    

}

// Ruleengineconfig
type Ruleengineconfig struct { 
    // Rules - List of rules to evaluate by the engine.
    Rules []Ruleconfig `json:"rules"`


    // Fallback - Fallback configuration.
    Fallback Fallback `json:"fallback"`

}

// String returns a JSON representation of the model
func (o *Ruleengineconfig) String() string {
     o.Rules = []Ruleconfig{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ruleengineconfig) MarshalJSON() ([]byte, error) {
    type Alias Ruleengineconfig

    if RuleengineconfigMarshalled {
        return []byte("{}"), nil
    }
    RuleengineconfigMarshalled = true

    return json.Marshal(&struct {
        
        Rules []Ruleconfig `json:"rules"`
        
        Fallback Fallback `json:"fallback"`
        *Alias
    }{

        
        Rules: []Ruleconfig{{}},
        


        

        Alias: (*Alias)(u),
    })
}

