package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatedecisiontablesnapshotrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatedecisiontablesnapshotrequestDud struct { 
    


    

}

// Createdecisiontablesnapshotrequest
type Createdecisiontablesnapshotrequest struct { 
    // SnapshotName - Display name for the snapshot
    SnapshotName string `json:"snapshotName"`


    // Notes - Optional notes for the snapshot
    Notes string `json:"notes"`

}

// String returns a JSON representation of the model
func (o *Createdecisiontablesnapshotrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createdecisiontablesnapshotrequest) MarshalJSON() ([]byte, error) {
    type Alias Createdecisiontablesnapshotrequest

    if CreatedecisiontablesnapshotrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatedecisiontablesnapshotrequestMarshalled = true

    return json.Marshal(&struct {
        
        SnapshotName string `json:"snapshotName"`
        
        Notes string `json:"notes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

