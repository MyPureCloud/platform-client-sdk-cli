package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShorttermforecastingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShorttermforecastingsettingsDud struct { 
    

}

// Shorttermforecastingsettings - Short Term Forecasting Settings
type Shorttermforecastingsettings struct { 
    // DefaultHistoryWeeks - The number of weeks to consider by default when generating a volume forecast
    DefaultHistoryWeeks int `json:"defaultHistoryWeeks"`

}

// String returns a JSON representation of the model
func (o *Shorttermforecastingsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shorttermforecastingsettings) MarshalJSON() ([]byte, error) {
    type Alias Shorttermforecastingsettings

    if ShorttermforecastingsettingsMarshalled {
        return []byte("{}"), nil
    }
    ShorttermforecastingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        DefaultHistoryWeeks int `json:"defaultHistoryWeeks"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

