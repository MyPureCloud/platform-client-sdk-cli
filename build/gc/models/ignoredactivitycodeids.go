package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IgnoredactivitycodeidsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IgnoredactivitycodeidsDud struct { 
    

}

// Ignoredactivitycodeids
type Ignoredactivitycodeids struct { 
    // Values - List of activity code IDs
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Ignoredactivitycodeids) String() string {
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ignoredactivitycodeids) MarshalJSON() ([]byte, error) {
    type Alias Ignoredactivitycodeids

    if IgnoredactivitycodeidsMarshalled {
        return []byte("{}"), nil
    }
    IgnoredactivitycodeidsMarshalled = true

    return json.Marshal(&struct {
        
        Values []string `json:"values"`
        *Alias
    }{

        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

