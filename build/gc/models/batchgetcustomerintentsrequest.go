package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchgetcustomerintentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchgetcustomerintentsrequestDud struct { 
    

}

// Batchgetcustomerintentsrequest
type Batchgetcustomerintentsrequest struct { 
    // Ids - Customer intent IDs to retrieve
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Batchgetcustomerintentsrequest) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchgetcustomerintentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Batchgetcustomerintentsrequest

    if BatchgetcustomerintentsrequestMarshalled {
        return []byte("{}"), nil
    }
    BatchgetcustomerintentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

