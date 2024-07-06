package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestionknowledgeanswerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestionknowledgeanswerDud struct { 
    Answer string `json:"answer"`


    StartIndex int `json:"startIndex"`


    EndIndex int `json:"endIndex"`

}

// Suggestionknowledgeanswer
type Suggestionknowledgeanswer struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Suggestionknowledgeanswer) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestionknowledgeanswer) MarshalJSON() ([]byte, error) {
    type Alias Suggestionknowledgeanswer

    if SuggestionknowledgeanswerMarshalled {
        return []byte("{}"), nil
    }
    SuggestionknowledgeanswerMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

