package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequestcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequestcontextDud struct { 
    

}

// Requestcontext
type Requestcontext struct { 
    // Patterns - A list of one or more patterns to match.
    Patterns []Requestcontextpattern `json:"patterns"`

}

// String returns a JSON representation of the model
func (o *Requestcontext) String() string {
     o.Patterns = []Requestcontextpattern{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requestcontext) MarshalJSON() ([]byte, error) {
    type Alias Requestcontext

    if RequestcontextMarshalled {
        return []byte("{}"), nil
    }
    RequestcontextMarshalled = true

    return json.Marshal(&struct {
        
        Patterns []Requestcontextpattern `json:"patterns"`
        *Alias
    }{

        
        Patterns: []Requestcontextpattern{{}},
        

        Alias: (*Alias)(u),
    })
}

