package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrappergroupsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrappergroupsettingsDud struct { 
    

}

// Valuewrappergroupsettings
type Valuewrappergroupsettings struct { 
    // Value - The value for the associated field
    Value Groupsettings `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrappergroupsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrappergroupsettings) MarshalJSON() ([]byte, error) {
    type Alias Valuewrappergroupsettings

    if ValuewrappergroupsettingsMarshalled {
        return []byte("{}"), nil
    }
    ValuewrappergroupsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Value Groupsettings `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

