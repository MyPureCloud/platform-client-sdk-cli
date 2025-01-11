package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobterminaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobterminaterequestDud struct { 
    

}

// Bulkjobterminaterequest
type Bulkjobterminaterequest struct { 
    // Entities - The list of bulk job entities to terminate.
    Entities []Bulkjobentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkjobterminaterequest) String() string {
     o.Entities = []Bulkjobentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobterminaterequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobterminaterequest

    if BulkjobterminaterequestMarshalled {
        return []byte("{}"), nil
    }
    BulkjobterminaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Bulkjobentity `json:"entities"`
        *Alias
    }{

        
        Entities: []Bulkjobentity{{}},
        

        Alias: (*Alias)(u),
    })
}

