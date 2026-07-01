package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesearchpreviewresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesearchpreviewresponseDud struct { 
    


    SearchId string `json:"searchId"`


    SessionId string `json:"sessionId"`


    Result Knowledgesearchresult `json:"result"`


    


    

}

// Knowledgesearchpreviewresponse
type Knowledgesearchpreviewresponse struct { 
    // Query - Query to search content in the knowledge base.
    Query string `json:"query"`


    


    


    


    // Application - The touchpoint application used for the preview.
    Application V3knowledgesearchpreviewclientapplication `json:"application"`


    // ConversationContext - The channel context used for the preview.
    ConversationContext Knowledgev3previewconversationcontext `json:"conversationContext"`

}

// String returns a JSON representation of the model
func (o *Knowledgesearchpreviewresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesearchpreviewresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesearchpreviewresponse

    if KnowledgesearchpreviewresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesearchpreviewresponseMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        Application V3knowledgesearchpreviewclientapplication `json:"application"`
        
        ConversationContext Knowledgev3previewconversationcontext `json:"conversationContext"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

