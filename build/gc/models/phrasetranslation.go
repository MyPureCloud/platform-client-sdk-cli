package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhrasetranslationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhrasetranslationDud struct { 
    


    


    

}

// Phrasetranslation
type Phrasetranslation struct { 
    // StartTimeMs - Epoch start time of the phrase
    StartTimeMs int `json:"startTimeMs"`


    // ParticipantPurpose - Purpose of the participant associated with the phrase
    ParticipantPurpose string `json:"participantPurpose"`


    // TranslatedText - Translation of the phrase
    TranslatedText string `json:"translatedText"`

}

// String returns a JSON representation of the model
func (o *Phrasetranslation) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phrasetranslation) MarshalJSON() ([]byte, error) {
    type Alias Phrasetranslation

    if PhrasetranslationMarshalled {
        return []byte("{}"), nil
    }
    PhrasetranslationMarshalled = true

    return json.Marshal(&struct {
        
        StartTimeMs int `json:"startTimeMs"`
        
        ParticipantPurpose string `json:"participantPurpose"`
        
        TranslatedText string `json:"translatedText"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

