package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestionknowledgesearchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestionknowledgesearchDud struct { 
    Title string `json:"title"`


    Snippets []string `json:"snippets"`


    Confidence float32 `json:"confidence"`


    SearchId string `json:"searchId"`


    Document Addressableentityref `json:"document"`


    Version Addressableentityref `json:"version"`


    KnowledgeAnswer Suggestionknowledgeanswer `json:"knowledgeAnswer"`


    Variations []Addressableentityref `json:"variations"`

}

// Suggestionknowledgesearch
type Suggestionknowledgesearch struct { 
    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Suggestionknowledgesearch) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestionknowledgesearch) MarshalJSON() ([]byte, error) {
    type Alias Suggestionknowledgesearch

    if SuggestionknowledgesearchMarshalled {
        return []byte("{}"), nil
    }
    SuggestionknowledgesearchMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

