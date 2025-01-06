package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataingestionrulestatuspatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataingestionrulestatuspatchrequestDud struct { 
    

}

// Dataingestionrulestatuspatchrequest
type Dataingestionrulestatuspatchrequest struct { 
    // Status - The status of the data ingestion rule.
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Dataingestionrulestatuspatchrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataingestionrulestatuspatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Dataingestionrulestatuspatchrequest

    if DataingestionrulestatuspatchrequestMarshalled {
        return []byte("{}"), nil
    }
    DataingestionrulestatuspatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

