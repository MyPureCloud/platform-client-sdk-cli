package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatebusinessunitsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatebusinessunitsettingsrequestDud struct { 
    


    


    


    

}

// Createbusinessunitsettingsrequest
type Createbusinessunitsettingsrequest struct { 
    // StartDayOfWeek - The start day of week for this business unit
    StartDayOfWeek string `json:"startDayOfWeek"`


    // TimeZone - The time zone for this business unit, using the Olsen tz database format
    TimeZone string `json:"timeZone"`


    // ShortTermForecasting - Short term forecasting settings
    ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`


    // Scheduling - Scheduling settings
    Scheduling Buschedulingsettingsrequest `json:"scheduling"`

}

// String returns a JSON representation of the model
func (o *Createbusinessunitsettingsrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createbusinessunitsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Createbusinessunitsettingsrequest

    if CreatebusinessunitsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatebusinessunitsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDayOfWeek string `json:"startDayOfWeek"`
        
        TimeZone string `json:"timeZone"`
        
        ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`
        
        Scheduling Buschedulingsettingsrequest `json:"scheduling"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

