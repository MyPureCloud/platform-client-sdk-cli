package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningscormuploadresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningscormuploadresponseDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Learningscormuploadresponse - Learning SCORM upload response
type Learningscormuploadresponse struct { 
    


    // Status - The status of the SCORM package
    Status string `json:"status"`


    // UploadUrl - The pre-signed URL. Use it with headers below to upload file to S3
    UploadUrl string `json:"uploadUrl"`


    // Headers - The additional headers that need to be included in the upload request
    Headers map[string]string `json:"headers"`


    

}

// String returns a JSON representation of the model
func (o *Learningscormuploadresponse) String() string {
    
    
     o.Headers = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningscormuploadresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningscormuploadresponse

    if LearningscormuploadresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningscormuploadresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        UploadUrl string `json:"uploadUrl"`
        
        Headers map[string]string `json:"headers"`
        *Alias
    }{

        


        


        


        
        Headers: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

