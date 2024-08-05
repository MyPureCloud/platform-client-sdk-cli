package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplanbidsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplanbidsDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Agentworkplanbids
type Agentworkplanbids struct { 
    


    // BusinessUnit - The business unit to which the bids belong
    BusinessUnit Businessunitreference `json:"businessUnit"`


    // AgentWorkPlanBids - Work plan bid summaries associated with this agent
    AgentWorkPlanBids []Agentworkplanbid `json:"agentWorkPlanBids"`


    

}

// String returns a JSON representation of the model
func (o *Agentworkplanbids) String() string {
    
     o.AgentWorkPlanBids = []Agentworkplanbid{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplanbids) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplanbids

    if AgentworkplanbidsMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplanbidsMarshalled = true

    return json.Marshal(&struct {
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        
        AgentWorkPlanBids []Agentworkplanbid `json:"agentWorkPlanBids"`
        *Alias
    }{

        


        


        
        AgentWorkPlanBids: []Agentworkplanbid{{}},
        


        

        Alias: (*Alias)(u),
    })
}

