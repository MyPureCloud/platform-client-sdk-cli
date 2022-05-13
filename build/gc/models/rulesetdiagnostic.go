package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RulesetdiagnosticMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RulesetdiagnosticDud struct { 
    RuleSet Domainentityref `json:"ruleSet"`


    Warnings []string `json:"warnings"`

}

// Rulesetdiagnostic
type Rulesetdiagnostic struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Rulesetdiagnostic) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Rulesetdiagnostic) MarshalJSON() ([]byte, error) {
    type Alias Rulesetdiagnostic

    if RulesetdiagnosticMarshalled {
        return []byte("{}"), nil
    }
    RulesetdiagnosticMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

