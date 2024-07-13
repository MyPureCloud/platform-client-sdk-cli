package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperusersearchruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperusersearchruleDud struct { 
    

}

// Valuewrapperusersearchrule
type Valuewrapperusersearchrule struct { 
    // Value - The value for the associated field
    Value Usersearchrule `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperusersearchrule) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperusersearchrule) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperusersearchrule

    if ValuewrapperusersearchruleMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperusersearchruleMarshalled = true

    return json.Marshal(&struct {
        
        Value Usersearchrule `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

