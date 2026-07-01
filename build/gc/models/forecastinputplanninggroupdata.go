package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ForecastinputplanninggroupdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ForecastinputplanninggroupdataDud struct { 
    


    

}

// Forecastinputplanninggroupdata
type Forecastinputplanninggroupdata struct { 
    // PlanningGroupId - The ID of the planning group for which this data applies
    PlanningGroupId string `json:"planningGroupId"`


    // CapacityPlanForecastMetrics - The capacity plan forecast metrics for this planning group
    CapacityPlanForecastMetrics Capacityplanforecastmetrics `json:"capacityPlanForecastMetrics"`

}

// String returns a JSON representation of the model
func (o *Forecastinputplanninggroupdata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Forecastinputplanninggroupdata) MarshalJSON() ([]byte, error) {
    type Alias Forecastinputplanninggroupdata

    if ForecastinputplanninggroupdataMarshalled {
        return []byte("{}"), nil
    }
    ForecastinputplanninggroupdataMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        CapacityPlanForecastMetrics Capacityplanforecastmetrics `json:"capacityPlanForecastMetrics"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

