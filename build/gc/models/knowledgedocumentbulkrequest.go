package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentbulkrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentbulkrequestDud struct { 
    


    


    


    


    


    

}

// Knowledgedocumentbulkrequest
type Knowledgedocumentbulkrequest struct { 
    // VarType - Document type according to assigned template
    VarType string `json:"type"`


    // ExternalUrl - External Url to the document
    ExternalUrl string `json:"externalUrl"`


    // Faq - Faq document details
    Faq Documentfaq `json:"faq"`


    // Categories - Document categories
    Categories []Documentcategoryinput `json:"categories"`


    // Article - Article details
    Article Documentarticle `json:"article"`


    // Id - Identifier of document for update. Omit for create new Document.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentbulkrequest) String() string {
    
    
    
     o.Categories = []Documentcategoryinput{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentbulkrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentbulkrequest

    if KnowledgedocumentbulkrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentbulkrequestMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        ExternalUrl string `json:"externalUrl"`
        
        Faq Documentfaq `json:"faq"`
        
        Categories []Documentcategoryinput `json:"categories"`
        
        Article Documentarticle `json:"article"`
        
        Id string `json:"id"`
        *Alias
    }{

        


        


        


        
        Categories: []Documentcategoryinput{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

