package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnansweredphraseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnansweredphraseDud struct { 
    


    


    

}

// Unansweredphrase
type Unansweredphrase struct { 
    // Id - Id of an unanswered phrase
    Id string `json:"id"`


    // Text - Phrase text of an unanswered phrase
    Text string `json:"text"`


    // UnlinkedPhraseHitCount - Hit count of an unlinked phrase
    UnlinkedPhraseHitCount int `json:"unlinkedPhraseHitCount"`

}

// String returns a JSON representation of the model
func (o *Unansweredphrase) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unansweredphrase) MarshalJSON() ([]byte, error) {
    type Alias Unansweredphrase

    if UnansweredphraseMarshalled {
        return []byte("{}"), nil
    }
    UnansweredphraseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        UnlinkedPhraseHitCount int `json:"unlinkedPhraseHitCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

