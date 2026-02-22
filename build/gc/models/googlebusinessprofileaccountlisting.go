package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GooglebusinessprofileaccountlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GooglebusinessprofileaccountlistingDud struct { 
    

}

// Googlebusinessprofileaccountlisting
type Googlebusinessprofileaccountlisting struct { 
    // Entities
    Entities []Googlebusinessprofileaccount `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Googlebusinessprofileaccountlisting) String() string {
     o.Entities = []Googlebusinessprofileaccount{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googlebusinessprofileaccountlisting) MarshalJSON() ([]byte, error) {
    type Alias Googlebusinessprofileaccountlisting

    if GooglebusinessprofileaccountlistingMarshalled {
        return []byte("{}"), nil
    }
    GooglebusinessprofileaccountlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Googlebusinessprofileaccount `json:"entities"`
        *Alias
    }{

        
        Entities: []Googlebusinessprofileaccount{{}},
        

        Alias: (*Alias)(u),
    })
}

