package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinertopicphraseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinertopicphraseDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Minertopicphrase
type Minertopicphrase struct { 
    


    // Name - Phrase name.
    Name string `json:"name"`


    // Topic - Topic associated with a phrase.
    Topic Minertopic `json:"topic"`


    // Utterances - List of utterances related to a phrase.
    Utterances []Utterance `json:"utterances"`


    // UtteranceCount - Number of utterances belonging to a phrase
    UtteranceCount int `json:"utteranceCount"`


    

}

// String returns a JSON representation of the model
func (o *Minertopicphrase) String() string {
    
    
     o.Utterances = []Utterance{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minertopicphrase) MarshalJSON() ([]byte, error) {
    type Alias Minertopicphrase

    if MinertopicphraseMarshalled {
        return []byte("{}"), nil
    }
    MinertopicphraseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Topic Minertopic `json:"topic"`
        
        Utterances []Utterance `json:"utterances"`
        
        UtteranceCount int `json:"utteranceCount"`
        *Alias
    }{

        


        


        


        
        Utterances: []Utterance{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

