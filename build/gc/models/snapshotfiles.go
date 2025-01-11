package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SnapshotfilesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SnapshotfilesDud struct { 
    


    


    

}

// Snapshotfiles
type Snapshotfiles struct { 
    // MetaData - Metadata for requested snapshot
    MetaData Snapshotmetadata `json:"metaData"`


    // Offered - Offered data for the requested snapshot
    Offered Snapshotmetricdata `json:"offered"`


    // AverageHandleTime - Average handle time data for the requested snapshot
    AverageHandleTime Snapshotmetricdata `json:"averageHandleTime"`

}

// String returns a JSON representation of the model
func (o *Snapshotfiles) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Snapshotfiles) MarshalJSON() ([]byte, error) {
    type Alias Snapshotfiles

    if SnapshotfilesMarshalled {
        return []byte("{}"), nil
    }
    SnapshotfilesMarshalled = true

    return json.Marshal(&struct {
        
        MetaData Snapshotmetadata `json:"metaData"`
        
        Offered Snapshotmetricdata `json:"offered"`
        
        AverageHandleTime Snapshotmetricdata `json:"averageHandleTime"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

