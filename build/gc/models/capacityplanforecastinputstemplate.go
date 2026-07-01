package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanforecastinputstemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanforecastinputstemplateDud struct { 
    


    


    


    


    

}

// Capacityplanforecastinputstemplate
type Capacityplanforecastinputstemplate struct { 
    // ReferenceBusinessUnitDate - The reference date for interval-based data relative to the business unit time zone for the forecast inputs. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    ReferenceBusinessUnitDate time.Time `json:"referenceBusinessUnitDate"`


    // Granularity - Granularity of the intervals
    Granularity string `json:"granularity"`


    // Months - The list of months covered by this capacity plan, formatted as yyyy-MM, populated for monthly granularity
    Months []string `json:"months"`


    // PlanningGroupsForecastData - The forecast data for the planning groups
    PlanningGroupsForecastData []Forecastinputplanninggroupdata `json:"planningGroupsForecastData"`


    // CapacityPlanForecastSummary - The summary of forecast inputs for this capacity plan, for the selected granularity
    CapacityPlanForecastSummary Capacityplanforecastmetrics `json:"capacityPlanForecastSummary"`

}

// String returns a JSON representation of the model
func (o *Capacityplanforecastinputstemplate) String() string {
    
    
     o.Months = []string{""} 
     o.PlanningGroupsForecastData = []Forecastinputplanninggroupdata{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanforecastinputstemplate) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanforecastinputstemplate

    if CapacityplanforecastinputstemplateMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanforecastinputstemplateMarshalled = true

    return json.Marshal(&struct {
        
        ReferenceBusinessUnitDate time.Time `json:"referenceBusinessUnitDate"`
        
        Granularity string `json:"granularity"`
        
        Months []string `json:"months"`
        
        PlanningGroupsForecastData []Forecastinputplanninggroupdata `json:"planningGroupsForecastData"`
        
        CapacityPlanForecastSummary Capacityplanforecastmetrics `json:"capacityPlanForecastSummary"`
        *Alias
    }{

        


        


        
        Months: []string{""},
        


        
        PlanningGroupsForecastData: []Forecastinputplanninggroupdata{{}},
        


        

        Alias: (*Alias)(u),
    })
}

