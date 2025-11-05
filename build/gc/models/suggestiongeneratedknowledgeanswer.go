package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestiongeneratedknowledgeanswerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestiongeneratedknowledgeanswerDud struct { 
    SearchId string `json:"searchId"`


    KnowledgeAnswerGenerated string `json:"knowledgeAnswerGenerated"`


    SuggestedSearchChunks []Suggestedsearchchunk `json:"suggestedSearchChunks"`

}

// Suggestiongeneratedknowledgeanswer
type Suggestiongeneratedknowledgeanswer struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Suggestiongeneratedknowledgeanswer) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestiongeneratedknowledgeanswer) MarshalJSON() ([]byte, error) {
    type Alias Suggestiongeneratedknowledgeanswer

    if SuggestiongeneratedknowledgeanswerMarshalled {
        return []byte("{}"), nil
    }
    SuggestiongeneratedknowledgeanswerMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

