package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchjourneyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchjourneyDud struct { 
    

}

// Patchjourney
type Patchjourney struct { 
    // Patterns - A list of zero or more patterns to match.
    Patterns []Patchjourneypattern `json:"patterns"`

}

// String returns a JSON representation of the model
func (o *Patchjourney) String() string {
     o.Patterns = []Patchjourneypattern{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchjourney) MarshalJSON() ([]byte, error) {
    type Alias Patchjourney

    if PatchjourneyMarshalled {
        return []byte("{}"), nil
    }
    PatchjourneyMarshalled = true

    return json.Marshal(&struct {
        
        Patterns []Patchjourneypattern `json:"patterns"`
        *Alias
    }{

        
        Patterns: []Patchjourneypattern{{}},
        

        Alias: (*Alias)(u),
    })
}

