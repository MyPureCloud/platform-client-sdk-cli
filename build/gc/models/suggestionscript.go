package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestionscriptMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestionscriptDud struct { 
    Script Addressableentityref `json:"script"`


    Page Addressableentityref `json:"page"`


    Data map[string]string `json:"data"`

}

// Suggestionscript
type Suggestionscript struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Suggestionscript) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestionscript) MarshalJSON() ([]byte, error) {
    type Alias Suggestionscript

    if SuggestionscriptMarshalled {
        return []byte("{}"), nil
    }
    SuggestionscriptMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

