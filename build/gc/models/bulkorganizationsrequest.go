package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkorganizationsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkorganizationsrequestDud struct { 
    

}

// Bulkorganizationsrequest
type Bulkorganizationsrequest struct { 
    // Entities
    Entities []Externalorganization `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkorganizationsrequest) String() string {
     o.Entities = []Externalorganization{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkorganizationsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkorganizationsrequest

    if BulkorganizationsrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkorganizationsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Externalorganization `json:"entities"`
        *Alias
    }{

        
        Entities: []Externalorganization{{}},
        

        Alias: (*Alias)(u),
    })
}

