package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkopportunitieserrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkopportunitieserrorDud struct { 
    

}

// Bulkopportunitieserror
type Bulkopportunitieserror struct { 
    // ErrorCode - The error code for the failed operation
    ErrorCode string `json:"errorCode"`

}

// String returns a JSON representation of the model
func (o *Bulkopportunitieserror) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkopportunitieserror) MarshalJSON() ([]byte, error) {
    type Alias Bulkopportunitieserror

    if BulkopportunitieserrorMarshalled {
        return []byte("{}"), nil
    }
    BulkopportunitieserrorMarshalled = true

    return json.Marshal(&struct {
        
        ErrorCode string `json:"errorCode"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

