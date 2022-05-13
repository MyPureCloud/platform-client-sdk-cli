package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatmessageDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Chatmessage
type Chatmessage struct { 
    // Body - The message body
    Body string `json:"body"`


    // Id
    Id string `json:"id"`


    // To - The message recipient
    To string `json:"to"`


    // From - The message sender
    From string `json:"from"`


    // Utc
    Utc string `json:"utc"`


    // Chat - The interaction id (if available)
    Chat string `json:"chat"`


    // Message - The message id
    Message string `json:"message"`


    // VarType
    VarType string `json:"type"`


    // BodyType - Type of the message body (v2 chats only)
    BodyType string `json:"bodyType"`


    // SenderCommunicationId - Communication of sender (v2 chats only)
    SenderCommunicationId string `json:"senderCommunicationId"`


    // ParticipantPurpose - Participant purpose of sender (v2 chats only)
    ParticipantPurpose string `json:"participantPurpose"`


    // User - The user information for the sender (if available)
    User Chatmessageuser `json:"user"`

}

// String returns a JSON representation of the model
func (o *Chatmessage) String() string {
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatmessage) MarshalJSON() ([]byte, error) {
    type Alias Chatmessage

    if ChatmessageMarshalled {
        return []byte("{}"), nil
    }
    ChatmessageMarshalled = true

    return json.Marshal(&struct {
        
        Body string `json:"body"`
        
        Id string `json:"id"`
        
        To string `json:"to"`
        
        From string `json:"from"`
        
        Utc string `json:"utc"`
        
        Chat string `json:"chat"`
        
        Message string `json:"message"`
        
        VarType string `json:"type"`
        
        BodyType string `json:"bodyType"`
        
        SenderCommunicationId string `json:"senderCommunicationId"`
        
        ParticipantPurpose string `json:"participantPurpose"`
        
        User Chatmessageuser `json:"user"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

