package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionmetadataDud struct { 
    


    

}

// Sessionmetadata
type Sessionmetadata struct { 
    // DownloadUrl - URL to fetch the meta data information. This field is populated only if session state is Complete
    DownloadUrl string `json:"downloadUrl"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult Sessionmetadataresult `json:"downloadResult"`

}

// String returns a JSON representation of the model
func (o *Sessionmetadata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionmetadata) MarshalJSON() ([]byte, error) {
    type Alias Sessionmetadata

    if SessionmetadataMarshalled {
        return []byte("{}"), nil
    }
    SessionmetadataMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadResult Sessionmetadataresult `json:"downloadResult"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

