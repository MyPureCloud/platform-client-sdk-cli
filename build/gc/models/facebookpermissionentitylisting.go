package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookpermissionentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookpermissionentitylistingDud struct { 
    

}

// Facebookpermissionentitylisting
type Facebookpermissionentitylisting struct { 
    // Entities
    Entities []Facebookpermission `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Facebookpermissionentitylisting) String() string {
     o.Entities = []Facebookpermission{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookpermissionentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Facebookpermissionentitylisting

    if FacebookpermissionentitylistingMarshalled {
        return []byte("{}"), nil
    }
    FacebookpermissionentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Facebookpermission `json:"entities"`
        *Alias
    }{

        
        Entities: []Facebookpermission{{}},
        

        Alias: (*Alias)(u),
    })
}

