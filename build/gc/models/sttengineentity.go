package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SttengineentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SttengineentityDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Sttengineentity
type Sttengineentity struct { 
    


    // Name
    Name string `json:"name"`


    // GrammarBased - The STT engine is intended to be used for Grammars
    GrammarBased bool `json:"grammarBased"`


    

}

// String returns a JSON representation of the model
func (o *Sttengineentity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sttengineentity) MarshalJSON() ([]byte, error) {
    type Alias Sttengineentity

    if SttengineentityMarshalled {
        return []byte("{}"), nil
    }
    SttengineentityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        GrammarBased bool `json:"grammarBased"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

