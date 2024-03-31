package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentcontenthighlightindexMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentcontenthighlightindexDud struct { 
    


    

}

// Documentcontenthighlightindex
type Documentcontenthighlightindex struct { 
    // StartIndex - Highlight text start index.
    StartIndex int `json:"startIndex"`


    // EndIndex - Highlight text end index.
    EndIndex int `json:"endIndex"`

}

// String returns a JSON representation of the model
func (o *Documentcontenthighlightindex) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentcontenthighlightindex) MarshalJSON() ([]byte, error) {
    type Alias Documentcontenthighlightindex

    if DocumentcontenthighlightindexMarshalled {
        return []byte("{}"), nil
    }
    DocumentcontenthighlightindexMarshalled = true

    return json.Marshal(&struct {
        
        StartIndex int `json:"startIndex"`
        
        EndIndex int `json:"endIndex"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

