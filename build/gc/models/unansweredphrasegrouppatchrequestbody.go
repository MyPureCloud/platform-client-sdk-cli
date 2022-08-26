package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnansweredphrasegrouppatchrequestbodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnansweredphrasegrouppatchrequestbodyDud struct { 
    

}

// Unansweredphrasegrouppatchrequestbody
type Unansweredphrasegrouppatchrequestbody struct { 
    // PhraseAssociations - List of phrases and documents to be linked
    PhraseAssociations []Phraseassociations `json:"phraseAssociations"`

}

// String returns a JSON representation of the model
func (o *Unansweredphrasegrouppatchrequestbody) String() string {
     o.PhraseAssociations = []Phraseassociations{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unansweredphrasegrouppatchrequestbody) MarshalJSON() ([]byte, error) {
    type Alias Unansweredphrasegrouppatchrequestbody

    if UnansweredphrasegrouppatchrequestbodyMarshalled {
        return []byte("{}"), nil
    }
    UnansweredphrasegrouppatchrequestbodyMarshalled = true

    return json.Marshal(&struct {
        
        PhraseAssociations []Phraseassociations `json:"phraseAssociations"`
        *Alias
    }{

        
        PhraseAssociations: []Phraseassociations{{}},
        

        Alias: (*Alias)(u),
    })
}

