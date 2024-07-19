package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentdirectroutingbackupsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentdirectroutingbackupsettingsDud struct { 
    


    


    


    


    BackedUpUsers []string `json:"backedUpUsers"`

}

// Agentdirectroutingbackupsettings
type Agentdirectroutingbackupsettings struct { 
    // QueueId - ID of queue to be used as backup. If queueId and userId are both specified, queue behaves as secondary backup.
    QueueId string `json:"queueId"`


    // UserId - ID of user to be used as backup. If queueId and userId are both specified, user behaves as primary backup.
    UserId string `json:"userId"`


    // WaitForAgent - Flag indicating if Direct Routing interactions should wait for Direct Routing agent or go immediately to selected backup.
    WaitForAgent bool `json:"waitForAgent"`


    // AgentWaitSeconds - Time (in seconds) that a Direct Routing interaction will wait for Direct Routing agent before going to selected backup. Valid range [60, 864000].
    AgentWaitSeconds int `json:"agentWaitSeconds"`


    

}

// String returns a JSON representation of the model
func (o *Agentdirectroutingbackupsettings) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentdirectroutingbackupsettings) MarshalJSON() ([]byte, error) {
    type Alias Agentdirectroutingbackupsettings

    if AgentdirectroutingbackupsettingsMarshalled {
        return []byte("{}"), nil
    }
    AgentdirectroutingbackupsettingsMarshalled = true

    return json.Marshal(&struct {
        
        QueueId string `json:"queueId"`
        
        UserId string `json:"userId"`
        
        WaitForAgent bool `json:"waitForAgent"`
        
        AgentWaitSeconds int `json:"agentWaitSeconds"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

