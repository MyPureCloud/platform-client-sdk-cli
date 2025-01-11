package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobentityDud struct { 
    

}

// Bulkjobentity
type Bulkjobentity struct { 
    // Id - The id of the bulk job entity.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Bulkjobentity) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobentity) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobentity

    if BulkjobentityMarshalled {
        return []byte("{}"), nil
    }
    BulkjobentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

