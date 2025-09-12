package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DownloadservicerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DownloadservicerequestDud struct { 
    

}

// Downloadservicerequest
type Downloadservicerequest struct { 
    // Files - List of file names to download
    Files []string `json:"files"`

}

// String returns a JSON representation of the model
func (o *Downloadservicerequest) String() string {
     o.Files = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Downloadservicerequest) MarshalJSON() ([]byte, error) {
    type Alias Downloadservicerequest

    if DownloadservicerequestMarshalled {
        return []byte("{}"), nil
    }
    DownloadservicerequestMarshalled = true

    return json.Marshal(&struct {
        
        Files []string `json:"files"`
        *Alias
    }{

        
        Files: []string{""},
        

        Alias: (*Alias)(u),
    })
}

