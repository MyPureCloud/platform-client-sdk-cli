package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontextDud struct { 
    


    


    

}

// Conversationcontext
type Conversationcontext struct { 
    // Id - ID of the conversation turn.
    Id string `json:"id"`


    // Participant - Participant type of the entity.
    Participant string `json:"participant"`


    // Text - The conversation text.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Conversationcontext) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontext) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontext

    if ConversationcontextMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontextMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Participant string `json:"participant"`
        
        Text string `json:"text"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

