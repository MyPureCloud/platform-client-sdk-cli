package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommunicationtranslationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommunicationtranslationDud struct { 
    


    

}

// Communicationtranslation
type Communicationtranslation struct { 
    // CommunicationId - Communication Id
    CommunicationId string `json:"communicationId"`


    // Transcripts - List of translated transcripts
    Transcripts []Transcripttranslation `json:"transcripts"`

}

// String returns a JSON representation of the model
func (o *Communicationtranslation) String() string {
    
     o.Transcripts = []Transcripttranslation{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Communicationtranslation) MarshalJSON() ([]byte, error) {
    type Alias Communicationtranslation

    if CommunicationtranslationMarshalled {
        return []byte("{}"), nil
    }
    CommunicationtranslationMarshalled = true

    return json.Marshal(&struct {
        
        CommunicationId string `json:"communicationId"`
        
        Transcripts []Transcripttranslation `json:"transcripts"`
        *Alias
    }{

        


        
        Transcripts: []Transcripttranslation{{}},
        

        Alias: (*Alias)(u),
    })
}

