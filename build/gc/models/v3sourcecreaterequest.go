package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcecreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcecreaterequestDud struct { 
    


    


    


    


    


    

}

// V3sourcecreaterequest
type V3sourcecreaterequest struct { 
    // Name - The name of the source.
    Name string `json:"name"`


    // VarType - The type of the source. Required if connectionId is not specified, inherits the connection type otherwise.
    VarType string `json:"type"`


    // ConnectionId - The id of the connection related to the source. Required if type is Sharepoint.
    ConnectionId string `json:"connectionId"`


    // TriggerType - The trigger type of the source. Default is Manual.
    TriggerType string `json:"triggerType"`


    // ScheduleSettings - Settings that determine when the source starts a sync. Required if triggerType is Scheduled.
    ScheduleSettings V3sourceschedulesettings `json:"scheduleSettings"`


    // Filters - Filters that determine what documents are synced.
    Filters V3sourcefilter `json:"filters"`

}

// String returns a JSON representation of the model
func (o *V3sourcecreaterequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcecreaterequest) MarshalJSON() ([]byte, error) {
    type Alias V3sourcecreaterequest

    if V3sourcecreaterequestMarshalled {
        return []byte("{}"), nil
    }
    V3sourcecreaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        ConnectionId string `json:"connectionId"`
        
        TriggerType string `json:"triggerType"`
        
        ScheduleSettings V3sourceschedulesettings `json:"scheduleSettings"`
        
        Filters V3sourcefilter `json:"filters"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

