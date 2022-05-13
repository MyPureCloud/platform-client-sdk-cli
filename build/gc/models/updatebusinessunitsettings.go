package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatebusinessunitsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatebusinessunitsettingsDud struct { 
    StartDayOfWeek string `json:"startDayOfWeek"`


    TimeZone string `json:"timeZone"`


    


    


    

}

// Updatebusinessunitsettings
type Updatebusinessunitsettings struct { 
    


    


    // ShortTermForecasting - Short term forecasting settings
    ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`


    // Scheduling - Scheduling settings
    Scheduling Buschedulingsettings `json:"scheduling"`


    // Metadata - Version metadata for this business unit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updatebusinessunitsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatebusinessunitsettings) MarshalJSON() ([]byte, error) {
    type Alias Updatebusinessunitsettings

    if UpdatebusinessunitsettingsMarshalled {
        return []byte("{}"), nil
    }
    UpdatebusinessunitsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`
        
        Scheduling Buschedulingsettings `json:"scheduling"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

