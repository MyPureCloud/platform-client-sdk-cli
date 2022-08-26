package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelectedanswerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelectedanswerDud struct { 
    

}

// Selectedanswer
type Selectedanswer struct { 
    // Document - The search result document chosen as the answer.
    Document Addressableentityref `json:"document"`

}

// String returns a JSON representation of the model
func (o *Selectedanswer) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selectedanswer) MarshalJSON() ([]byte, error) {
    type Alias Selectedanswer

    if SelectedanswerMarshalled {
        return []byte("{}"), nil
    }
    SelectedanswerMarshalled = true

    return json.Marshal(&struct {
        
        Document Addressableentityref `json:"document"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

