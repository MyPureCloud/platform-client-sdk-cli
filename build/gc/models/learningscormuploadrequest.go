package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningscormuploadrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningscormuploadrequestDud struct { 
    

}

// Learningscormuploadrequest - Learning SCORM upload request
type Learningscormuploadrequest struct { 
    // ContentMd5 - The MD5 content of the SCORM package
    ContentMd5 string `json:"contentMd5"`

}

// String returns a JSON representation of the model
func (o *Learningscormuploadrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningscormuploadrequest) MarshalJSON() ([]byte, error) {
    type Alias Learningscormuploadrequest

    if LearningscormuploadrequestMarshalled {
        return []byte("{}"), nil
    }
    LearningscormuploadrequestMarshalled = true

    return json.Marshal(&struct {
        
        ContentMd5 string `json:"contentMd5"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

