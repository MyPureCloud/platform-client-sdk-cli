package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PublishscriptrequestdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PublishscriptrequestdataDud struct { 
    


    

}

// Publishscriptrequestdata
type Publishscriptrequestdata struct { 
    // ScriptId - The id of the script to publish
    ScriptId string `json:"scriptId"`


    // VersionId
    VersionId string `json:"versionId"`

}

// String returns a JSON representation of the model
func (o *Publishscriptrequestdata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Publishscriptrequestdata) MarshalJSON() ([]byte, error) {
    type Alias Publishscriptrequestdata

    if PublishscriptrequestdataMarshalled {
        return []byte("{}"), nil
    }
    PublishscriptrequestdataMarshalled = true

    return json.Marshal(&struct {
        
        ScriptId string `json:"scriptId"`
        
        VersionId string `json:"versionId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

