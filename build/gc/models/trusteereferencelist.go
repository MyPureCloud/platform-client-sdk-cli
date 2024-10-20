package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrusteereferencelistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrusteereferencelistDud struct { 
    

}

// Trusteereferencelist
type Trusteereferencelist struct { 
    // Entities
    Entities []Addressableentityref `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Trusteereferencelist) String() string {
     o.Entities = []Addressableentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trusteereferencelist) MarshalJSON() ([]byte, error) {
    type Alias Trusteereferencelist

    if TrusteereferencelistMarshalled {
        return []byte("{}"), nil
    }
    TrusteereferencelistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Addressableentityref `json:"entities"`
        *Alias
    }{

        
        Entities: []Addressableentityref{{}},
        

        Alias: (*Alias)(u),
    })
}

