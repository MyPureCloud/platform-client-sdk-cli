package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgebasereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgebasereferenceDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgebasereference
type Knowledgebasereference struct { 
    // Id - The globally unique identifier for the knowledge base.
    Id string `json:"id"`


    // LanguageCode - Language of the knowledge base
    LanguageCode string `json:"languageCode"`


    // ContentSearchEnabled - Flag that indicates the search on content is enabled for the knowledge base.
    ContentSearchEnabled bool `json:"contentSearchEnabled"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgebasereference) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgebasereference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgebasereference

    if KnowledgebasereferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgebasereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        LanguageCode string `json:"languageCode"`
        
        ContentSearchEnabled bool `json:"contentSearchEnabled"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

