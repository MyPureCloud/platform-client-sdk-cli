package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentsuggestionresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentsuggestionresultDud struct { 
    

}

// Knowledgeguestdocumentsuggestionresult
type Knowledgeguestdocumentsuggestionresult struct { 
    // MatchedPhrase - Matched phrase to the autocomplete suggestions query.
    MatchedPhrase string `json:"matchedPhrase"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentsuggestionresult) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentsuggestionresult) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentsuggestionresult

    if KnowledgeguestdocumentsuggestionresultMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentsuggestionresultMarshalled = true

    return json.Marshal(&struct {
        
        MatchedPhrase string `json:"matchedPhrase"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

