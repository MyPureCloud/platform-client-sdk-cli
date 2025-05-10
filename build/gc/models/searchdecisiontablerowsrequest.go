package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchdecisiontablerowsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchdecisiontablerowsrequestDud struct { 
    

}

// Searchdecisiontablerowsrequest
type Searchdecisiontablerowsrequest struct { 
    // Filter - The filter criteria for searching decision table rows
    Filter Rowsearchfilter `json:"filter"`

}

// String returns a JSON representation of the model
func (o *Searchdecisiontablerowsrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchdecisiontablerowsrequest) MarshalJSON() ([]byte, error) {
    type Alias Searchdecisiontablerowsrequest

    if SearchdecisiontablerowsrequestMarshalled {
        return []byte("{}"), nil
    }
    SearchdecisiontablerowsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Filter Rowsearchfilter `json:"filter"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

