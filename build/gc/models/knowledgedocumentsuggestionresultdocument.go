package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentsuggestionresultdocumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentsuggestionresultdocumentDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgedocumentsuggestionresultdocument
type Knowledgedocumentsuggestionresultdocument struct { 
    // Id - The globally unique identifier for the document.
    Id string `json:"id"`


    // KnowledgeBase - The knowledge base that the document belongs to.
    KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`


    // Title - The title of the document.
    Title string `json:"title"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsuggestionresultdocument) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentsuggestionresultdocument) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentsuggestionresultdocument

    if KnowledgedocumentsuggestionresultdocumentMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentsuggestionresultdocumentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`
        
        Title string `json:"title"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

