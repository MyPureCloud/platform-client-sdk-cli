package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulecoverartrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulecoverartrequestDud struct { 
    

}

// Learningmodulecoverartrequest
type Learningmodulecoverartrequest struct { 
    // Id - The key identifier for the cover art
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Learningmodulecoverartrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulecoverartrequest) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulecoverartrequest

    if LearningmodulecoverartrequestMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulecoverartrequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

