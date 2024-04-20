package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BackgroundimagesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BackgroundimagesettingsDud struct { 
    

}

// Backgroundimagesettings - The settings to Agent Video background image
type Backgroundimagesettings struct { 
    // Url - BackgroundImage URL for agent video settings
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Backgroundimagesettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Backgroundimagesettings) MarshalJSON() ([]byte, error) {
    type Alias Backgroundimagesettings

    if BackgroundimagesettingsMarshalled {
        return []byte("{}"), nil
    }
    BackgroundimagesettingsMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

