package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentgreetingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentgreetingDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentgreeting
type Agentgreeting struct { 
    


    // Name
    Name string `json:"name"`


    // InboundPrompt - The agent greeting prompt to use when inbound calls are connected
    InboundPrompt Prompt `json:"inboundPrompt"`


    // OutboundPrompt - The agent greeting prompt to use when outbound calls are connected
    OutboundPrompt Prompt `json:"outboundPrompt"`


    // InboundPromptDefaultLanguage - The default language to use for the agent greeting inbound prompt
    InboundPromptDefaultLanguage string `json:"inboundPromptDefaultLanguage"`


    // OutboundPromptDefaultLanguage - The default language to use for the agent greeting outbound prompt
    OutboundPromptDefaultLanguage string `json:"outboundPromptDefaultLanguage"`


    

}

// String returns a JSON representation of the model
func (o *Agentgreeting) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentgreeting) MarshalJSON() ([]byte, error) {
    type Alias Agentgreeting

    if AgentgreetingMarshalled {
        return []byte("{}"), nil
    }
    AgentgreetingMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        InboundPrompt Prompt `json:"inboundPrompt"`
        
        OutboundPrompt Prompt `json:"outboundPrompt"`
        
        InboundPromptDefaultLanguage string `json:"inboundPromptDefaultLanguage"`
        
        OutboundPromptDefaultLanguage string `json:"outboundPromptDefaultLanguage"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

