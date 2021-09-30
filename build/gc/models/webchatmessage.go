package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatmessageDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Webchatmessage
type Webchatmessage struct { 
    


    // Name
    Name string `json:"name"`


    // Conversation - The identifier of the conversation
    Conversation Webchatconversation `json:"conversation"`


    // Sender - The member who sent the message
    Sender Webchatmemberinfo `json:"sender"`


    // Body - The message body.
    Body string `json:"body"`


    // BodyType - The purpose of the message within the conversation, such as a standard text entry versus a greeting.
    BodyType string `json:"bodyType"`


    // Timestamp - The timestamp of the message, in ISO-8601 format
    Timestamp time.Time `json:"timestamp"`


    

}

// String returns a JSON representation of the model
func (o *Webchatmessage) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatmessage) MarshalJSON() ([]byte, error) {
    type Alias Webchatmessage

    if WebchatmessageMarshalled {
        return []byte("{}"), nil
    }
    WebchatmessageMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Conversation Webchatconversation `json:"conversation"`
        
        Sender Webchatmemberinfo `json:"sender"`
        
        Body string `json:"body"`
        
        BodyType string `json:"bodyType"`
        
        Timestamp time.Time `json:"timestamp"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

