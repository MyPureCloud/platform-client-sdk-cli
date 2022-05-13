package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperplanningperiodsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperplanningperiodsettingsDud struct { 
    

}

// Valuewrapperplanningperiodsettings - An object to provide context to nullable fields in PATCH requests
type Valuewrapperplanningperiodsettings struct { 
    // Value - The value for the associated field
    Value Planningperiodsettings `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperplanningperiodsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperplanningperiodsettings) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperplanningperiodsettings

    if ValuewrapperplanningperiodsettingsMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperplanningperiodsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Value Planningperiodsettings `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

