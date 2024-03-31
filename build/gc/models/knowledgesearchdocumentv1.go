package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Knowledgesearchdocumentv1Marshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Knowledgesearchdocumentv1Dud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    KnowledgeBase Knowledgebase `json:"knowledgeBase"`


    


    


    Confidence float64 `json:"confidence"`


    SelfUri string `json:"selfUri"`

}

// Knowledgesearchdocumentv1
type Knowledgesearchdocumentv1 struct { 
    


    // Name
    Name string `json:"name"`


    // LanguageCode - Language of the document
    LanguageCode string `json:"languageCode"`


    // VarType - Document type
    VarType string `json:"type"`


    // Faq - FAQ document details
    Faq Documentfaq `json:"faq"`


    


    


    // Categories - Document categories
    Categories []Knowledgecategory `json:"categories"`


    


    // ExternalUrl - External URL to the document
    ExternalUrl string `json:"externalUrl"`


    // Article - Article
    Article Documentarticle `json:"article"`


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgesearchdocumentv1) String() string {
    
    
    
    
     o.Categories = []Knowledgecategory{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesearchdocumentv1) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesearchdocumentv1

    if Knowledgesearchdocumentv1Marshalled {
        return []byte("{}"), nil
    }
    Knowledgesearchdocumentv1Marshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        LanguageCode string `json:"languageCode"`
        
        VarType string `json:"type"`
        
        Faq Documentfaq `json:"faq"`
        
        Categories []Knowledgecategory `json:"categories"`
        
        ExternalUrl string `json:"externalUrl"`
        
        Article Documentarticle `json:"article"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Categories: []Knowledgecategory{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

