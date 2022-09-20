package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitsettingsresponseDud struct { 
    


    


    


    


    

}

// Businessunitsettingsresponse
type Businessunitsettingsresponse struct { 
    // StartDayOfWeek - The start day of week for this business unit
    StartDayOfWeek string `json:"startDayOfWeek"`


    // TimeZone - The time zone for this business unit, using the Olsen tz database format
    TimeZone string `json:"timeZone"`


    // ShortTermForecasting - Short term forecasting settings
    ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`


    // Scheduling - Scheduling settings
    Scheduling Buschedulingsettingsresponse `json:"scheduling"`


    // Metadata - Version metadata for this business unit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Businessunitsettingsresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunitsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Businessunitsettingsresponse

    if BusinessunitsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        StartDayOfWeek string `json:"startDayOfWeek"`
        
        TimeZone string `json:"timeZone"`
        
        ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`
        
        Scheduling Buschedulingsettingsresponse `json:"scheduling"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

