package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobaddrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobaddrequestDud struct { 
    

}

// Bulkjobaddrequest
type Bulkjobaddrequest struct { 
    // Entities - The list of workitem entities to create.
    Entities []Workitemcommoncreate `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkjobaddrequest) String() string {
     o.Entities = []Workitemcommoncreate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobaddrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobaddrequest

    if BulkjobaddrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkjobaddrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitemcommoncreate `json:"entities"`
        *Alias
    }{

        
        Entities: []Workitemcommoncreate{{}},
        

        Alias: (*Alias)(u),
    })
}

