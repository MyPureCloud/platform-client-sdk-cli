package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImportscheduleuploadschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImportscheduleuploadschemaDud struct { 
    


    


    


    


    


    

}

// Importscheduleuploadschema
type Importscheduleuploadschema struct { 
    // Description - The description for the imported schedule
    Description string `json:"description"`


    // WeekCount - The number of weeks the imported schedule will cover
    WeekCount int `json:"weekCount"`


    // Published - Whether the imported schedule should be immediately published
    Published bool `json:"published"`


    // ShortTermForecast - The short term forecast to associate with the imported schedule
    ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`


    // HeadcountForecast - The headcount forecast to associate with the imported schedule
    HeadcountForecast Buheadcountforecast `json:"headcountForecast"`


    // AgentSchedules - Individual agent schedules
    AgentSchedules []Buimportagentscheduleuploadschema `json:"agentSchedules"`

}

// String returns a JSON representation of the model
func (o *Importscheduleuploadschema) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AgentSchedules = []Buimportagentscheduleuploadschema{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Importscheduleuploadschema) MarshalJSON() ([]byte, error) {
    type Alias Importscheduleuploadschema

    if ImportscheduleuploadschemaMarshalled {
        return []byte("{}"), nil
    }
    ImportscheduleuploadschemaMarshalled = true

    return json.Marshal(&struct { 
        Description string `json:"description"`
        
        WeekCount int `json:"weekCount"`
        
        Published bool `json:"published"`
        
        ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`
        
        HeadcountForecast Buheadcountforecast `json:"headcountForecast"`
        
        AgentSchedules []Buimportagentscheduleuploadschema `json:"agentSchedules"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        AgentSchedules: []Buimportagentscheduleuploadschema{{}},
        

        
        Alias: (*Alias)(u),
    })
}

