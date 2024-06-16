package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewupdateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewupdateresponseDud struct { 
    


    

}

// Learningmodulepreviewupdateresponse - Learning module preview update response
type Learningmodulepreviewupdateresponse struct { 
    // Id - The Learning Module id
    Id string `json:"id"`


    // Assignment - The Assignment Preview
    Assignment Learningmodulepreviewupdateresponseassignment `json:"assignment"`

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdateresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewupdateresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewupdateresponse

    if LearningmodulepreviewupdateresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewupdateresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Assignment Learningmodulepreviewupdateresponseassignment `json:"assignment"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

