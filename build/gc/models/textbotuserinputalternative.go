package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotuserinputalternativeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotuserinputalternativeDud struct { 
    

}

// Textbotuserinputalternative - User input data used in a bot flow turn.
type Textbotuserinputalternative struct { 
    // Transcript - The user input transcript.
    Transcript Textbottranscript `json:"transcript"`

}

// String returns a JSON representation of the model
func (o *Textbotuserinputalternative) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotuserinputalternative) MarshalJSON() ([]byte, error) {
    type Alias Textbotuserinputalternative

    if TextbotuserinputalternativeMarshalled {
        return []byte("{}"), nil
    }
    TextbotuserinputalternativeMarshalled = true

    return json.Marshal(&struct { 
        Transcript Textbottranscript `json:"transcript"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

