package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationuploadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationuploadDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// V3synchronizationupload
type V3synchronizationupload struct { 
    // FileId - The unique identifier for the upload object.
    FileId string `json:"fileId"`


    // FileName - Name of the uploaded file.
    FileName string `json:"fileName"`


    // Metadata - The metadata of the uploaded file
    Metadata V3synchronizationuploadmetadata `json:"metadata"`


    // Synchronization - The synchronization of the file upload.
    Synchronization V3synchronizationref `json:"synchronization"`


    

}

// String returns a JSON representation of the model
func (o *V3synchronizationupload) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationupload) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationupload

    if V3synchronizationuploadMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationuploadMarshalled = true

    return json.Marshal(&struct {
        
        FileId string `json:"fileId"`
        
        FileName string `json:"fileName"`
        
        Metadata V3synchronizationuploadmetadata `json:"metadata"`
        
        Synchronization V3synchronizationref `json:"synchronization"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

