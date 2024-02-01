package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DirectroutingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DirectroutingDud struct { 
    


    


    


    


    


    

}

// Directrouting
type Directrouting struct { 
    // CallMediaSettings - Direct Routing Settings specific to Call media.
    CallMediaSettings Directroutingmediasettings `json:"callMediaSettings"`


    // EmailMediaSettings - Direct Routing Settings specific to Email media.
    EmailMediaSettings Directroutingmediasettings `json:"emailMediaSettings"`


    // MessageMediaSettings - Direct Routing Settings specific to Message media.
    MessageMediaSettings Directroutingmediasettings `json:"messageMediaSettings"`


    // BackupQueueId - ID of another queue to be used as the default backup if an agent does not have their Backup Settings configured. If not set, the current queue will be used as backup, but with Direct Routing criteria removed from the conversation.
    BackupQueueId string `json:"backupQueueId"`


    // WaitForAgent - Flag indicating if Direct Routing interactions should wait for Direct Routing agent or go immediately to selected backup.
    WaitForAgent bool `json:"waitForAgent"`


    // AgentWaitSeconds - Time (in seconds) that a Direct Routing interaction will wait for Direct Routing agent before going to selected backup. Valid range [60, 864000].
    AgentWaitSeconds int `json:"agentWaitSeconds"`

}

// String returns a JSON representation of the model
func (o *Directrouting) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Directrouting) MarshalJSON() ([]byte, error) {
    type Alias Directrouting

    if DirectroutingMarshalled {
        return []byte("{}"), nil
    }
    DirectroutingMarshalled = true

    return json.Marshal(&struct {
        
        CallMediaSettings Directroutingmediasettings `json:"callMediaSettings"`
        
        EmailMediaSettings Directroutingmediasettings `json:"emailMediaSettings"`
        
        MessageMediaSettings Directroutingmediasettings `json:"messageMediaSettings"`
        
        BackupQueueId string `json:"backupQueueId"`
        
        WaitForAgent bool `json:"waitForAgent"`
        
        AgentWaitSeconds int `json:"agentWaitSeconds"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

