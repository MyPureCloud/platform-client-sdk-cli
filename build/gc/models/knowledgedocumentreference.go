package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentreferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgedocumentreference
type Knowledgedocumentreference struct { 
    // Id - The globally unique identifier for the document.
    Id string `json:"id"`


    // KnowledgeBase - The knowledge base that the document belongs to.
    KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentreference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentreference

    if KnowledgedocumentreferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

