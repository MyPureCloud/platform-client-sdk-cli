package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GrammarlanguagefilemetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GrammarlanguagefilemetadataDud struct { 
    


    


    


    

}

// Grammarlanguagefilemetadata
type Grammarlanguagefilemetadata struct { 
    // FileName - The name of the file as defined by the user
    FileName string `json:"fileName"`


    // FileSizeBytes - The size of the file in bytes
    FileSizeBytes int `json:"fileSizeBytes"`


    // DateUploaded - The date the file was uploaded. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateUploaded time.Time `json:"dateUploaded"`


    // FileType - The extension of the file
    FileType string `json:"fileType"`

}

// String returns a JSON representation of the model
func (o *Grammarlanguagefilemetadata) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Grammarlanguagefilemetadata) MarshalJSON() ([]byte, error) {
    type Alias Grammarlanguagefilemetadata

    if GrammarlanguagefilemetadataMarshalled {
        return []byte("{}"), nil
    }
    GrammarlanguagefilemetadataMarshalled = true

    return json.Marshal(&struct {
        
        FileName string `json:"fileName"`
        
        FileSizeBytes int `json:"fileSizeBytes"`
        
        DateUploaded time.Time `json:"dateUploaded"`
        
        FileType string `json:"fileType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

