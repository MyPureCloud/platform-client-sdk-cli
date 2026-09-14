package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivitylistingDud struct { 
    

}

// Useractivitylisting
type Useractivitylisting struct { 
    // Entities
    Entities []Useractivity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Useractivitylisting) String() string {
     o.Entities = []Useractivity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivitylisting) MarshalJSON() ([]byte, error) {
    type Alias Useractivitylisting

    if UseractivitylistingMarshalled {
        return []byte("{}"), nil
    }
    UseractivitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Useractivity `json:"entities"`
        *Alias
    }{

        
        Entities: []Useractivity{{}},
        

        Alias: (*Alias)(u),
    })
}

