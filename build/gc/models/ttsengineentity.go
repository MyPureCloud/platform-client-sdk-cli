package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TtsengineentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TtsengineentityDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Ttsengineentity
type Ttsengineentity struct { 
    


    // Name
    Name string `json:"name"`


    // Languages - The set of languages the TTS engine supports
    Languages []string `json:"languages"`


    // OutputFormats - The set of output formats the TTS engine can produce
    OutputFormats []string `json:"outputFormats"`


    // Voices - The set of voices the TTS engine supports
    Voices []Ttsvoiceentity `json:"voices"`


    // IsDefault - The TTS engine is the global default engine
    IsDefault bool `json:"isDefault"`


    // IsSecure - The TTS engine can be used in a secure call flow
    IsSecure bool `json:"isSecure"`


    

}

// String returns a JSON representation of the model
func (o *Ttsengineentity) String() string {
    
     o.Languages = []string{""} 
     o.OutputFormats = []string{""} 
     o.Voices = []Ttsvoiceentity{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ttsengineentity) MarshalJSON() ([]byte, error) {
    type Alias Ttsengineentity

    if TtsengineentityMarshalled {
        return []byte("{}"), nil
    }
    TtsengineentityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Languages []string `json:"languages"`
        
        OutputFormats []string `json:"outputFormats"`
        
        Voices []Ttsvoiceentity `json:"voices"`
        
        IsDefault bool `json:"isDefault"`
        
        IsSecure bool `json:"isSecure"`
        *Alias
    }{

        


        


        
        Languages: []string{""},
        


        
        OutputFormats: []string{""},
        


        
        Voices: []Ttsvoiceentity{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

