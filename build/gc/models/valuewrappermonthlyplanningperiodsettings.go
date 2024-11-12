package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrappermonthlyplanningperiodsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrappermonthlyplanningperiodsettingsDud struct { 
    

}

// Valuewrappermonthlyplanningperiodsettings
type Valuewrappermonthlyplanningperiodsettings struct { 
    // Value - The value for the associated field
    Value Monthlyplanningperiodsettings `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrappermonthlyplanningperiodsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrappermonthlyplanningperiodsettings) MarshalJSON() ([]byte, error) {
    type Alias Valuewrappermonthlyplanningperiodsettings

    if ValuewrappermonthlyplanningperiodsettingsMarshalled {
        return []byte("{}"), nil
    }
    ValuewrappermonthlyplanningperiodsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Value Monthlyplanningperiodsettings `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

