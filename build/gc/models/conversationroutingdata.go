package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationroutingdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationroutingdataDud struct { 
    


    


    


    


    


    

}

// Conversationroutingdata
type Conversationroutingdata struct { 
    // Queue - The queue to use for routing decisions
    Queue Addressableentityref `json:"queue"`


    // Language - The language to use for routing decisions
    Language Addressableentityref `json:"language"`


    // Priority - The priority of the conversation to use for routing decisions
    Priority int `json:"priority"`


    // Skills - The skills to use for routing decisions
    Skills []Addressableentityref `json:"skills"`


    // ScoredAgents - A collection of agents and their assigned scores for this conversation (0 - 100, higher being better), for use in routing to preferred agents
    ScoredAgents []Scoredagent `json:"scoredAgents"`


    // Label - An optional label that categorizes the conversation.  Max-utilization settings can be configured at a per-label level
    Label string `json:"label"`

}

// String returns a JSON representation of the model
func (o *Conversationroutingdata) String() string {
    
    
    
     o.Skills = []Addressableentityref{{}} 
     o.ScoredAgents = []Scoredagent{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationroutingdata) MarshalJSON() ([]byte, error) {
    type Alias Conversationroutingdata

    if ConversationroutingdataMarshalled {
        return []byte("{}"), nil
    }
    ConversationroutingdataMarshalled = true

    return json.Marshal(&struct {
        
        Queue Addressableentityref `json:"queue"`
        
        Language Addressableentityref `json:"language"`
        
        Priority int `json:"priority"`
        
        Skills []Addressableentityref `json:"skills"`
        
        ScoredAgents []Scoredagent `json:"scoredAgents"`
        
        Label string `json:"label"`
        *Alias
    }{

        


        


        


        
        Skills: []Addressableentityref{{}},
        


        
        ScoredAgents: []Scoredagent{{}},
        


        

        Alias: (*Alias)(u),
    })
}

