package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FileuploadmodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FileuploadmodeDud struct { 
    


    

}

// Fileuploadmode
type Fileuploadmode struct { 
    // FileTypes - A list of supported content types for uploading files
    FileTypes []string `json:"fileTypes"`


    // MaxFileSizeKB - The maximum file size for file uploads in kilobytes. Default is 10240 (10 MB)
    MaxFileSizeKB int `json:"maxFileSizeKB"`

}

// String returns a JSON representation of the model
func (o *Fileuploadmode) String() string {
    
    
     o.FileTypes = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fileuploadmode) MarshalJSON() ([]byte, error) {
    type Alias Fileuploadmode

    if FileuploadmodeMarshalled {
        return []byte("{}"), nil
    }
    FileuploadmodeMarshalled = true

    return json.Marshal(&struct { 
        FileTypes []string `json:"fileTypes"`
        
        MaxFileSizeKB int `json:"maxFileSizeKB"`
        
        *Alias
    }{
        

        
        FileTypes: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

