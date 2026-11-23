package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DiscoveryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DiscoveryresponseDud struct { 
    

}

// Discoveryresponse
type Discoveryresponse struct { 
    // SnapshotId - The snapshot ID for the discovery job.
    SnapshotId string `json:"snapshotId"`

}

// String returns a JSON representation of the model
func (o *Discoveryresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Discoveryresponse) MarshalJSON() ([]byte, error) {
    type Alias Discoveryresponse

    if DiscoveryresponseMarshalled {
        return []byte("{}"), nil
    }
    DiscoveryresponseMarshalled = true

    return json.Marshal(&struct {
        
        SnapshotId string `json:"snapshotId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

