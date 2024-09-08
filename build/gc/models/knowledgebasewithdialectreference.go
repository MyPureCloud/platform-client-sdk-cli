package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgebasewithdialectreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgebasewithdialectreferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgebasewithdialectreference
type Knowledgebasewithdialectreference struct { 
    // Id - The globally unique identifier for the knowledge base.
    Id string `json:"id"`


    // LanguageCode - The dialect for the knowledge base.
    LanguageCode string `json:"languageCode"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgebasewithdialectreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgebasewithdialectreference) MarshalJSON() ([]byte, error) {
    type Alias Knowledgebasewithdialectreference

    if KnowledgebasewithdialectreferenceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgebasewithdialectreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        LanguageCode string `json:"languageCode"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

