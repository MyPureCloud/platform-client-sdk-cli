package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentsuggestionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentsuggestionDud struct { 
    


    


    SessionId string `json:"sessionId"`


    Results []Knowledgeguestdocumentsuggestionresult `json:"results"`

}

// Knowledgeguestdocumentsuggestion
type Knowledgeguestdocumentsuggestion struct { 
    // Query - Query to get autocomplete suggestions for the matching knowledge documents.
    Query string `json:"query"`


    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentsuggestion) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentsuggestion) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentsuggestion

    if KnowledgeguestdocumentsuggestionMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentsuggestionMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

