package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatemanagementunitsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatemanagementunitsettingsrequestDud struct { 
    


    


    


    


    

}

// Createmanagementunitsettingsrequest
type Createmanagementunitsettingsrequest struct { 
    // Adherence - Adherence settings for this management unit
    Adherence Adherencesettings `json:"adherence"`


    // ShortTermForecasting - Short term forecasting settings for this management unit.  Moving to Business Unit
    ShortTermForecasting Shorttermforecastingsettings `json:"shortTermForecasting"`


    // TimeOff - Time off request settings for this management unit
    TimeOff Timeoffsettingsrequest `json:"timeOff"`


    // Scheduling - Scheduling settings for this management unit
    Scheduling Schedulingsettingsrequest `json:"scheduling"`


    // ShiftTrading - Shift trade settings for this management unit
    ShiftTrading Shifttradesettings `json:"shiftTrading"`

}

// String returns a JSON representation of the model
func (o *Createmanagementunitsettingsrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createmanagementunitsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Createmanagementunitsettingsrequest

    if CreatemanagementunitsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatemanagementunitsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Adherence Adherencesettings `json:"adherence"`
        
        ShortTermForecasting Shorttermforecastingsettings `json:"shortTermForecasting"`
        
        TimeOff Timeoffsettingsrequest `json:"timeOff"`
        
        Scheduling Schedulingsettingsrequest `json:"scheduling"`
        
        ShiftTrading Shifttradesettings `json:"shiftTrading"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

