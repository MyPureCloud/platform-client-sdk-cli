package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequestjourneyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequestjourneyDud struct { 
    

}

// Requestjourney
type Requestjourney struct { 
    // Patterns - A list of zero or more patterns to match.
    Patterns []Requestjourneypattern `json:"patterns"`

}

// String returns a JSON representation of the model
func (o *Requestjourney) String() string {
     o.Patterns = []Requestjourneypattern{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requestjourney) MarshalJSON() ([]byte, error) {
    type Alias Requestjourney

    if RequestjourneyMarshalled {
        return []byte("{}"), nil
    }
    RequestjourneyMarshalled = true

    return json.Marshal(&struct {
        
        Patterns []Requestjourneypattern `json:"patterns"`
        *Alias
    }{

        
        Patterns: []Requestjourneypattern{{}},
        

        Alias: (*Alias)(u),
    })
}

