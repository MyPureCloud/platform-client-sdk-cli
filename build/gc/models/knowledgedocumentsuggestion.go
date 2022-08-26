package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentsuggestionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentsuggestionDud struct { 
    


    


    Results []Knowledgedocumentsuggestionresult `json:"results"`

}

// Knowledgedocumentsuggestion
type Knowledgedocumentsuggestion struct { 
    // Query - Query to get autocomplete suggestions for the matching knowledge documents.
    Query string `json:"query"`


    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsuggestion) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentsuggestion) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentsuggestion

    if KnowledgedocumentsuggestionMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentsuggestionMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

