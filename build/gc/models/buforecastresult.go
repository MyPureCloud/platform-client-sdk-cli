package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuforecastresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuforecastresultDud struct { 
    


    


    


    

}

// Buforecastresult
type Buforecastresult struct { 
    // ReferenceStartDate - The reference start date for interval-based data for this forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReferenceStartDate time.Time `json:"referenceStartDate"`


    // PlanningGroups - The forecast data broken up by planning group
    PlanningGroups []Forecastplanninggroupdata `json:"planningGroups"`


    // WeekNumber - The week number represented by this response
    WeekNumber int `json:"weekNumber"`


    // WeekCount - The number of weeks in this forecast
    WeekCount int `json:"weekCount"`

}

// String returns a JSON representation of the model
func (o *Buforecastresult) String() string {
    
     o.PlanningGroups = []Forecastplanninggroupdata{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buforecastresult) MarshalJSON() ([]byte, error) {
    type Alias Buforecastresult

    if BuforecastresultMarshalled {
        return []byte("{}"), nil
    }
    BuforecastresultMarshalled = true

    return json.Marshal(&struct {
        
        ReferenceStartDate time.Time `json:"referenceStartDate"`
        
        PlanningGroups []Forecastplanninggroupdata `json:"planningGroups"`
        
        WeekNumber int `json:"weekNumber"`
        
        WeekCount int `json:"weekCount"`
        *Alias
    }{

        


        
        PlanningGroups: []Forecastplanninggroupdata{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

