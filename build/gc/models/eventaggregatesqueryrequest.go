package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventaggregatesqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventaggregatesqueryrequestDud struct { 
    

}

// Eventaggregatesqueryrequest
type Eventaggregatesqueryrequest struct { 
    // Interval - Date and time range to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`

}

// String returns a JSON representation of the model
func (o *Eventaggregatesqueryrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventaggregatesqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Eventaggregatesqueryrequest

    if EventaggregatesqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    EventaggregatesqueryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

