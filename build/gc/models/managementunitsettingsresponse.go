package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManagementunitsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManagementunitsettingsresponseDud struct { 
    


    


    


    


    


    

}

// Managementunitsettingsresponse
type Managementunitsettingsresponse struct { 
    // Adherence - Adherence settings for this management unit
    Adherence Adherencesettings `json:"adherence"`


    // ShortTermForecasting - Short term forecasting settings for this management unit
    ShortTermForecasting Shorttermforecastingsettings `json:"shortTermForecasting"`


    // TimeOff - Time off request settings for this management unit
    TimeOff Timeoffrequestsettings `json:"timeOff"`


    // Scheduling - Scheduling settings for this management unit. These settings are only available if you have the permission wfm:managementUnit:view
    Scheduling Schedulingsettingsresponse `json:"scheduling"`


    // ShiftTrading - Shift trade settings for this management unit
    ShiftTrading Shifttradesettings `json:"shiftTrading"`


    // Metadata - Version info metadata for the associated management unit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Managementunitsettingsresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Managementunitsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Managementunitsettingsresponse

    if ManagementunitsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    ManagementunitsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Adherence Adherencesettings `json:"adherence"`
        
        ShortTermForecasting Shorttermforecastingsettings `json:"shortTermForecasting"`
        
        TimeOff Timeoffrequestsettings `json:"timeOff"`
        
        Scheduling Schedulingsettingsresponse `json:"scheduling"`
        
        ShiftTrading Shifttradesettings `json:"shiftTrading"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

