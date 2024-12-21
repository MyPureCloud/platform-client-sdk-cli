package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstatecountsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstatecountsrequestDud struct { 
    


    

}

// Agentstatecountsrequest
type Agentstatecountsrequest struct { 
    // UserFilter - Filters that target user-level data
    UserFilter Agentstateuserfilter `json:"userFilter"`


    // SessionFilter - Filters that target session-level data
    SessionFilter Agentstatesessionfilter `json:"sessionFilter"`

}

// String returns a JSON representation of the model
func (o *Agentstatecountsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstatecountsrequest) MarshalJSON() ([]byte, error) {
    type Alias Agentstatecountsrequest

    if AgentstatecountsrequestMarshalled {
        return []byte("{}"), nil
    }
    AgentstatecountsrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserFilter Agentstateuserfilter `json:"userFilter"`
        
        SessionFilter Agentstatesessionfilter `json:"sessionFilter"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

