package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepublishrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepublishrequestDud struct { 
    

}

// Learningmodulepublishrequest - Learning module publish request
type Learningmodulepublishrequest struct { 
    // TermsAndConditionsAccepted - Whether the terms and conditions were accepted
    TermsAndConditionsAccepted bool `json:"termsAndConditionsAccepted"`

}

// String returns a JSON representation of the model
func (o *Learningmodulepublishrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepublishrequest) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepublishrequest

    if LearningmodulepublishrequestMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepublishrequestMarshalled = true

    return json.Marshal(&struct {
        
        TermsAndConditionsAccepted bool `json:"termsAndConditionsAccepted"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

