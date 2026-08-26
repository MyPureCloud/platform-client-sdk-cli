package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkprefixesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkprefixesrequestDud struct { 
    

}

// Bulkprefixesrequest
type Bulkprefixesrequest struct { 
    // Entities - List of prefixes to save - add or remove
    Entities []Prefix `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Bulkprefixesrequest) String() string {
     o.Entities = []Prefix{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkprefixesrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkprefixesrequest

    if BulkprefixesrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkprefixesrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Prefix `json:"entities"`
        *Alias
    }{

        
        Entities: []Prefix{{}},
        

        Alias: (*Alias)(u),
    })
}

