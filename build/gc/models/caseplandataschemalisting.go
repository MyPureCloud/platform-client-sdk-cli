package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseplandataschemalistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseplandataschemalistingDud struct { 
    

}

// Caseplandataschemalisting
type Caseplandataschemalisting struct { 
    // Entities
    Entities []Caseplandataschema `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Caseplandataschemalisting) String() string {
     o.Entities = []Caseplandataschema{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseplandataschemalisting) MarshalJSON() ([]byte, error) {
    type Alias Caseplandataschemalisting

    if CaseplandataschemalistingMarshalled {
        return []byte("{}"), nil
    }
    CaseplandataschemalistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Caseplandataschema `json:"entities"`
        *Alias
    }{

        
        Entities: []Caseplandataschema{{}},
        

        Alias: (*Alias)(u),
    })
}

