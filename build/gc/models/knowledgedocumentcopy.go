package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentcopyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentcopyDud struct { 
    


    


    


    


    


    SessionId string `json:"sessionId"`


    


    

}

// Knowledgedocumentcopy
type Knowledgedocumentcopy struct { 
    // DocumentVariationId - The variation of the document whose content was copied.
    DocumentVariationId string `json:"documentVariationId"`


    // DocumentVersionId - The version of the document whose content was copied.
    DocumentVersionId string `json:"documentVersionId"`


    // SearchId - The search that surfaced the document whose content was copied.
    SearchId string `json:"searchId"`


    // QueryType - The type of the query that surfaced the document.
    QueryType string `json:"queryType"`


    // SurfacingMethod - The method how knowledge was surfaced. Article: Full article was shown. Snippet: A snippet from the article was shown. Highlight: A highlighted answer in a snippet was shown.Generative: A generated answer in a snippet was shown.
    SurfacingMethod string `json:"surfacingMethod"`


    


    // ConversationContext - Conversation context information, if the document content is copied in the context of a conversation.
    ConversationContext Knowledgeconversationcontext `json:"conversationContext"`


    // Application - The client application in which the document content was copied.
    Application Knowledgesearchclientapplication `json:"application"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentcopy) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentcopy) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentcopy

    if KnowledgedocumentcopyMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentcopyMarshalled = true

    return json.Marshal(&struct {
        
        DocumentVariationId string `json:"documentVariationId"`
        
        DocumentVersionId string `json:"documentVersionId"`
        
        SearchId string `json:"searchId"`
        
        QueryType string `json:"queryType"`
        
        SurfacingMethod string `json:"surfacingMethod"`
        
        ConversationContext Knowledgeconversationcontext `json:"conversationContext"`
        
        Application Knowledgesearchclientapplication `json:"application"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

