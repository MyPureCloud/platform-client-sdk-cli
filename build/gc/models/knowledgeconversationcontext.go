package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeconversationcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeconversationcontextDud struct { 
    


    


    


    

}

// Knowledgeconversationcontext
type Knowledgeconversationcontext struct { 
    // ConversationId - The unique identifier of the conversation.
    ConversationId string `json:"conversationId"`


    // MediaType - The media type of the conversation.
    MediaType string `json:"mediaType"`


    // QueueId - The unique identifier of the queue used to assign the interaction to the user.
    QueueId string `json:"queueId"`


    // ExternalContactId - The external contact identifier of the end-user participant.
    ExternalContactId string `json:"externalContactId"`

}

// String returns a JSON representation of the model
func (o *Knowledgeconversationcontext) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeconversationcontext) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeconversationcontext

    if KnowledgeconversationcontextMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeconversationcontextMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        
        MediaType string `json:"mediaType"`
        
        QueueId string `json:"queueId"`
        
        ExternalContactId string `json:"externalContactId"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

