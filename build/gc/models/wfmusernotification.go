package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmusernotificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmusernotificationDud struct { 
    


    


    Timestamp time.Time `json:"timestamp"`


    VarType string `json:"type"`


    ShiftTrade Shifttradenotification `json:"shiftTrade"`


    TimeOffRequest Timeoffrequestnotification `json:"timeOffRequest"`


    AdherenceExplanation Adherenceexplanationnotification `json:"adherenceExplanation"`


    


    AgentNotification bool `json:"agentNotification"`


    OtherNotificationIdsInGroup []string `json:"otherNotificationIdsInGroup"`

}

// Wfmusernotification
type Wfmusernotification struct { 
    // Id - The immutable globally unique identifier for the object.
    Id string `json:"id"`


    // MutableGroupId - The group ID of the notification (mutable, may change  on update)
    MutableGroupId string `json:"mutableGroupId"`


    


    


    


    


    


    // MarkedAsRead - Whether this notification has been marked \"read\"
    MarkedAsRead bool `json:"markedAsRead"`


    


    

}

// String returns a JSON representation of the model
func (o *Wfmusernotification) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmusernotification) MarshalJSON() ([]byte, error) {
    type Alias Wfmusernotification

    if WfmusernotificationMarshalled {
        return []byte("{}"), nil
    }
    WfmusernotificationMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        MutableGroupId string `json:"mutableGroupId"`
        
        
        
        
        
        
        
        
        
        
        
        MarkedAsRead bool `json:"markedAsRead"`
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

