package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalysummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalysummaryDud struct { 
    


    


    

}

// Anomalysummary
type Anomalysummary struct { 
    // DateIntervalStart - The start of the time interval for the anomaly summary count, in ISO-8601 format
    DateIntervalStart time.Time `json:"dateIntervalStart"`


    // DateIntervalEnd - The end of the time interval for the anomaly summary count, in ISO-8601 format
    DateIntervalEnd time.Time `json:"dateIntervalEnd"`


    // AnomalyCount - The number of anomalies detected in the time window
    AnomalyCount int `json:"anomalyCount"`

}

// String returns a JSON representation of the model
func (o *Anomalysummary) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalysummary) MarshalJSON() ([]byte, error) {
    type Alias Anomalysummary

    if AnomalysummaryMarshalled {
        return []byte("{}"), nil
    }
    AnomalysummaryMarshalled = true

    return json.Marshal(&struct {
        
        DateIntervalStart time.Time `json:"dateIntervalStart"`
        
        DateIntervalEnd time.Time `json:"dateIntervalEnd"`
        
        AnomalyCount int `json:"anomalyCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

