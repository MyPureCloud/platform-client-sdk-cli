package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LongtermforecastplanninggroupdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LongtermforecastplanninggroupdataDud struct { 
    


    


    

}

// Longtermforecastplanninggroupdata
type Longtermforecastplanninggroupdata struct { 
    // PlanningGroupId - The ID of the planning group to which this data applies. Note this is a snapshot of the planning group at the time of forecast creation and may not correspond to the current configuration
    PlanningGroupId string `json:"planningGroupId"`


    // OfferedPerDay - Forecast offered counts per day for this planning group
    OfferedPerDay []float64 `json:"offeredPerDay"`


    // AverageHandleTimeSecondsPerDay - Forecast average handle time per day in seconds
    AverageHandleTimeSecondsPerDay []float64 `json:"averageHandleTimeSecondsPerDay"`

}

// String returns a JSON representation of the model
func (o *Longtermforecastplanninggroupdata) String() string {
    
     o.OfferedPerDay = []float64{0.0} 
     o.AverageHandleTimeSecondsPerDay = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Longtermforecastplanninggroupdata) MarshalJSON() ([]byte, error) {
    type Alias Longtermforecastplanninggroupdata

    if LongtermforecastplanninggroupdataMarshalled {
        return []byte("{}"), nil
    }
    LongtermforecastplanninggroupdataMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        OfferedPerDay []float64 `json:"offeredPerDay"`
        
        AverageHandleTimeSecondsPerDay []float64 `json:"averageHandleTimeSecondsPerDay"`
        *Alias
    }{

        


        
        OfferedPerDay: []float64{0.0},
        


        
        AverageHandleTimeSecondsPerDay: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

