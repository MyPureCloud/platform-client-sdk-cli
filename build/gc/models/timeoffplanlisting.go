package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffplanlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffplanlistingDud struct { 
    

}

// Timeoffplanlisting
type Timeoffplanlisting struct { 
    // Entities
    Entities []Timeoffplan `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Timeoffplanlisting) String() string {
     o.Entities = []Timeoffplan{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffplanlisting) MarshalJSON() ([]byte, error) {
    type Alias Timeoffplanlisting

    if TimeoffplanlistingMarshalled {
        return []byte("{}"), nil
    }
    TimeoffplanlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Timeoffplan `json:"entities"`
        *Alias
    }{

        
        Entities: []Timeoffplan{{}},
        

        Alias: (*Alias)(u),
    })
}

