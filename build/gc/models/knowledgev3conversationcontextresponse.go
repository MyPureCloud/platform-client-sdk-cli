package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Knowledgev3conversationcontextresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Knowledgev3conversationcontextresponseDud struct { 
    


    


    


    

}

// Knowledgev3conversationcontextresponse
type Knowledgev3conversationcontextresponse struct { 
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
func (o *Knowledgev3conversationcontextresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgev3conversationcontextresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgev3conversationcontextresponse

    if Knowledgev3conversationcontextresponseMarshalled {
        return []byte("{}"), nil
    }
    Knowledgev3conversationcontextresponseMarshalled = true

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

