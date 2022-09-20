package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatebusinessunitsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatebusinessunitsettingsrequestDud struct { 
    StartDayOfWeek string `json:"startDayOfWeek"`


    TimeZone string `json:"timeZone"`


    


    


    

}

// Updatebusinessunitsettingsrequest
type Updatebusinessunitsettingsrequest struct { 
    


    


    // ShortTermForecasting - Short term forecasting settings
    ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`


    // Scheduling - Scheduling settings
    Scheduling Buschedulingsettingsrequest `json:"scheduling"`


    // Metadata - Version metadata for this business unit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updatebusinessunitsettingsrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatebusinessunitsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatebusinessunitsettingsrequest

    if UpdatebusinessunitsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatebusinessunitsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        ShortTermForecasting Bushorttermforecastingsettings `json:"shortTermForecasting"`
        
        Scheduling Buschedulingsettingsrequest `json:"scheduling"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

