package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicyerrorsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicyerrorsDud struct { 
    

}

// Policyerrors
type Policyerrors struct { 
    // PolicyErrorMessages
    PolicyErrorMessages []Policyerrormessage `json:"policyErrorMessages"`

}

// String returns a JSON representation of the model
func (o *Policyerrors) String() string {
     o.PolicyErrorMessages = []Policyerrormessage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policyerrors) MarshalJSON() ([]byte, error) {
    type Alias Policyerrors

    if PolicyerrorsMarshalled {
        return []byte("{}"), nil
    }
    PolicyerrorsMarshalled = true

    return json.Marshal(&struct {
        
        PolicyErrorMessages []Policyerrormessage `json:"policyErrorMessages"`
        *Alias
    }{

        
        PolicyErrorMessages: []Policyerrormessage{{}},
        

        Alias: (*Alias)(u),
    })
}

