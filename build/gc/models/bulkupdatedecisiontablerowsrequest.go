package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdatedecisiontablerowsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdatedecisiontablerowsrequestDud struct { 
    

}

// Bulkupdatedecisiontablerowsrequest
type Bulkupdatedecisiontablerowsrequest struct { 
    // Rows - The list of rows to update. Maximum 15 rows per request. Each row must have a unique rowId.
    Rows []Row `json:"rows"`

}

// String returns a JSON representation of the model
func (o *Bulkupdatedecisiontablerowsrequest) String() string {
     o.Rows = []Row{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdatedecisiontablerowsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdatedecisiontablerowsrequest

    if BulkupdatedecisiontablerowsrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkupdatedecisiontablerowsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Rows []Row `json:"rows"`
        *Alias
    }{

        
        Rows: []Row{{}},
        

        Alias: (*Alias)(u),
    })
}

