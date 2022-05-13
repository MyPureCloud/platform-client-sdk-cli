package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingsettingsDud struct { 
    

}

// Routingsettings
type Routingsettings struct { 
    // ResetAgentScoreOnPresenceChange - Reset agent score when agent presence changes from off-queue to on-queue
    ResetAgentScoreOnPresenceChange bool `json:"resetAgentScoreOnPresenceChange"`

}

// String returns a JSON representation of the model
func (o *Routingsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingsettings) MarshalJSON() ([]byte, error) {
    type Alias Routingsettings

    if RoutingsettingsMarshalled {
        return []byte("{}"), nil
    }
    RoutingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ResetAgentScoreOnPresenceChange bool `json:"resetAgentScoreOnPresenceChange"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

