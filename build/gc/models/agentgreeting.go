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


    // InboundPrompt - The agent greeting prompt to use when the call is connected
    InboundPrompt Prompt `json:"inboundPrompt"`


    // OutboundPrompt - The agent greeting prompt to use when the call is about to be disconnected
    OutboundPrompt Prompt `json:"outboundPrompt"`


    

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
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

