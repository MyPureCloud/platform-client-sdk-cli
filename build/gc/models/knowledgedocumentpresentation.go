package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentpresentationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentpresentationDud struct { 
    


    


    


    SessionId string `json:"sessionId"`


    


    

}

// Knowledgedocumentpresentation
type Knowledgedocumentpresentation struct { 
    // Documents - The presented documents
    Documents []Knowledgedocumentversionvariationreference `json:"documents"`


    // SearchId - The search that surfaced the documents that were presented.
    SearchId string `json:"searchId"`


    // QueryType - The type of the query that surfaced the documents.
    QueryType string `json:"queryType"`


    


    // ConversationContext - Conversation context information if the documents were presented in the context of a conversation.
    ConversationContext Knowledgeconversationcontext `json:"conversationContext"`


    // Application - The client application in which the documents were presented.
    Application Knowledgesearchclientapplication `json:"application"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentpresentation) String() string {
     o.Documents = []Knowledgedocumentversionvariationreference{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentpresentation) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentpresentation

    if KnowledgedocumentpresentationMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentpresentationMarshalled = true

    return json.Marshal(&struct {
        
        Documents []Knowledgedocumentversionvariationreference `json:"documents"`
        
        SearchId string `json:"searchId"`
        
        QueryType string `json:"queryType"`
        
        ConversationContext Knowledgeconversationcontext `json:"conversationContext"`
        
        Application Knowledgesearchclientapplication `json:"application"`
        *Alias
    }{

        
        Documents: []Knowledgedocumentversionvariationreference{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

