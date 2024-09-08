package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesuggestionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesuggestionconfigDud struct { 
    


    


    

}

// Knowledgesuggestionconfig
type Knowledgesuggestionconfig struct { 
    // VendorName - The name of vendor used for knowledge suggestions.
    VendorName string `json:"vendorName"`


    // KnowledgeBase - The ID of knowledge base to query when Genesys is the knowledge suggestions provider.
    KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`


    // KnowledgeBases - The knowledge bases to query based on dialect, when Genesys is the knowledge suggestions provider.
    KnowledgeBases []Knowledgebasewithdialectreference `json:"knowledgeBases"`

}

// String returns a JSON representation of the model
func (o *Knowledgesuggestionconfig) String() string {
    
    
     o.KnowledgeBases = []Knowledgebasewithdialectreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesuggestionconfig) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesuggestionconfig

    if KnowledgesuggestionconfigMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesuggestionconfigMarshalled = true

    return json.Marshal(&struct {
        
        VendorName string `json:"vendorName"`
        
        KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`
        
        KnowledgeBases []Knowledgebasewithdialectreference `json:"knowledgeBases"`
        *Alias
    }{

        


        


        
        KnowledgeBases: []Knowledgebasewithdialectreference{{}},
        

        Alias: (*Alias)(u),
    })
}

