package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectiontimeintervaldataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectiontimeintervaldataDud struct { 
    


    


    


    


    

}

// Anomalydetectiontimeintervaldata
type Anomalydetectiontimeintervaldata struct { 
    // DateIntervalStart - Start time of the interval, in ISO-8601 format
    DateIntervalStart time.Time `json:"dateIntervalStart"`


    // DateIntervalEnd - End time of the interval, in ISO-8601 format
    DateIntervalEnd time.Time `json:"dateIntervalEnd"`


    // ExpectedValue - Expected value for this time interval
    ExpectedValue float64 `json:"expectedValue"`


    // LowerBound - Lower bound of the forecasted range
    LowerBound float64 `json:"lowerBound"`


    // UpperBound - Upper bound of the forecasted range
    UpperBound float64 `json:"upperBound"`

}

// String returns a JSON representation of the model
func (o *Anomalydetectiontimeintervaldata) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectiontimeintervaldata) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectiontimeintervaldata

    if AnomalydetectiontimeintervaldataMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectiontimeintervaldataMarshalled = true

    return json.Marshal(&struct {
        
        DateIntervalStart time.Time `json:"dateIntervalStart"`
        
        DateIntervalEnd time.Time `json:"dateIntervalEnd"`
        
        ExpectedValue float64 `json:"expectedValue"`
        
        LowerBound float64 `json:"lowerBound"`
        
        UpperBound float64 `json:"upperBound"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

