package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CsvuploadresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CsvuploadresponseDud struct { 
    


    


    

}

// Csvuploadresponse
type Csvuploadresponse struct { 
    // UploadId - Id of the upload
    UploadId string `json:"uploadId"`


    // UploadUrl - Url for the upload
    UploadUrl string `json:"uploadUrl"`


    // UploadHeaders - Required headers for the upload
    UploadHeaders []Header `json:"uploadHeaders"`

}

// String returns a JSON representation of the model
func (o *Csvuploadresponse) String() string {
    
    
     o.UploadHeaders = []Header{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Csvuploadresponse) MarshalJSON() ([]byte, error) {
    type Alias Csvuploadresponse

    if CsvuploadresponseMarshalled {
        return []byte("{}"), nil
    }
    CsvuploadresponseMarshalled = true

    return json.Marshal(&struct {
        
        UploadId string `json:"uploadId"`
        
        UploadUrl string `json:"uploadUrl"`
        
        UploadHeaders []Header `json:"uploadHeaders"`
        *Alias
    }{

        


        


        
        UploadHeaders: []Header{{}},
        

        Alias: (*Alias)(u),
    })
}

