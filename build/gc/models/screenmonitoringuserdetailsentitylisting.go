package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenmonitoringuserdetailsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenmonitoringuserdetailsentitylistingDud struct { 
    

}

// Screenmonitoringuserdetailsentitylisting
type Screenmonitoringuserdetailsentitylisting struct { 
    // Entities
    Entities []Screenmonitoringuserdetails `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Screenmonitoringuserdetailsentitylisting) String() string {
     o.Entities = []Screenmonitoringuserdetails{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenmonitoringuserdetailsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Screenmonitoringuserdetailsentitylisting

    if ScreenmonitoringuserdetailsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ScreenmonitoringuserdetailsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Screenmonitoringuserdetails `json:"entities"`
        *Alias
    }{

        
        Entities: []Screenmonitoringuserdetails{{}},
        

        Alias: (*Alias)(u),
    })
}

