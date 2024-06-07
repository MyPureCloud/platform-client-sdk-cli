package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmuserentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmuserentitylistingDud struct { 
    

}

// Wfmuserentitylisting
type Wfmuserentitylisting struct { 
    // Entities
    Entities []Userreference `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Wfmuserentitylisting) String() string {
     o.Entities = []Userreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmuserentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Wfmuserentitylisting

    if WfmuserentitylistingMarshalled {
        return []byte("{}"), nil
    }
    WfmuserentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userreference `json:"entities"`
        *Alias
    }{

        
        Entities: []Userreference{{}},
        

        Alias: (*Alias)(u),
    })
}

