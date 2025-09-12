package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IgnorephrasesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IgnorephrasesrequestDud struct { 
    

}

// Ignorephrasesrequest
type Ignorephrasesrequest struct { 
    // Phrases - List of phrases to be ignored
    Phrases []Ignorephrase `json:"phrases"`

}

// String returns a JSON representation of the model
func (o *Ignorephrasesrequest) String() string {
     o.Phrases = []Ignorephrase{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ignorephrasesrequest) MarshalJSON() ([]byte, error) {
    type Alias Ignorephrasesrequest

    if IgnorephrasesrequestMarshalled {
        return []byte("{}"), nil
    }
    IgnorephrasesrequestMarshalled = true

    return json.Marshal(&struct {
        
        Phrases []Ignorephrase `json:"phrases"`
        *Alias
    }{

        
        Phrases: []Ignorephrase{{}},
        

        Alias: (*Alias)(u),
    })
}

