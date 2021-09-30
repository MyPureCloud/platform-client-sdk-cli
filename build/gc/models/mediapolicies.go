package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediapoliciesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediapoliciesDud struct { 
    


    


    


    

}

// Mediapolicies
type Mediapolicies struct { 
    // CallPolicy - Conditions and actions for calls
    CallPolicy Callmediapolicy `json:"callPolicy"`


    // ChatPolicy - Conditions and actions for chats
    ChatPolicy Chatmediapolicy `json:"chatPolicy"`


    // EmailPolicy - Conditions and actions for emails
    EmailPolicy Emailmediapolicy `json:"emailPolicy"`


    // MessagePolicy - Conditions and actions for messages
    MessagePolicy Messagemediapolicy `json:"messagePolicy"`

}

// String returns a JSON representation of the model
func (o *Mediapolicies) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediapolicies) MarshalJSON() ([]byte, error) {
    type Alias Mediapolicies

    if MediapoliciesMarshalled {
        return []byte("{}"), nil
    }
    MediapoliciesMarshalled = true

    return json.Marshal(&struct { 
        CallPolicy Callmediapolicy `json:"callPolicy"`
        
        ChatPolicy Chatmediapolicy `json:"chatPolicy"`
        
        EmailPolicy Emailmediapolicy `json:"emailPolicy"`
        
        MessagePolicy Messagemediapolicy `json:"messagePolicy"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

