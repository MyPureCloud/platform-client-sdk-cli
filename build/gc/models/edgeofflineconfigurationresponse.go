package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeofflineconfigurationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeofflineconfigurationresponseDud struct { 
    

}

// Edgeofflineconfigurationresponse
type Edgeofflineconfigurationresponse struct { 
    // DownloadUrl - The URL to download the Edge configuration file. Save the file to a USB key and plug it into the Edge.
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Edgeofflineconfigurationresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeofflineconfigurationresponse) MarshalJSON() ([]byte, error) {
    type Alias Edgeofflineconfigurationresponse

    if EdgeofflineconfigurationresponseMarshalled {
        return []byte("{}"), nil
    }
    EdgeofflineconfigurationresponseMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

