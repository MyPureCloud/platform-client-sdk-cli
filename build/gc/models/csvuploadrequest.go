package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CsvuploadrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CsvuploadrequestDud struct { 
    


    

}

// Csvuploadrequest
type Csvuploadrequest struct { 
    // FileName - Name of the file to upload
    FileName string `json:"fileName"`


    // FileSize - The size of the upload file
    FileSize int `json:"fileSize"`

}

// String returns a JSON representation of the model
func (o *Csvuploadrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Csvuploadrequest) MarshalJSON() ([]byte, error) {
    type Alias Csvuploadrequest

    if CsvuploadrequestMarshalled {
        return []byte("{}"), nil
    }
    CsvuploadrequestMarshalled = true

    return json.Marshal(&struct {
        
        FileName string `json:"fileName"`
        
        FileSize int `json:"fileSize"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

