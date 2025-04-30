package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConditionalgroupactivationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConditionalgroupactivationDud struct { 
    


    

}

// Conditionalgroupactivation
type Conditionalgroupactivation struct { 
    // PilotRule - The pilot rule for this queue, which executes periodically to determine queue health
    PilotRule Conditionalgroupactivationpilotrule `json:"pilotRule"`


    // Rules - The set of rules to be periodically executed on the queue (if the pilot rule evaluates as true or there is no pilot rule)
    Rules []Conditionalgroupactivationrule `json:"rules"`

}

// String returns a JSON representation of the model
func (o *Conditionalgroupactivation) String() string {
    
     o.Rules = []Conditionalgroupactivationrule{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conditionalgroupactivation) MarshalJSON() ([]byte, error) {
    type Alias Conditionalgroupactivation

    if ConditionalgroupactivationMarshalled {
        return []byte("{}"), nil
    }
    ConditionalgroupactivationMarshalled = true

    return json.Marshal(&struct {
        
        PilotRule Conditionalgroupactivationpilotrule `json:"pilotRule"`
        
        Rules []Conditionalgroupactivationrule `json:"rules"`
        *Alias
    }{

        


        
        Rules: []Conditionalgroupactivationrule{{}},
        

        Alias: (*Alias)(u),
    })
}

