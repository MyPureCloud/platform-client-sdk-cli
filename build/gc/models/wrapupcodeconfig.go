package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WrapupcodeconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WrapupcodeconfigDud struct { 
    

}

// Wrapupcodeconfig
type Wrapupcodeconfig struct { 
    // Values - A set of valid Wrap Up Code UUIDs used to optimize a KPI.
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Wrapupcodeconfig) String() string {
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wrapupcodeconfig) MarshalJSON() ([]byte, error) {
    type Alias Wrapupcodeconfig

    if WrapupcodeconfigMarshalled {
        return []byte("{}"), nil
    }
    WrapupcodeconfigMarshalled = true

    return json.Marshal(&struct {
        
        Values []string `json:"values"`
        *Alias
    }{

        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

