package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingdataDud struct { 
    


    


    


    


    


    


    


    

}

// Routingdata
type Routingdata struct { 
    // QueueId - The identifier of the routing queue
    QueueId string `json:"queueId"`


    // LanguageId - The identifier of a language to be considered in routing
    LanguageId string `json:"languageId"`


    // Label - An optional label that categorizes the conversation.  Max-utilization settings can be configured at a per-label level
    Label string `json:"label"`


    // Priority - The priority for routing
    Priority int `json:"priority"`


    // SkillIds - A list of skill identifiers to be considered in routing
    SkillIds []string `json:"skillIds"`


    // PreferredAgentIds - A list of agents to be preferred in routing
    PreferredAgentIds []string `json:"preferredAgentIds"`


    // ScoredAgents - A list of scored agents for routing decisions. For Agent Owned Callbacks use one scored agent with a score of 100.
    ScoredAgents []Scoredagent `json:"scoredAgents"`


    // RoutingFlags - An array of flags indicating how the conversation should be routed. Use \"AGENT_OWNED_CALLBACK\" when creating an Agent Owned Callback.
    RoutingFlags []string `json:"routingFlags"`

}

// String returns a JSON representation of the model
func (o *Routingdata) String() string {
    
    
    
    
     o.SkillIds = []string{""} 
     o.PreferredAgentIds = []string{""} 
     o.ScoredAgents = []Scoredagent{{}} 
     o.RoutingFlags = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingdata) MarshalJSON() ([]byte, error) {
    type Alias Routingdata

    if RoutingdataMarshalled {
        return []byte("{}"), nil
    }
    RoutingdataMarshalled = true

    return json.Marshal(&struct {
        
        QueueId string `json:"queueId"`
        
        LanguageId string `json:"languageId"`
        
        Label string `json:"label"`
        
        Priority int `json:"priority"`
        
        SkillIds []string `json:"skillIds"`
        
        PreferredAgentIds []string `json:"preferredAgentIds"`
        
        ScoredAgents []Scoredagent `json:"scoredAgents"`
        
        RoutingFlags []string `json:"routingFlags"`
        *Alias
    }{

        


        


        


        


        
        SkillIds: []string{""},
        


        
        PreferredAgentIds: []string{""},
        


        
        ScoredAgents: []Scoredagent{{}},
        


        
        RoutingFlags: []string{""},
        

        Alias: (*Alias)(u),
    })
}

