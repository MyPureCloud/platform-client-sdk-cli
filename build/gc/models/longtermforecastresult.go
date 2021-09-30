package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LongtermforecastresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LongtermforecastresultDud struct { 
    


    


    

}

// Longtermforecastresult
type Longtermforecastresult struct { 
    // PlanningGroups - The forecast data broken up by planning group
    PlanningGroups []Longtermforecastplanninggroupdata `json:"planningGroups"`


    // ReferenceStartDate - The reference start date relative to the business unit time zone in this forecast. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    ReferenceStartDate time.Time `json:"referenceStartDate"`


    // WeekCount - The number of weeks in this forecast
    WeekCount int `json:"weekCount"`

}

// String returns a JSON representation of the model
func (o *Longtermforecastresult) String() string {
    
    
     o.PlanningGroups = []Longtermforecastplanninggroupdata{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Longtermforecastresult) MarshalJSON() ([]byte, error) {
    type Alias Longtermforecastresult

    if LongtermforecastresultMarshalled {
        return []byte("{}"), nil
    }
    LongtermforecastresultMarshalled = true

    return json.Marshal(&struct { 
        PlanningGroups []Longtermforecastplanninggroupdata `json:"planningGroups"`
        
        ReferenceStartDate time.Time `json:"referenceStartDate"`
        
        WeekCount int `json:"weekCount"`
        
        *Alias
    }{
        

        
        PlanningGroups: []Longtermforecastplanninggroupdata{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

