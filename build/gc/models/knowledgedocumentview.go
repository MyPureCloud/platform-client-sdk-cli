package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentviewDud struct { 
    


    


    


    


    


    


    


    

}

// Knowledgedocumentview
type Knowledgedocumentview struct { 
    // DocumentVariationId - The variation of the viewed document.
    DocumentVariationId string `json:"documentVariationId"`


    // DocumentVersionId - The version of the viewed document.
    DocumentVersionId string `json:"documentVersionId"`


    // SearchId - The search that surfaced the viewed document.
    SearchId string `json:"searchId"`


    // QueryType - The type of the query that surfaced the document.
    QueryType string `json:"queryType"`


    // SurfacingMethod - The method how knowledge was surfaced. Article: Full article was shown. Snippet: A snippet from the article was shown. Highlight: A highlighted answer in a snippet was shown.
    SurfacingMethod string `json:"surfacingMethod"`


    // Application - The client application from which the document was viewed.
    Application Knowledgesearchclientapplication `json:"application"`


    // SessionId - The unique identifier of the knowledge session in which the document was viewed.
    SessionId string `json:"sessionId"`


    // ConversationContext - Conversation context information if the document was viewed in the context of a conversation.
    ConversationContext Knowledgeconversationcontext `json:"conversationContext"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentview) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentview) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentview

    if KnowledgedocumentviewMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentviewMarshalled = true

    return json.Marshal(&struct {
        
        DocumentVariationId string `json:"documentVariationId"`
        
        DocumentVersionId string `json:"documentVersionId"`
        
        SearchId string `json:"searchId"`
        
        QueryType string `json:"queryType"`
        
        SurfacingMethod string `json:"surfacingMethod"`
        
        Application Knowledgesearchclientapplication `json:"application"`
        
        SessionId string `json:"sessionId"`
        
        ConversationContext Knowledgeconversationcontext `json:"conversationContext"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

