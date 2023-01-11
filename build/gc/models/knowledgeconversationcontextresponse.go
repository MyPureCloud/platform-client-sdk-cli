package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeconversationcontextresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeconversationcontextresponseDud struct { 
    


    


    


    

}

// Knowledgeconversationcontextresponse
type Knowledgeconversationcontextresponse struct { 
    // Conversation - The conversation.
    Conversation Addressableentityref `json:"conversation"`


    // Queue - The queue used to assign the interaction to the user.
    Queue Addressableentityref `json:"queue"`


    // ExternalContact - The end-user participant of the conversation.
    ExternalContact Addressableentityref `json:"externalContact"`


    // MediaType - The media type of the conversation.
    MediaType string `json:"mediaType"`

}

// String returns a JSON representation of the model
func (o *Knowledgeconversationcontextresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeconversationcontextresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeconversationcontextresponse

    if KnowledgeconversationcontextresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeconversationcontextresponseMarshalled = true

    return json.Marshal(&struct {
        
        Conversation Addressableentityref `json:"conversation"`
        
        Queue Addressableentityref `json:"queue"`
        
        ExternalContact Addressableentityref `json:"externalContact"`
        
        MediaType string `json:"mediaType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

