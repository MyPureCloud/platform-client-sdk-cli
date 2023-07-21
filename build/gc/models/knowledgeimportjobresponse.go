package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeimportjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeimportjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    Status string `json:"status"`


    Report Knowledgeimportjobreport `json:"report"`


    KnowledgeBase Knowledgebase `json:"knowledgeBase"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    SelfUri string `json:"selfUri"`

}

// Knowledgeimportjobresponse
type Knowledgeimportjobresponse struct { 
    


    // UploadKey - Upload key
    UploadKey string `json:"uploadKey"`


    // FileType - File type of the document
    FileType string `json:"fileType"`


    // Settings - Additional optional settings
    Settings Knowledgeimportjobsettings `json:"settings"`


    


    


    


    


    


    // SkipConfirmationStep - If enabled pre-validation step will be skipped.
    SkipConfirmationStep bool `json:"skipConfirmationStep"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeimportjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeimportjobresponse

    if KnowledgeimportjobresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeimportjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        
        FileType string `json:"fileType"`
        
        Settings Knowledgeimportjobsettings `json:"settings"`
        
        SkipConfirmationStep bool `json:"skipConfirmationStep"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

