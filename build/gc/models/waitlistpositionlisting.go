package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WaitlistpositionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WaitlistpositionlistingDud struct { 
    

}

// Waitlistpositionlisting
type Waitlistpositionlisting struct { 
    // Entities
    Entities []Waitlistposition `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Waitlistpositionlisting) String() string {
     o.Entities = []Waitlistposition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Waitlistpositionlisting) MarshalJSON() ([]byte, error) {
    type Alias Waitlistpositionlisting

    if WaitlistpositionlistingMarshalled {
        return []byte("{}"), nil
    }
    WaitlistpositionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Waitlistposition `json:"entities"`
        *Alias
    }{

        
        Entities: []Waitlistposition{{}},
        

        Alias: (*Alias)(u),
    })
}

