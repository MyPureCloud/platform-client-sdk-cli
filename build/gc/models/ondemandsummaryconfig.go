package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OndemandsummaryconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OndemandsummaryconfigDud struct { 
    

}

// Ondemandsummaryconfig
type Ondemandsummaryconfig struct { 
    // Enabled - On demand summary is enabled.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Ondemandsummaryconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ondemandsummaryconfig) MarshalJSON() ([]byte, error) {
    type Alias Ondemandsummaryconfig

    if OndemandsummaryconfigMarshalled {
        return []byte("{}"), nil
    }
    OndemandsummaryconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

