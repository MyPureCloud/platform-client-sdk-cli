package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnansweredphrasegroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnansweredphrasegroupDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Unansweredphrasegroup
type Unansweredphrasegroup struct { 
    


    // Label - Knowledge base phrase group label
    Label string `json:"label"`


    // Phrases - List of unanswered phrases in a phrase group
    Phrases []Unansweredphrase `json:"phrases"`


    // UnlinkedPhraseHitCount - Hit count of the unlinked phrase group
    UnlinkedPhraseHitCount int `json:"unlinkedPhraseHitCount"`


    // UnlinkedPhraseCount - Unique phrase count of the unlinked phrase group
    UnlinkedPhraseCount int `json:"unlinkedPhraseCount"`


    

}

// String returns a JSON representation of the model
func (o *Unansweredphrasegroup) String() string {
    
     o.Phrases = []Unansweredphrase{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unansweredphrasegroup) MarshalJSON() ([]byte, error) {
    type Alias Unansweredphrasegroup

    if UnansweredphrasegroupMarshalled {
        return []byte("{}"), nil
    }
    UnansweredphrasegroupMarshalled = true

    return json.Marshal(&struct {
        
        Label string `json:"label"`
        
        Phrases []Unansweredphrase `json:"phrases"`
        
        UnlinkedPhraseHitCount int `json:"unlinkedPhraseHitCount"`
        
        UnlinkedPhraseCount int `json:"unlinkedPhraseCount"`
        *Alias
    }{

        


        


        
        Phrases: []Unansweredphrase{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

