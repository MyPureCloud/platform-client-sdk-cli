package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BugenerateschedulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BugenerateschedulerequestDud struct { 
    


    


    

}

// Bugenerateschedulerequest
type Bugenerateschedulerequest struct { 
    // Description - The description for the schedule
    Description string `json:"description"`


    // ShortTermForecast - The forecast to use when generating the schedule.  Note that the forecast must fully encompass the schedule's start week + week count
    ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`


    // WeekCount - The number of weeks in the schedule. One extra day is added at the end
    WeekCount int `json:"weekCount"`

}

// String returns a JSON representation of the model
func (o *Bugenerateschedulerequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bugenerateschedulerequest) MarshalJSON() ([]byte, error) {
    type Alias Bugenerateschedulerequest

    if BugenerateschedulerequestMarshalled {
        return []byte("{}"), nil
    }
    BugenerateschedulerequestMarshalled = true

    return json.Marshal(&struct { 
        Description string `json:"description"`
        
        ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`
        
        WeekCount int `json:"weekCount"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

