package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommonalertbulkupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommonalertbulkupdaterequestDud struct { 
    


    


    


    

}

// Commonalertbulkupdaterequest
type Commonalertbulkupdaterequest struct { 
    // VarType - The action to take
    VarType string `json:"type"`


    // AlertIds - The user supplied alert ids to be muted
    AlertIds []string `json:"alertIds"`


    // MuteSnooze - The fields need for a mute or snooze requests
    MuteSnooze Mutesnoozefields `json:"muteSnooze"`


    // Unread - The fields need for an unread update requests
    Unread Unreadfields `json:"unread"`

}

// String returns a JSON representation of the model
func (o *Commonalertbulkupdaterequest) String() string {
    
     o.AlertIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commonalertbulkupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Commonalertbulkupdaterequest

    if CommonalertbulkupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    CommonalertbulkupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        AlertIds []string `json:"alertIds"`
        
        MuteSnooze Mutesnoozefields `json:"muteSnooze"`
        
        Unread Unreadfields `json:"unread"`
        *Alias
    }{

        


        
        AlertIds: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

