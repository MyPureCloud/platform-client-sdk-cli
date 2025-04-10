package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkcontactsenrichrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkcontactsenrichrequestDud struct { 
    

}

// Bulkcontactsenrichrequest
type Bulkcontactsenrichrequest struct { 
    // Entities
    Entities []Contactenrichrequest `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkcontactsenrichrequest) String() string {
     o.Entities = []Contactenrichrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkcontactsenrichrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkcontactsenrichrequest

    if BulkcontactsenrichrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkcontactsenrichrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contactenrichrequest `json:"entities"`
        *Alias
    }{

        
        Entities: []Contactenrichrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

