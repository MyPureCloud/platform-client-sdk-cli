package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgebaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgebaseDud struct { 
    Id string `json:"id"`


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    FaqCount int `json:"faqCount"`


    DateDocumentLastModified time.Time `json:"dateDocumentLastModified"`


    ArticleCount int `json:"articleCount"`


    Published bool `json:"published"`


    SelfUri string `json:"selfUri"`

}

// Knowledgebase
type Knowledgebase struct { 
    


    // Name
    Name string `json:"name"`


    // Description - Knowledge base description
    Description string `json:"description"`


    // CoreLanguage - Core language for knowledge base in which initial content must be created, language codes [en-US, en-UK, en-AU, de-DE] are supported currently. However, the new DX knowledge will support all these language codes, along with 'early preview' language codes [ca-ES, tr-TR, sv-SE, fi-FI, nb-NO, da-DK, ja-JP, ar-AE, zh-CN, zh-TW, zh-HK, ko-KR, pl-PL, hi-IN, th-TH, hu-HU, vi-VN, uk-UA] which might have a lower accuracy.
    CoreLanguage string `json:"coreLanguage"`


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgebase) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgebase) MarshalJSON() ([]byte, error) {
    type Alias Knowledgebase

    if KnowledgebaseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgebaseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        CoreLanguage string `json:"coreLanguage"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

