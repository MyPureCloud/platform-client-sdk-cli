package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnectionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnectionlistingDud struct { 
    

}

// Connectionlisting
type Connectionlisting struct { 
    // Entities
    Entities []Connectionresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Connectionlisting) String() string {
     o.Entities = []Connectionresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connectionlisting) MarshalJSON() ([]byte, error) {
    type Alias Connectionlisting

    if ConnectionlistingMarshalled {
        return []byte("{}"), nil
    }
    ConnectionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Connectionresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Connectionresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

