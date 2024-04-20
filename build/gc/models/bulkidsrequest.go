package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkidsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkidsrequestDud struct { 
    

}

// Bulkidsrequest
type Bulkidsrequest struct { 
    // Entities
    Entities []Externalcontactsentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkidsrequest) String() string {
     o.Entities = []Externalcontactsentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkidsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkidsrequest

    if BulkidsrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkidsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Externalcontactsentity `json:"entities"`
        *Alias
    }{

        
        Entities: []Externalcontactsentity{{}},
        

        Alias: (*Alias)(u),
    })
}

