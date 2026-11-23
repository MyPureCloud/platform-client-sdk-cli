package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalysummarysettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalysummarysettingsDud struct { 
    

}

// Anomalysummarysettings
type Anomalysummarysettings struct { 
    // Granularity - The granularity used for grouping anomaly counts into time windows
    Granularity string `json:"granularity"`

}

// String returns a JSON representation of the model
func (o *Anomalysummarysettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalysummarysettings) MarshalJSON() ([]byte, error) {
    type Alias Anomalysummarysettings

    if AnomalysummarysettingsMarshalled {
        return []byte("{}"), nil
    }
    AnomalysummarysettingsMarshalled = true

    return json.Marshal(&struct {
        
        Granularity string `json:"granularity"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

