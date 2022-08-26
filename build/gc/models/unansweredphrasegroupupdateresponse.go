package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnansweredphrasegroupupdateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnansweredphrasegroupupdateresponseDud struct { 
    


    

}

// Unansweredphrasegroupupdateresponse
type Unansweredphrasegroupupdateresponse struct { 
    // PhraseAssociations - List of phrases and documents linked in the patch request
    PhraseAssociations []Phraseassociations `json:"phraseAssociations"`


    // Group - Knowledge base unanswered group response
    Group Unansweredgroup `json:"group"`

}

// String returns a JSON representation of the model
func (o *Unansweredphrasegroupupdateresponse) String() string {
     o.PhraseAssociations = []Phraseassociations{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unansweredphrasegroupupdateresponse) MarshalJSON() ([]byte, error) {
    type Alias Unansweredphrasegroupupdateresponse

    if UnansweredphrasegroupupdateresponseMarshalled {
        return []byte("{}"), nil
    }
    UnansweredphrasegroupupdateresponseMarshalled = true

    return json.Marshal(&struct {
        
        PhraseAssociations []Phraseassociations `json:"phraseAssociations"`
        
        Group Unansweredgroup `json:"group"`
        *Alias
    }{

        
        PhraseAssociations: []Phraseassociations{{}},
        


        

        Alias: (*Alias)(u),
    })
}

