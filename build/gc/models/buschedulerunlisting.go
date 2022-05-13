package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulerunlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulerunlistingDud struct { 
    

}

// Buschedulerunlisting
type Buschedulerunlisting struct { 
    // Entities
    Entities []Buschedulerun `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Buschedulerunlisting) String() string {
     o.Entities = []Buschedulerun{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulerunlisting) MarshalJSON() ([]byte, error) {
    type Alias Buschedulerunlisting

    if BuschedulerunlistingMarshalled {
        return []byte("{}"), nil
    }
    BuschedulerunlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Buschedulerun `json:"entities"`
        *Alias
    }{

        
        Entities: []Buschedulerun{{}},
        

        Alias: (*Alias)(u),
    })
}

