package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OperatorpositionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OperatorpositionDud struct { 
    


    

}

// Operatorposition
type Operatorposition struct { 
    // VoiceSecondsPosition - Number of seconds (for voice interactions) from operand match
    VoiceSecondsPosition int `json:"voiceSecondsPosition"`


    // DigitalWordsPosition - Number of words (for digital interactions) from operand match
    DigitalWordsPosition int `json:"digitalWordsPosition"`

}

// String returns a JSON representation of the model
func (o *Operatorposition) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Operatorposition) MarshalJSON() ([]byte, error) {
    type Alias Operatorposition

    if OperatorpositionMarshalled {
        return []byte("{}"), nil
    }
    OperatorpositionMarshalled = true

    return json.Marshal(&struct {
        
        VoiceSecondsPosition int `json:"voiceSecondsPosition"`
        
        DigitalWordsPosition int `json:"digitalWordsPosition"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

