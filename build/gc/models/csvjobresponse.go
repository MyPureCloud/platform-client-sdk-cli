package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CsvjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CsvjobresponseDud struct { 
    

}

// Csvjobresponse
type Csvjobresponse struct { 
    // Job - Job for the import
    Job Addressableentityref `json:"job"`

}

// String returns a JSON representation of the model
func (o *Csvjobresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Csvjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Csvjobresponse

    if CsvjobresponseMarshalled {
        return []byte("{}"), nil
    }
    CsvjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Addressableentityref `json:"job"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

