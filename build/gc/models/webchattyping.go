package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchattypingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchattypingDud struct { 
    


    


    


    

}

// Webchattyping
type Webchattyping struct { 
    // Id - The event identifier of this typing indicator event (useful to guard against event re-delivery
    Id string `json:"id"`


    // Conversation - The identifier of the conversation
    Conversation Webchatconversation `json:"conversation"`


    // Sender - The member who sent the message
    Sender Webchatmemberinfo `json:"sender"`


    // Timestamp - The timestamp of the message, in ISO-8601 format
    Timestamp time.Time `json:"timestamp"`

}

// String returns a JSON representation of the model
func (o *Webchattyping) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchattyping) MarshalJSON() ([]byte, error) {
    type Alias Webchattyping

    if WebchattypingMarshalled {
        return []byte("{}"), nil
    }
    WebchattypingMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Conversation Webchatconversation `json:"conversation"`
        
        Sender Webchatmemberinfo `json:"sender"`
        
        Timestamp time.Time `json:"timestamp"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

