package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StationsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StationsettingsDud struct { 
    

}

// Stationsettings - Organization settings for stations
type Stationsettings struct { 
    // FreeSeatingConfiguration - Configuration options for free-seating
    FreeSeatingConfiguration Freeseatingconfiguration `json:"freeSeatingConfiguration"`

}

// String returns a JSON representation of the model
func (o *Stationsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stationsettings) MarshalJSON() ([]byte, error) {
    type Alias Stationsettings

    if StationsettingsMarshalled {
        return []byte("{}"), nil
    }
    StationsettingsMarshalled = true

    return json.Marshal(&struct {
        
        FreeSeatingConfiguration Freeseatingconfiguration `json:"freeSeatingConfiguration"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

