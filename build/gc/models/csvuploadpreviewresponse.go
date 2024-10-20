package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CsvuploadpreviewresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CsvuploadpreviewresponseDud struct { 
    


    


    

}

// Csvuploadpreviewresponse
type Csvuploadpreviewresponse struct { 
    // UploadId - Id for the upload
    UploadId string `json:"uploadId"`


    // Headers - List of headers in the CSV
    Headers []string `json:"headers"`


    // Entries - List of entries in the CSV
    Entries [][]string `json:"entries"`

}

// String returns a JSON representation of the model
func (o *Csvuploadpreviewresponse) String() string {
    
     o.Headers = []string{""} 
     o.Entries = [][]string{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Csvuploadpreviewresponse) MarshalJSON() ([]byte, error) {
    type Alias Csvuploadpreviewresponse

    if CsvuploadpreviewresponseMarshalled {
        return []byte("{}"), nil
    }
    CsvuploadpreviewresponseMarshalled = true

    return json.Marshal(&struct {
        
        UploadId string `json:"uploadId"`
        
        Headers []string `json:"headers"`
        
        Entries [][]string `json:"entries"`
        *Alias
    }{

        


        
        Headers: []string{""},
        


        
        Entries: [][]string{{}},
        

        Alias: (*Alias)(u),
    })
}

