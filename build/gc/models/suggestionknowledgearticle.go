package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestionknowledgearticleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestionknowledgearticleDud struct { 
    Title string `json:"title"`


    Snippets []string `json:"snippets"`


    Document Addressableentityref `json:"document"`


    Version Addressableentityref `json:"version"`


    KnowledgeAnswer Suggestionknowledgeanswer `json:"knowledgeAnswer"`


    Variations []Addressableentityref `json:"variations"`

}

// Suggestionknowledgearticle
type Suggestionknowledgearticle struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Suggestionknowledgearticle) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestionknowledgearticle) MarshalJSON() ([]byte, error) {
    type Alias Suggestionknowledgearticle

    if SuggestionknowledgearticleMarshalled {
        return []byte("{}"), nil
    }
    SuggestionknowledgearticleMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

