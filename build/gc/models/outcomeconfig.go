package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeconfigDud struct { 
    

}

// Outcomeconfig
type Outcomeconfig struct { 
    // Values - A set of valid Outcome UUIDs used to optimize a KPI.
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Outcomeconfig) String() string {
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeconfig) MarshalJSON() ([]byte, error) {
    type Alias Outcomeconfig

    if OutcomeconfigMarshalled {
        return []byte("{}"), nil
    }
    OutcomeconfigMarshalled = true

    return json.Marshal(&struct {
        
        Values []string `json:"values"`
        *Alias
    }{

        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

