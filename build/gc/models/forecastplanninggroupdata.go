package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastplanninggroupdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastplanninggroupdataDud struct { 
    


    


    

}

// Forecastplanninggroupdata
type Forecastplanninggroupdata struct { 
    // PlanningGroupId - The ID of the planning group to which this data applies. Note this is a snapshot of the planning group at the time of forecast creation and may not correspond to the current configuration
    PlanningGroupId string `json:"planningGroupId"`


    // OfferedPerInterval - Forecast offered counts per 15 minute interval for this week of the forecast
    OfferedPerInterval []float64 `json:"offeredPerInterval"`


    // AverageHandleTimeSecondsPerInterval - Forecast average handle time per 15 minute interval in seconds
    AverageHandleTimeSecondsPerInterval []float64 `json:"averageHandleTimeSecondsPerInterval"`

}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupdata) String() string {
    
    
    
    
    
    
     o.OfferedPerInterval = []float64{0.0} 
    
    
    
     o.AverageHandleTimeSecondsPerInterval = []float64{0.0} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastplanninggroupdata) MarshalJSON() ([]byte, error) {
    type Alias Forecastplanninggroupdata

    if ForecastplanninggroupdataMarshalled {
        return []byte("{}"), nil
    }
    ForecastplanninggroupdataMarshalled = true

    return json.Marshal(&struct { 
        PlanningGroupId string `json:"planningGroupId"`
        
        OfferedPerInterval []float64 `json:"offeredPerInterval"`
        
        AverageHandleTimeSecondsPerInterval []float64 `json:"averageHandleTimeSecondsPerInterval"`
        
        *Alias
    }{
        

        

        

        
        OfferedPerInterval: []float64{0.0},
        

        

        
        AverageHandleTimeSecondsPerInterval: []float64{0.0},
        

        
        Alias: (*Alias)(u),
    })
}

