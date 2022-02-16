package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitsettingsDud struct { 
    


    


    


    


    

}

// Businessunitsettings
type Businessunitsettings struct { 
    // StartDayOfWeek - The start day of week for this business unit
    StartDayOfWeek string `json:"startDayOfWeek"`


    // TimeZone - The time zone for this business unit, using the Olsen tz database format
    TimeZone string `json:"timeZone"`


    // ShortTermForecasting - Short term forecasting settings
    ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`


    // Scheduling - Scheduling settings
    Scheduling Buschedulingsettings `json:"scheduling"`


    // Metadata - Version metadata for this business unit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Businessunitsettings) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunitsettings) MarshalJSON() ([]byte, error) {
    type Alias Businessunitsettings

    if BusinessunitsettingsMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitsettingsMarshalled = true

    return json.Marshal(&struct { 
        StartDayOfWeek string `json:"startDayOfWeek"`
        
        TimeZone string `json:"timeZone"`
        
        ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`
        
        Scheduling Buschedulingsettings `json:"scheduling"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

