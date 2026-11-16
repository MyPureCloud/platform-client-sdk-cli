package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastactualvaluesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastactualvaluesDud struct { 
    


    

}

// Capacityplanforecastactualvalues
type Capacityplanforecastactualvalues struct { 
    // OfferedPerInterval - Forecast offered counts per interval, ordered by the granularity, starting from the business unit's start day of the week
    OfferedPerInterval []float64 `json:"offeredPerInterval"`


    // AverageHandleTimeSecondsPerInterval - Forecast average handle time in seconds per interval, ordered by the granularity, starting from the business unit's start day of the week
    AverageHandleTimeSecondsPerInterval []float64 `json:"averageHandleTimeSecondsPerInterval"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastactualvalues) String() string {
     o.OfferedPerInterval = []float64{0.0} 
     o.AverageHandleTimeSecondsPerInterval = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastactualvalues) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastactualvalues

    if CapacityplanforecastactualvaluesMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastactualvaluesMarshalled = true

    return json.Marshal(&struct {
        
        OfferedPerInterval []float64 `json:"offeredPerInterval"`
        
        AverageHandleTimeSecondsPerInterval []float64 `json:"averageHandleTimeSecondsPerInterval"`
        *Alias
    }{

        
        OfferedPerInterval: []float64{0.0},
        


        
        AverageHandleTimeSecondsPerInterval: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

