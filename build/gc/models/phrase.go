package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhraseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhraseDud struct { 
    


    


    

}

// Phrase
type Phrase struct { 
    // Text - The phrase text
    Text string `json:"text"`


    // Strictness - The phrase strictness, default value is null
    Strictness string `json:"strictness"`


    // Sentiment - The phrase sentiment, default value is Unspecified.
    Sentiment string `json:"sentiment"`

}

// String returns a JSON representation of the model
func (o *Phrase) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phrase) MarshalJSON() ([]byte, error) {
    type Alias Phrase

    if PhraseMarshalled {
        return []byte("{}"), nil
    }
    PhraseMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Strictness string `json:"strictness"`
        
        Sentiment string `json:"sentiment"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

