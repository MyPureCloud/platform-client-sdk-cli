package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkremovesourceintentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkremovesourceintentsrequestDud struct { 
    

}

// Bulkremovesourceintentsrequest
type Bulkremovesourceintentsrequest struct { 
    // Items - List of ids to delete
    Items []string `json:"items"`

}

// String returns a JSON representation of the model
func (o *Bulkremovesourceintentsrequest) String() string {
     o.Items = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkremovesourceintentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkremovesourceintentsrequest

    if BulkremovesourceintentsrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkremovesourceintentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Items []string `json:"items"`
        *Alias
    }{

        
        Items: []string{""},
        

        Alias: (*Alias)(u),
    })
}

