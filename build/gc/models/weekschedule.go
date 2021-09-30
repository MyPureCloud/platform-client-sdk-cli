package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeekscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeekscheduleDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`


    


    


    


    


    


    


    


    


    

}

// Weekschedule - Week schedule information
type Weekschedule struct { 
    


    


    // WeekDate - First day of this week schedule in yyyy-MM-dd format
    WeekDate string `json:"weekDate"`


    // Description - Description of the week schedule
    Description string `json:"description"`


    // Published - Whether the week schedule is published
    Published bool `json:"published"`


    // GenerationResults - Summary of the results from the schedule run
    GenerationResults Weekschedulegenerationresult `json:"generationResults"`


    // ShortTermForecast - Short term forecast associated with this schedule
    ShortTermForecast Shorttermforecastreference `json:"shortTermForecast"`


    // Metadata - Version metadata for this work plan
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // UserSchedules - User schedules in the week
    UserSchedules map[string]Userschedule `json:"userSchedules"`


    // HeadcountForecast - Headcount information for the week schedule
    HeadcountForecast Headcountforecast `json:"headcountForecast"`


    // AgentSchedulesVersion - Version of agent schedules in the week schedule
    AgentSchedulesVersion int `json:"agentSchedulesVersion"`

}

// String returns a JSON representation of the model
func (o *Weekschedule) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.UserSchedules = map[string]Userschedule{"": {}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekschedule) MarshalJSON() ([]byte, error) {
    type Alias Weekschedule

    if WeekscheduleMarshalled {
        return []byte("{}"), nil
    }
    WeekscheduleMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        WeekDate string `json:"weekDate"`
        
        Description string `json:"description"`
        
        Published bool `json:"published"`
        
        GenerationResults Weekschedulegenerationresult `json:"generationResults"`
        
        ShortTermForecast Shorttermforecastreference `json:"shortTermForecast"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        UserSchedules map[string]Userschedule `json:"userSchedules"`
        
        HeadcountForecast Headcountforecast `json:"headcountForecast"`
        
        AgentSchedulesVersion int `json:"agentSchedulesVersion"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        UserSchedules: map[string]Userschedule{"": {}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

