package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DictionaryfeedbackexamplephraseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DictionaryfeedbackexamplephraseDud struct { 
    


    

}

// Dictionaryfeedbackexamplephrase
type Dictionaryfeedbackexamplephrase struct { 
    // Phrase - The Example Phrase text. At least 3 words and up to 20 words
    Phrase string `json:"phrase"`


    // Source - The source of the given Example Phrase
    Source string `json:"source"`

}

// String returns a JSON representation of the model
func (o *Dictionaryfeedbackexamplephrase) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dictionaryfeedbackexamplephrase) MarshalJSON() ([]byte, error) {
    type Alias Dictionaryfeedbackexamplephrase

    if DictionaryfeedbackexamplephraseMarshalled {
        return []byte("{}"), nil
    }
    DictionaryfeedbackexamplephraseMarshalled = true

    return json.Marshal(&struct {
        
        Phrase string `json:"phrase"`
        
        Source string `json:"source"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

