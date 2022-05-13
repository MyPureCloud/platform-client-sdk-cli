package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediatranscriptionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediatranscriptionDud struct { 
    


    


    

}

// Mediatranscription
type Mediatranscription struct { 
    // DisplayName
    DisplayName string `json:"displayName"`


    // TranscriptionProvider
    TranscriptionProvider string `json:"transcriptionProvider"`


    // IntegrationId
    IntegrationId string `json:"integrationId"`

}

// String returns a JSON representation of the model
func (o *Mediatranscription) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediatranscription) MarshalJSON() ([]byte, error) {
    type Alias Mediatranscription

    if MediatranscriptionMarshalled {
        return []byte("{}"), nil
    }
    MediatranscriptionMarshalled = true

    return json.Marshal(&struct {
        
        DisplayName string `json:"displayName"`
        
        TranscriptionProvider string `json:"transcriptionProvider"`
        
        IntegrationId string `json:"integrationId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

