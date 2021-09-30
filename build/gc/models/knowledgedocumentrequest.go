package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentrequestDud struct { 
    


    


    


    


    

}

// Knowledgedocumentrequest
type Knowledgedocumentrequest struct { 
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

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Categories = []Documentcategoryinput{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentrequest

    if KnowledgedocumentrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentrequestMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        ExternalUrl string `json:"externalUrl"`
        
        Faq Documentfaq `json:"faq"`
        
        Categories []Documentcategoryinput `json:"categories"`
        
        Article Documentarticle `json:"article"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Categories: []Documentcategoryinput{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

