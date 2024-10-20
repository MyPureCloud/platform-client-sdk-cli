package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CsvuploaddetailsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CsvuploaddetailsresponseDud struct { 
    


    


    

}

// Csvuploaddetailsresponse
type Csvuploaddetailsresponse struct { 
    // UploadId - Id for the upload
    UploadId string `json:"uploadId"`


    // FileName - File name for the upload
    FileName string `json:"fileName"`


    // ValidationResult - Validation for the upload
    ValidationResult Validationresult `json:"validationResult"`

}

// String returns a JSON representation of the model
func (o *Csvuploaddetailsresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Csvuploaddetailsresponse) MarshalJSON() ([]byte, error) {
    type Alias Csvuploaddetailsresponse

    if CsvuploaddetailsresponseMarshalled {
        return []byte("{}"), nil
    }
    CsvuploaddetailsresponseMarshalled = true

    return json.Marshal(&struct {
        
        UploadId string `json:"uploadId"`
        
        FileName string `json:"fileName"`
        
        ValidationResult Validationresult `json:"validationResult"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

