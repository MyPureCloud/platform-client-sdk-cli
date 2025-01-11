package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobupdateDud struct { 
    

}

// Bulkjobupdate
type Bulkjobupdate struct { 
    // State - The destination state of the bulk job.
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Bulkjobupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobupdate) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobupdate

    if BulkjobupdateMarshalled {
        return []byte("{}"), nil
    }
    BulkjobupdateMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

