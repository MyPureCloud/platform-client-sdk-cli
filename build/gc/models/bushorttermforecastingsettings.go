package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BushorttermforecastingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BushorttermforecastingsettingsDud struct { 
    

}

// Bushorttermforecastingsettings
type Bushorttermforecastingsettings struct { 
    // DefaultHistoryWeeks - The number of historical weeks to consider when creating a forecast. This setting is only used for legacy weighted average forecasts
    DefaultHistoryWeeks int `json:"defaultHistoryWeeks"`

}

// String returns a JSON representation of the model
func (o *Bushorttermforecastingsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bushorttermforecastingsettings) MarshalJSON() ([]byte, error) {
    type Alias Bushorttermforecastingsettings

    if BushorttermforecastingsettingsMarshalled {
        return []byte("{}"), nil
    }
    BushorttermforecastingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        DefaultHistoryWeeks int `json:"defaultHistoryWeeks"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

