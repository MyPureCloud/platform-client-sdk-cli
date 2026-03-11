package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Knowledgev3conversationcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Knowledgev3conversationcontextDud struct { 
    


    


    


    


    

}

// Knowledgev3conversationcontext
type Knowledgev3conversationcontext struct { 
    // ConversationId - The unique identifier of the conversation.
    ConversationId string `json:"conversationId"`


    // MediaType - The media type of the conversation.
    MediaType string `json:"mediaType"`


    // MessageType - The message type of the conversation.
    MessageType string `json:"messageType"`


    // QueueId - The unique identifier of the queue used to assign the interaction to the user.
    QueueId string `json:"queueId"`


    // ExternalContactId - The external contact identifier of the end-user participant.
    ExternalContactId string `json:"externalContactId"`

}

// String returns a JSON representation of the model
func (o *Knowledgev3conversationcontext) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgev3conversationcontext) MarshalJSON() ([]byte, error) {
    type Alias Knowledgev3conversationcontext

    if Knowledgev3conversationcontextMarshalled {
        return []byte("{}"), nil
    }
    Knowledgev3conversationcontextMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        
        MediaType string `json:"mediaType"`
        
        MessageType string `json:"messageType"`
        
        QueueId string `json:"queueId"`
        
        ExternalContactId string `json:"externalContactId"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

