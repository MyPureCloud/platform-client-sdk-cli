package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SnapshotmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SnapshotmetadataDud struct { 
    


    

}

// Snapshotmetadata
type Snapshotmetadata struct { 
    // DownloadUrl - URL to fetch the snapshot meta data information
    DownloadUrl string `json:"downloadUrl"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult Snapshotmetadataresult `json:"downloadResult"`

}

// String returns a JSON representation of the model
func (o *Snapshotmetadata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Snapshotmetadata) MarshalJSON() ([]byte, error) {
    type Alias Snapshotmetadata

    if SnapshotmetadataMarshalled {
        return []byte("{}"), nil
    }
    SnapshotmetadataMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadResult Snapshotmetadataresult `json:"downloadResult"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

