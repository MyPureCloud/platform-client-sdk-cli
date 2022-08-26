package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentsuggestionresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentsuggestionresultDud struct { 
    


    

}

// Knowledgedocumentsuggestionresult
type Knowledgedocumentsuggestionresult struct { 
    // MatchedPhrase - Matched phrase to the autocomplete suggestions query.
    MatchedPhrase string `json:"matchedPhrase"`


    // Document
    Document Knowledgedocumentsuggestionresultdocument `json:"document"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsuggestionresult) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentsuggestionresult) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentsuggestionresult

    if KnowledgedocumentsuggestionresultMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentsuggestionresultMarshalled = true

    return json.Marshal(&struct {
        
        MatchedPhrase string `json:"matchedPhrase"`
        
        Document Knowledgedocumentsuggestionresultdocument `json:"document"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

