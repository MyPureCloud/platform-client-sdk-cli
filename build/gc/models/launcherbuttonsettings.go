package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LauncherbuttonsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LauncherbuttonsettingsDud struct { 
    

}

// Launcherbuttonsettings - The settings for the launcher button
type Launcherbuttonsettings struct { 
    // Visibility - The visibility settings for the button
    Visibility string `json:"visibility"`

}

// String returns a JSON representation of the model
func (o *Launcherbuttonsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Launcherbuttonsettings) MarshalJSON() ([]byte, error) {
    type Alias Launcherbuttonsettings

    if LauncherbuttonsettingsMarshalled {
        return []byte("{}"), nil
    }
    LauncherbuttonsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Visibility string `json:"visibility"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

