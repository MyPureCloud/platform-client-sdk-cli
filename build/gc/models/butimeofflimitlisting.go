package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButimeofflimitlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButimeofflimitlistingDud struct { 
    

}

// Butimeofflimitlisting
type Butimeofflimitlisting struct { 
    // Entities
    Entities []Butimeofflimitresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Butimeofflimitlisting) String() string {
     o.Entities = []Butimeofflimitresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Butimeofflimitlisting) MarshalJSON() ([]byte, error) {
    type Alias Butimeofflimitlisting

    if ButimeofflimitlistingMarshalled {
        return []byte("{}"), nil
    }
    ButimeofflimitlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Butimeofflimitresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Butimeofflimitresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

