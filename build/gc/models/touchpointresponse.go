package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TouchpointresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TouchpointresponseDud struct { 
    


    


    


    


    


    


    

}

// Touchpointresponse
type Touchpointresponse struct { 
    // ConversationId - ID of conversation.
    ConversationId string `json:"conversationId"`


    // AgentId - ID of agent.
    AgentId string `json:"agentId"`


    // AssociatedValue - The value attributed to this touchpoint.
    AssociatedValue float32 `json:"associatedValue"`


    // MediaType - Media Type of the touchpoint; allowed values are Email, Message and Voice.
    MediaType string `json:"mediaType"`


    // State - Outcome Attribution Touchpoint status.
    State string `json:"state"`


    // Message - Additional information on the state of the touchpoint entity, populated when the touchpoint has an error.
    Message string `json:"message"`


    // CreatedDate - Date conversation happened. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Touchpointresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Touchpointresponse) MarshalJSON() ([]byte, error) {
    type Alias Touchpointresponse

    if TouchpointresponseMarshalled {
        return []byte("{}"), nil
    }
    TouchpointresponseMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        
        AgentId string `json:"agentId"`
        
        AssociatedValue float32 `json:"associatedValue"`
        
        MediaType string `json:"mediaType"`
        
        State string `json:"state"`
        
        Message string `json:"message"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

