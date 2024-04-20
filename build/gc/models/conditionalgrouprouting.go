package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConditionalgrouproutingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConditionalgrouproutingDud struct { 
    

}

// Conditionalgrouprouting
type Conditionalgrouprouting struct { 
    // Rules - The set of rules to be executed for each conversation
    Rules []Conditionalgrouproutingrule `json:"rules"`

}

// String returns a JSON representation of the model
func (o *Conditionalgrouprouting) String() string {
     o.Rules = []Conditionalgrouproutingrule{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conditionalgrouprouting) MarshalJSON() ([]byte, error) {
    type Alias Conditionalgrouprouting

    if ConditionalgrouproutingMarshalled {
        return []byte("{}"), nil
    }
    ConditionalgrouproutingMarshalled = true

    return json.Marshal(&struct {
        
        Rules []Conditionalgrouproutingrule `json:"rules"`
        *Alias
    }{

        
        Rules: []Conditionalgrouproutingrule{{}},
        

        Alias: (*Alias)(u),
    })
}

