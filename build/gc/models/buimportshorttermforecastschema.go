package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuimportshorttermforecastschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuimportshorttermforecastschemaDud struct { 
    


    


    


    


    

}

// Buimportshorttermforecastschema
type Buimportshorttermforecastschema struct { 
    // Description - The description for the forecast
    Description string `json:"description"`


    // WeekCount - The number of weeks covered by the forecast
    WeekCount int `json:"weekCount"`


    // PlanningGroups - The short term planning group data
    PlanningGroups []Forecastplanninggroupdata `json:"planningGroups"`


    // LongTermPlanningGroups - The long term planning group data
    LongTermPlanningGroups []Longtermforecastplanninggroupdata `json:"longTermPlanningGroups"`


    // CanUseForScheduling - Whether this forecast can be used for scheduling
    CanUseForScheduling bool `json:"canUseForScheduling"`

}

// String returns a JSON representation of the model
func (o *Buimportshorttermforecastschema) String() string {
    
    
     o.PlanningGroups = []Forecastplanninggroupdata{{}} 
     o.LongTermPlanningGroups = []Longtermforecastplanninggroupdata{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buimportshorttermforecastschema) MarshalJSON() ([]byte, error) {
    type Alias Buimportshorttermforecastschema

    if BuimportshorttermforecastschemaMarshalled {
        return []byte("{}"), nil
    }
    BuimportshorttermforecastschemaMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        
        WeekCount int `json:"weekCount"`
        
        PlanningGroups []Forecastplanninggroupdata `json:"planningGroups"`
        
        LongTermPlanningGroups []Longtermforecastplanninggroupdata `json:"longTermPlanningGroups"`
        
        CanUseForScheduling bool `json:"canUseForScheduling"`
        *Alias
    }{

        


        


        
        PlanningGroups: []Forecastplanninggroupdata{{}},
        


        
        LongTermPlanningGroups: []Longtermforecastplanninggroupdata{{}},
        


        

        Alias: (*Alias)(u),
    })
}

