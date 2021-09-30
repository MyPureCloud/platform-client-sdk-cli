package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CrossplatformmediapoliciesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CrossplatformmediapoliciesDud struct { 
    


    


    


    

}

// Crossplatformmediapolicies
type Crossplatformmediapolicies struct { 
    // CallPolicy - Conditions and actions for calls
    CallPolicy Crossplatformcallmediapolicy `json:"callPolicy"`


    // ChatPolicy - Conditions and actions for chats
    ChatPolicy Crossplatformchatmediapolicy `json:"chatPolicy"`


    // EmailPolicy - Conditions and actions for emails
    EmailPolicy Crossplatformemailmediapolicy `json:"emailPolicy"`


    // MessagePolicy - Conditions and actions for messages
    MessagePolicy Crossplatformmessagemediapolicy `json:"messagePolicy"`

}

// String returns a JSON representation of the model
func (o *Crossplatformmediapolicies) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Crossplatformmediapolicies) MarshalJSON() ([]byte, error) {
    type Alias Crossplatformmediapolicies

    if CrossplatformmediapoliciesMarshalled {
        return []byte("{}"), nil
    }
    CrossplatformmediapoliciesMarshalled = true

    return json.Marshal(&struct { 
        CallPolicy Crossplatformcallmediapolicy `json:"callPolicy"`
        
        ChatPolicy Crossplatformchatmediapolicy `json:"chatPolicy"`
        
        EmailPolicy Crossplatformemailmediapolicy `json:"emailPolicy"`
        
        MessagePolicy Crossplatformmessagemediapolicy `json:"messagePolicy"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

