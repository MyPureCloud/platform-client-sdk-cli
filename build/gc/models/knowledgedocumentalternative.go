package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentalternativeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentalternativeDud struct { 
    


    

}

// Knowledgedocumentalternative
type Knowledgedocumentalternative struct { 
    // Phrase - Alternate phrasing to the document title, having a limit of 500 words.
    Phrase string `json:"phrase"`


    // Autocomplete - Autocomplete enabled for the alternate phrase.
    Autocomplete bool `json:"autocomplete"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentalternative) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentalternative) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentalternative

    if KnowledgedocumentalternativeMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentalternativeMarshalled = true

    return json.Marshal(&struct {
        
        Phrase string `json:"phrase"`
        
        Autocomplete bool `json:"autocomplete"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

