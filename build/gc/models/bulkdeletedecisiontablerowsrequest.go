package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkdeletedecisiontablerowsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkdeletedecisiontablerowsrequestDud struct { 
    

}

// Bulkdeletedecisiontablerowsrequest
type Bulkdeletedecisiontablerowsrequest struct { 
    // RowIds - The set of unique row IDs to be deleted. Maximum 49 rows per request.
    RowIds []string `json:"rowIds"`

}

// String returns a JSON representation of the model
func (o *Bulkdeletedecisiontablerowsrequest) String() string {
     o.RowIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkdeletedecisiontablerowsrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkdeletedecisiontablerowsrequest

    if BulkdeletedecisiontablerowsrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkdeletedecisiontablerowsrequestMarshalled = true

    return json.Marshal(&struct {
        
        RowIds []string `json:"rowIds"`
        *Alias
    }{

        
        RowIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

