package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnactionagentdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnactionagentdetailsDud struct { 
    


    


    

}

// Reportingturnactionagentdetails
type Reportingturnactionagentdetails struct { 
    // AgentId - The agent ID used in an action.
    AgentId string `json:"agentId"`


    // AgentName - The agent name used in an action.
    AgentName string `json:"agentName"`


    // AgentVersion - The agent version used in an action.
    AgentVersion string `json:"agentVersion"`

}

// String returns a JSON representation of the model
func (o *Reportingturnactionagentdetails) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnactionagentdetails) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnactionagentdetails

    if ReportingturnactionagentdetailsMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnactionagentdetailsMarshalled = true

    return json.Marshal(&struct {
        
        AgentId string `json:"agentId"`
        
        AgentName string `json:"agentName"`
        
        AgentVersion string `json:"agentVersion"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

