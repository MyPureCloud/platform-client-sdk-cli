package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationnotificationtemplatefooterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationnotificationtemplatefooterDud struct { 
    

}

// Conversationnotificationtemplatefooter - Template footer object.
type Conversationnotificationtemplatefooter struct { 
    // Text - Footer text. For WhatsApp, ignored.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplatefooter) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationnotificationtemplatefooter) MarshalJSON() ([]byte, error) {
    type Alias Conversationnotificationtemplatefooter

    if ConversationnotificationtemplatefooterMarshalled {
        return []byte("{}"), nil
    }
    ConversationnotificationtemplatefooterMarshalled = true

    return json.Marshal(&struct { 
        Text string `json:"text"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

