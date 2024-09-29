package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscripttranslationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscripttranslationDud struct { 
    


    


    

}

// Transcripttranslation
type Transcripttranslation struct { 
    // Id - Transcript Id
    Id string `json:"id"`


    // Phrases - List of translated phrases, if translation succeeded
    Phrases []Phrasetranslation `json:"phrases"`


    // TranslateError - Translation error, if translation failed
    TranslateError string `json:"translateError"`

}

// String returns a JSON representation of the model
func (o *Transcripttranslation) String() string {
    
     o.Phrases = []Phrasetranslation{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcripttranslation) MarshalJSON() ([]byte, error) {
    type Alias Transcripttranslation

    if TranscripttranslationMarshalled {
        return []byte("{}"), nil
    }
    TranscripttranslationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Phrases []Phrasetranslation `json:"phrases"`
        
        TranslateError string `json:"translateError"`
        *Alias
    }{

        


        
        Phrases: []Phrasetranslation{{}},
        


        

        Alias: (*Alias)(u),
    })
}

