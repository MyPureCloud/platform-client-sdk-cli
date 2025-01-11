package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContinuousforecastsnapshotresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContinuousforecastsnapshotresponseDud struct { 
    


    


    

}

// Continuousforecastsnapshotresponse
type Continuousforecastsnapshotresponse struct { 
    // SessionId - Session Id of the continuous forecast
    SessionId string `json:"sessionId"`


    // SnapshotId - Snapshot Id of the continuous forecast session
    SnapshotId string `json:"snapshotId"`


    // Files - Link to the files containing data for requested snapshot
    Files Snapshotfiles `json:"files"`

}

// String returns a JSON representation of the model
func (o *Continuousforecastsnapshotresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Continuousforecastsnapshotresponse) MarshalJSON() ([]byte, error) {
    type Alias Continuousforecastsnapshotresponse

    if ContinuousforecastsnapshotresponseMarshalled {
        return []byte("{}"), nil
    }
    ContinuousforecastsnapshotresponseMarshalled = true

    return json.Marshal(&struct {
        
        SessionId string `json:"sessionId"`
        
        SnapshotId string `json:"snapshotId"`
        
        Files Snapshotfiles `json:"files"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

