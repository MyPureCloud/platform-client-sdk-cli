package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatebusinessunitsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatebusinessunitsettingsDud struct { 
    


    


    


    

}

// Createbusinessunitsettings
type Createbusinessunitsettings struct { 
    // StartDayOfWeek - The start day of week for this business unit
    StartDayOfWeek string `json:"startDayOfWeek"`


    // TimeZone - The time zone for this business unit, using the Olsen tz database format
    TimeZone string `json:"timeZone"`


    // ShortTermForecasting - Short term forecasting settings
    ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`


    // Scheduling - Scheduling settings
    Scheduling Buschedulingsettings `json:"scheduling"`

}

// String returns a JSON representation of the model
func (o *Createbusinessunitsettings) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createbusinessunitsettings) MarshalJSON() ([]byte, error) {
    type Alias Createbusinessunitsettings

    if CreatebusinessunitsettingsMarshalled {
        return []byte("{}"), nil
    }
    CreatebusinessunitsettingsMarshalled = true

    return json.Marshal(&struct {
        
        StartDayOfWeek string `json:"startDayOfWeek"`
        
        TimeZone string `json:"timeZone"`
        
        ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`
        
        Scheduling Buschedulingsettings `json:"scheduling"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

