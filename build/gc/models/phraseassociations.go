package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhraseassociationsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhraseassociationsDud struct { 
    


    

}

// Phraseassociations
type Phraseassociations struct { 
    // PhraseId - Id of the phrase to be linked
    PhraseId string `json:"phraseId"`


    // DocumentId - Id of the document to be linked
    DocumentId string `json:"documentId"`

}

// String returns a JSON representation of the model
func (o *Phraseassociations) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phraseassociations) MarshalJSON() ([]byte, error) {
    type Alias Phraseassociations

    if PhraseassociationsMarshalled {
        return []byte("{}"), nil
    }
    PhraseassociationsMarshalled = true

    return json.Marshal(&struct {
        
        PhraseId string `json:"phraseId"`
        
        DocumentId string `json:"documentId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

