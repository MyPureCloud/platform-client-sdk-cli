package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkadddecisiontablerowsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkadddecisiontablerowsrequestDud struct { 
    

}

// Bulkadddecisiontablerowsrequest
type Bulkadddecisiontablerowsrequest struct { 
    // Rows - The list of rows to create. Maximum 15 rows per request. RowIndex is not supported for bulk add - all rows will be appended to the end of the table in the order provided.
    Rows []Createdecisiontablerowrequest `json:"rows"`

}

// String returns a JSON representation of the model
func (o *Bulkadddecisiontablerowsrequest) String() string {
     o.Rows = []Createdecisiontablerowrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkadddecisiontablerowsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkadddecisiontablerowsrequest

    if BulkadddecisiontablerowsrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkadddecisiontablerowsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Rows []Createdecisiontablerowrequest `json:"rows"`
        *Alias
    }{

        
        Rows: []Createdecisiontablerowrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

