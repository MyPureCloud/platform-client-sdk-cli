package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentfeedbackupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentfeedbackupdaterequestDud struct { 
    


    


    


    

}

// Knowledgedocumentfeedbackupdaterequest
type Knowledgedocumentfeedbackupdaterequest struct { 
    // Rating - Feedback rating.
    Rating string `json:"rating"`


    // Reason - Feedback reason
    Reason string `json:"reason"`


    // Comment - Feedback comment
    Comment string `json:"comment"`


    // State - Feedback state
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentfeedbackupdaterequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentfeedbackupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentfeedbackupdaterequest

    if KnowledgedocumentfeedbackupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentfeedbackupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Rating string `json:"rating"`
        
        Reason string `json:"reason"`
        
        Comment string `json:"comment"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

