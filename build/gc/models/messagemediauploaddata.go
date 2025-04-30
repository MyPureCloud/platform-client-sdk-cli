package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagemediauploaddataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagemediauploaddataDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Messagemediauploaddata
type Messagemediauploaddata struct { 
    // Id - A unique media ID
    Id string `json:"id"`


    // Name - The name of the file
    Name string `json:"name"`


    // UploadUrl - URL to upload the file
    UploadUrl string `json:"uploadUrl"`


    // UploadHeaders - Required headers when uploading a file with the upload URL
    UploadHeaders map[string]string `json:"uploadHeaders"`


    

}

// String returns a JSON representation of the model
func (o *Messagemediauploaddata) String() string {
    
    
    
     o.UploadHeaders = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagemediauploaddata) MarshalJSON() ([]byte, error) {
    type Alias Messagemediauploaddata

    if MessagemediauploaddataMarshalled {
        return []byte("{}"), nil
    }
    MessagemediauploaddataMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        UploadUrl string `json:"uploadUrl"`
        
        UploadHeaders map[string]string `json:"uploadHeaders"`
        *Alias
    }{

        


        


        


        
        UploadHeaders: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

