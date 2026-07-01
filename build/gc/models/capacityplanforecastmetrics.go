package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastmetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastmetricsDud struct { 
    


    

}

// Capacityplanforecastmetrics
type Capacityplanforecastmetrics struct { 
    // Volume - Forecast offered counts per requested granularity interval
    Volume []float64 `json:"volume"`


    // AverageHandleTime - Average handle time in seconds per requested granularity interval
    AverageHandleTime []float64 `json:"averageHandleTime"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastmetrics) String() string {
     o.Volume = []float64{0.0} 
     o.AverageHandleTime = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastmetrics) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastmetrics

    if CapacityplanforecastmetricsMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastmetricsMarshalled = true

    return json.Marshal(&struct {
        
        Volume []float64 `json:"volume"`
        
        AverageHandleTime []float64 `json:"averageHandleTime"`
        *Alias
    }{

        
        Volume: []float64{0.0},
        


        
        AverageHandleTime: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

