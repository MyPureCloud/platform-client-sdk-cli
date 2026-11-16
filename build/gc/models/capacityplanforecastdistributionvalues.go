package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastdistributionvaluesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastdistributionvaluesDud struct { 
    


    

}

// Capacityplanforecastdistributionvalues
type Capacityplanforecastdistributionvalues struct { 
    // OfferedDistributions - Hourly distribution of forecast offered volume for the forecast week, with one entry per day of the week. Each value is a proportional weight between 0.0 to 1.0. Each day's 24 values must sum to approximately 1.0 (0.999 to 1.001 to allow for rounding)
    OfferedDistributions []Capacityplanforecastdistribution `json:"offeredDistributions"`


    // AverageHandleTimeDistributions - Hourly distribution of forecast average handle time for the forecast week, with one entry per day of the week. Each value is a weighting factor relative to the daily average handle time. Each day's 24 values must average to approximately 1.0 (0.999 to 1.001 to allow for rounding)
    AverageHandleTimeDistributions []Capacityplanforecastdistribution `json:"averageHandleTimeDistributions"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastdistributionvalues) String() string {
     o.OfferedDistributions = []Capacityplanforecastdistribution{{}} 
     o.AverageHandleTimeDistributions = []Capacityplanforecastdistribution{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastdistributionvalues) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastdistributionvalues

    if CapacityplanforecastdistributionvaluesMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastdistributionvaluesMarshalled = true

    return json.Marshal(&struct {
        
        OfferedDistributions []Capacityplanforecastdistribution `json:"offeredDistributions"`
        
        AverageHandleTimeDistributions []Capacityplanforecastdistribution `json:"averageHandleTimeDistributions"`
        *Alias
    }{

        
        OfferedDistributions: []Capacityplanforecastdistribution{{}},
        


        
        AverageHandleTimeDistributions: []Capacityplanforecastdistribution{{}},
        

        Alias: (*Alias)(u),
    })
}

