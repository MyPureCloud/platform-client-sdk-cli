package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulemigrateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulemigrateresponseDud struct { 
    

}

// Learningmodulemigrateresponse - Migrate response
type Learningmodulemigrateresponse struct { 
    // AutoAssign - The autoAssign Response
    AutoAssign Learningmoduleautoassignresponse `json:"autoAssign"`

}

// String returns a JSON representation of the model
func (o *Learningmodulemigrateresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulemigrateresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulemigrateresponse

    if LearningmodulemigrateresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulemigrateresponseMarshalled = true

    return json.Marshal(&struct {
        
        AutoAssign Learningmoduleautoassignresponse `json:"autoAssign"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

