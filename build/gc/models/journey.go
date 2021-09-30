package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyDud struct { 
    

}

// Journey
type Journey struct { 
    // Patterns - A list of zero or more patterns to match.
    Patterns []Journeypattern `json:"patterns"`

}

// String returns a JSON representation of the model
func (o *Journey) String() string {
    
    
     o.Patterns = []Journeypattern{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journey) MarshalJSON() ([]byte, error) {
    type Alias Journey

    if JourneyMarshalled {
        return []byte("{}"), nil
    }
    JourneyMarshalled = true

    return json.Marshal(&struct { 
        Patterns []Journeypattern `json:"patterns"`
        
        *Alias
    }{
        

        
        Patterns: []Journeypattern{{}},
        

        
        Alias: (*Alias)(u),
    })
}

