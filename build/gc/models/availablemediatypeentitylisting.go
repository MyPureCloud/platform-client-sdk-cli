package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailablemediatypeentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailablemediatypeentitylistingDud struct { 
    

}

// Availablemediatypeentitylisting
type Availablemediatypeentitylisting struct { 
    // Entities
    Entities []Availablemediatype `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Availablemediatypeentitylisting) String() string {
     o.Entities = []Availablemediatype{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availablemediatypeentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Availablemediatypeentitylisting

    if AvailablemediatypeentitylistingMarshalled {
        return []byte("{}"), nil
    }
    AvailablemediatypeentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Availablemediatype `json:"entities"`
        *Alias
    }{

        
        Entities: []Availablemediatype{{}},
        

        Alias: (*Alias)(u),
    })
}

