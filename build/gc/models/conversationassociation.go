package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationassociationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationassociationDud struct { 
    


    


    


    

}

// Conversationassociation
type Conversationassociation struct { 
    // ExternalContactId - An external contact ID.  If not supplied, implies the conversation should be disassociated with any external contact.
    ExternalContactId string `json:"externalContactId"`


    // ConversationId - Conversation ID
    ConversationId string `json:"conversationId"`


    // CommunicationId - Communication ID
    CommunicationId string `json:"communicationId"`


    // MediaType - Media type
    MediaType string `json:"mediaType"`

}

// String returns a JSON representation of the model
func (o *Conversationassociation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationassociation) MarshalJSON() ([]byte, error) {
    type Alias Conversationassociation

    if ConversationassociationMarshalled {
        return []byte("{}"), nil
    }
    ConversationassociationMarshalled = true

    return json.Marshal(&struct { 
        ExternalContactId string `json:"externalContactId"`
        
        ConversationId string `json:"conversationId"`
        
        CommunicationId string `json:"communicationId"`
        
        MediaType string `json:"mediaType"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

