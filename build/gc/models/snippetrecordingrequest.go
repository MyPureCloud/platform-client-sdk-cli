package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SnippetrecordingrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SnippetrecordingrequestDud struct { 
    


    

}

// Snippetrecordingrequest
type Snippetrecordingrequest struct { 
    // VarType - Type of recording to apply to the participant.
    VarType string `json:"type"`


    // Record
    Record bool `json:"record"`

}

// String returns a JSON representation of the model
func (o *Snippetrecordingrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Snippetrecordingrequest) MarshalJSON() ([]byte, error) {
    type Alias Snippetrecordingrequest

    if SnippetrecordingrequestMarshalled {
        return []byte("{}"), nil
    }
    SnippetrecordingrequestMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Record bool `json:"record"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

