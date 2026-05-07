package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourceexpandablelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourceexpandablelistingDud struct { 
    

}

// V3sourceexpandablelisting
type V3sourceexpandablelisting struct { 
    // Entities
    Entities []V3sourceexpandablelistresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *V3sourceexpandablelisting) String() string {
     o.Entities = []V3sourceexpandablelistresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourceexpandablelisting) MarshalJSON() ([]byte, error) {
    type Alias V3sourceexpandablelisting

    if V3sourceexpandablelistingMarshalled {
        return []byte("{}"), nil
    }
    V3sourceexpandablelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []V3sourceexpandablelistresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []V3sourceexpandablelistresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

