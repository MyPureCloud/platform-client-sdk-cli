package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdditionallanguagessynonymsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdditionallanguagessynonymsDud struct { 
    Synonyms []string `json:"synonyms"`

}

// Additionallanguagessynonyms
type Additionallanguagessynonyms struct { 
    

}

// String returns a JSON representation of the model
func (o *Additionallanguagessynonyms) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Additionallanguagessynonyms) MarshalJSON() ([]byte, error) {
    type Alias Additionallanguagessynonyms

    if AdditionallanguagessynonymsMarshalled {
        return []byte("{}"), nil
    }
    AdditionallanguagessynonymsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

