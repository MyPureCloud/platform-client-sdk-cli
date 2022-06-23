package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulejobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulejobrequestDud struct { 
    

}

// Learningmodulejobrequest - Learning module job request
type Learningmodulejobrequest struct { 
    // Action - The type for the learning module job
    Action string `json:"action"`

}

// String returns a JSON representation of the model
func (o *Learningmodulejobrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulejobrequest) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulejobrequest

    if LearningmodulejobrequestMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulejobrequestMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

