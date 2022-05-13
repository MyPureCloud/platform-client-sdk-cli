package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataavailabilityresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataavailabilityresponseDud struct { 
    

}

// Dataavailabilityresponse
type Dataavailabilityresponse struct { 
    // DataAvailabilityDate - Date and time before which data is guaranteed to be available in the datalake. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DataAvailabilityDate time.Time `json:"dataAvailabilityDate"`

}

// String returns a JSON representation of the model
func (o *Dataavailabilityresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataavailabilityresponse) MarshalJSON() ([]byte, error) {
    type Alias Dataavailabilityresponse

    if DataavailabilityresponseMarshalled {
        return []byte("{}"), nil
    }
    DataavailabilityresponseMarshalled = true

    return json.Marshal(&struct {
        
        DataAvailabilityDate time.Time `json:"dataAvailabilityDate"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

