package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkrelationshipsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkrelationshipsrequestDud struct { 
    

}

// Bulkrelationshipsrequest
type Bulkrelationshipsrequest struct { 
    // Entities
    Entities []Relationship `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkrelationshipsrequest) String() string {
     o.Entities = []Relationship{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkrelationshipsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkrelationshipsrequest

    if BulkrelationshipsrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkrelationshipsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Relationship `json:"entities"`
        *Alias
    }{

        
        Entities: []Relationship{{}},
        

        Alias: (*Alias)(u),
    })
}

