package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamiccontactqueueingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamiccontactqueueingsettingsDud struct { 
    

}

// Dynamiccontactqueueingsettings
type Dynamiccontactqueueingsettings struct { 
    // Sort - Whether to sort contacts dynamically
    Sort bool `json:"sort"`

}

// String returns a JSON representation of the model
func (o *Dynamiccontactqueueingsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamiccontactqueueingsettings) MarshalJSON() ([]byte, error) {
    type Alias Dynamiccontactqueueingsettings

    if DynamiccontactqueueingsettingsMarshalled {
        return []byte("{}"), nil
    }
    DynamiccontactqueueingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Sort bool `json:"sort"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

