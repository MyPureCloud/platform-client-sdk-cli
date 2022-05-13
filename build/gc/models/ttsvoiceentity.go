package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TtsvoiceentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TtsvoiceentityDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Ttsvoiceentity
type Ttsvoiceentity struct { 
    


    // Name
    Name string `json:"name"`


    // Gender - The gender of the TTS voice
    Gender string `json:"gender"`


    // Language - The language supported by the TTS voice
    Language string `json:"language"`


    // Engine - Ths TTS engine this voice belongs to
    Engine Ttsengineentity `json:"engine"`


    // IsDefault - The voice is the default voice for its language
    IsDefault bool `json:"isDefault"`


    

}

// String returns a JSON representation of the model
func (o *Ttsvoiceentity) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ttsvoiceentity) MarshalJSON() ([]byte, error) {
    type Alias Ttsvoiceentity

    if TtsvoiceentityMarshalled {
        return []byte("{}"), nil
    }
    TtsvoiceentityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Gender string `json:"gender"`
        
        Language string `json:"language"`
        
        Engine Ttsengineentity `json:"engine"`
        
        IsDefault bool `json:"isDefault"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

