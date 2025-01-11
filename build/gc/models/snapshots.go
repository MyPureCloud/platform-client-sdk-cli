package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SnapshotsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SnapshotsDud struct { 
    


    

}

// Snapshots
type Snapshots struct { 
    // Id - The snapshot Id
    Id string `json:"id"`


    // DaysInPast - The number of days from today denoting when the snapshot was captured
    DaysInPast int `json:"daysInPast"`

}

// String returns a JSON representation of the model
func (o *Snapshots) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Snapshots) MarshalJSON() ([]byte, error) {
    type Alias Snapshots

    if SnapshotsMarshalled {
        return []byte("{}"), nil
    }
    SnapshotsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DaysInPast int `json:"daysInPast"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

