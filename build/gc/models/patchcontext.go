package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchcontextDud struct { 
    

}

// Patchcontext
type Patchcontext struct { 
    // Patterns - A list of one or more patterns to match.
    Patterns []Patchcontextpattern `json:"patterns"`

}

// String returns a JSON representation of the model
func (o *Patchcontext) String() string {
     o.Patterns = []Patchcontextpattern{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchcontext) MarshalJSON() ([]byte, error) {
    type Alias Patchcontext

    if PatchcontextMarshalled {
        return []byte("{}"), nil
    }
    PatchcontextMarshalled = true

    return json.Marshal(&struct {
        
        Patterns []Patchcontextpattern `json:"patterns"`
        *Alias
    }{

        
        Patterns: []Patchcontextpattern{{}},
        

        Alias: (*Alias)(u),
    })
}

