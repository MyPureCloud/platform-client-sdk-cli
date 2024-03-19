package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenrecordingactivesessionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenrecordingactivesessionsDud struct { 
    

}

// Screenrecordingactivesessions
type Screenrecordingactivesessions struct { 
    // Count - Current concurrent active screen recordings count for organization
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Screenrecordingactivesessions) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenrecordingactivesessions) MarshalJSON() ([]byte, error) {
    type Alias Screenrecordingactivesessions

    if ScreenrecordingactivesessionsMarshalled {
        return []byte("{}"), nil
    }
    ScreenrecordingactivesessionsMarshalled = true

    return json.Marshal(&struct {
        
        Count int `json:"count"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

