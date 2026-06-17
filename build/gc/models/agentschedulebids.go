package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentschedulebidsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentschedulebidsDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Agentschedulebids
type Agentschedulebids struct { 
    


    // BusinessUnit - The business unit to which the bids belong
    BusinessUnit Businessunitreference `json:"businessUnit"`


    // AgentScheduleBids - Schedule bid summaries associated with this agent
    AgentScheduleBids []Agentschedulebid `json:"agentScheduleBids"`


    

}

// String returns a JSON representation of the model
func (o *Agentschedulebids) String() string {
    
     o.AgentScheduleBids = []Agentschedulebid{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentschedulebids) MarshalJSON() ([]byte, error) {
    type Alias Agentschedulebids

    if AgentschedulebidsMarshalled {
        return []byte("{}"), nil
    }
    AgentschedulebidsMarshalled = true

    return json.Marshal(&struct {
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        
        AgentScheduleBids []Agentschedulebid `json:"agentScheduleBids"`
        *Alias
    }{

        


        


        
        AgentScheduleBids: []Agentschedulebid{{}},
        


        

        Alias: (*Alias)(u),
    })
}

