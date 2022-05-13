package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DraftintentsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DraftintentsDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Draftintents
type Draftintents struct { 
    // Id - Id for an intent.
    Id string `json:"id"`


    // Name - Name/Label for an intent.
    Name string `json:"name"`


    // Utterances - The utterances that are extracted for an Intent.
    Utterances []string `json:"utterances"`


    

}

// String returns a JSON representation of the model
func (o *Draftintents) String() string {
    
    
     o.Utterances = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Draftintents) MarshalJSON() ([]byte, error) {
    type Alias Draftintents

    if DraftintentsMarshalled {
        return []byte("{}"), nil
    }
    DraftintentsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Utterances []string `json:"utterances"`
        *Alias
    }{

        


        


        
        Utterances: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

