package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NlufeedbackrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NlufeedbackrequestDud struct { 
    


    


    


    

}

// Nlufeedbackrequest
type Nlufeedbackrequest struct { 
    // Text - The feedback text.
    Text string `json:"text"`


    // Intents - Detected intent of the utterance
    Intents []Intentfeedback `json:"intents"`


    // VersionId - The domain version ID of the feedback.
    VersionId string `json:"versionId"`


    // Language - The language of the version to which feedback is linked, e.g. en-us, de-de
    Language string `json:"language"`

}

// String returns a JSON representation of the model
func (o *Nlufeedbackrequest) String() string {
    
     o.Intents = []Intentfeedback{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nlufeedbackrequest) MarshalJSON() ([]byte, error) {
    type Alias Nlufeedbackrequest

    if NlufeedbackrequestMarshalled {
        return []byte("{}"), nil
    }
    NlufeedbackrequestMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Intents []Intentfeedback `json:"intents"`
        
        VersionId string `json:"versionId"`
        
        Language string `json:"language"`
        *Alias
    }{

        


        
        Intents: []Intentfeedback{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

