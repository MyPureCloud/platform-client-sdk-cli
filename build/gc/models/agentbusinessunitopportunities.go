package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentbusinessunitopportunitiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentbusinessunitopportunitiesDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Agentbusinessunitopportunities
type Agentbusinessunitopportunities struct { 
    


    // Opportunities - The opportunities for the agent in this business unit
    Opportunities []Agentqueryopportunityresult `json:"opportunities"`


    

}

// String returns a JSON representation of the model
func (o *Agentbusinessunitopportunities) String() string {
     o.Opportunities = []Agentqueryopportunityresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentbusinessunitopportunities) MarshalJSON() ([]byte, error) {
    type Alias Agentbusinessunitopportunities

    if AgentbusinessunitopportunitiesMarshalled {
        return []byte("{}"), nil
    }
    AgentbusinessunitopportunitiesMarshalled = true

    return json.Marshal(&struct {
        
        Opportunities []Agentqueryopportunityresult `json:"opportunities"`
        *Alias
    }{

        


        
        Opportunities: []Agentqueryopportunityresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

