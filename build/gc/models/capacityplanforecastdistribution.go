package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastdistributionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastdistributionDud struct { 
    


    

}

// Capacityplanforecastdistribution
type Capacityplanforecastdistribution struct { 
    // DayOfWeek - The day of the week to which this distribution applies
    DayOfWeek string `json:"dayOfWeek"`


    // DistributionPerInterval - Hourly distribution values for the day, ordered from hour 0 to hour 23
    DistributionPerInterval []float64 `json:"distributionPerInterval"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastdistribution) String() string {
    
     o.DistributionPerInterval = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastdistribution) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastdistribution

    if CapacityplanforecastdistributionMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastdistributionMarshalled = true

    return json.Marshal(&struct {
        
        DayOfWeek string `json:"dayOfWeek"`
        
        DistributionPerInterval []float64 `json:"distributionPerInterval"`
        *Alias
    }{

        


        
        DistributionPerInterval: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

