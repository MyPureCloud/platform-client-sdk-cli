package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourceupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourceupdaterequestDud struct { 
    


    


    

}

// V3sourceupdaterequest
type V3sourceupdaterequest struct { 
    // Name - The name of the source.
    Name string `json:"name"`


    // TriggerType - The trigger type of the source.
    TriggerType string `json:"triggerType"`


    // ScheduleSettings - Settings that determine when the source starts a sync.
    ScheduleSettings V3sourceschedulesettings `json:"scheduleSettings"`

}

// String returns a JSON representation of the model
func (o *V3sourceupdaterequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourceupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias V3sourceupdaterequest

    if V3sourceupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    V3sourceupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        TriggerType string `json:"triggerType"`
        
        ScheduleSettings V3sourceschedulesettings `json:"scheduleSettings"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

