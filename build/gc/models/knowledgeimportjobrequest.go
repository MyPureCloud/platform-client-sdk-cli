package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeimportjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeimportjobrequestDud struct { 
    


    


    

}

// Knowledgeimportjobrequest
type Knowledgeimportjobrequest struct { 
    // UploadKey - Upload key
    UploadKey string `json:"uploadKey"`


    // FileType - File type of the document
    FileType string `json:"fileType"`


    // Settings - Additional optional settings
    Settings Knowledgeimportjobsettings `json:"settings"`

}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeimportjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeimportjobrequest

    if KnowledgeimportjobrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeimportjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        
        FileType string `json:"fileType"`
        
        Settings Knowledgeimportjobsettings `json:"settings"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

