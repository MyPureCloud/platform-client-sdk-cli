package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcewitherrorlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcewitherrorlistingDud struct { 
    

}

// V3sourcewitherrorlisting
type V3sourcewitherrorlisting struct { 
    // Entities
    Entities []V3sourcewitherrorresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *V3sourcewitherrorlisting) String() string {
     o.Entities = []V3sourcewitherrorresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcewitherrorlisting) MarshalJSON() ([]byte, error) {
    type Alias V3sourcewitherrorlisting

    if V3sourcewitherrorlistingMarshalled {
        return []byte("{}"), nil
    }
    V3sourcewitherrorlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []V3sourcewitherrorresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []V3sourcewitherrorresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

