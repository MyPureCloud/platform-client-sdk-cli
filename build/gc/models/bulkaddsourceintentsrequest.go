package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkaddsourceintentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkaddsourceintentsrequestDud struct { 
    

}

// Bulkaddsourceintentsrequest
type Bulkaddsourceintentsrequest struct { 
    // Items - List of items to add
    Items []Sourceintent `json:"items"`

}

// String returns a JSON representation of the model
func (o *Bulkaddsourceintentsrequest) String() string {
     o.Items = []Sourceintent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkaddsourceintentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkaddsourceintentsrequest

    if BulkaddsourceintentsrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkaddsourceintentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Items []Sourceintent `json:"items"`
        *Alias
    }{

        
        Items: []Sourceintent{{}},
        

        Alias: (*Alias)(u),
    })
}

