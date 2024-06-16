package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TermMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TermDud struct { 
    


    

}

// Term
type Term struct { 
    // Word - Find term in interaction
    Word string `json:"word"`


    // ParticipantType - Dictates if term operand must come from the internal, external or both participants
    ParticipantType string `json:"participantType"`

}

// String returns a JSON representation of the model
func (o *Term) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Term) MarshalJSON() ([]byte, error) {
    type Alias Term

    if TermMarshalled {
        return []byte("{}"), nil
    }
    TermMarshalled = true

    return json.Marshal(&struct {
        
        Word string `json:"word"`
        
        ParticipantType string `json:"participantType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

