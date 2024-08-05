package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnswergenerationdocumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnswergenerationdocumentDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Answergenerationdocument
type Answergenerationdocument struct { 
    


    // Title - The document title.
    Title string `json:"title"`


    

}

// String returns a JSON representation of the model
func (o *Answergenerationdocument) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Answergenerationdocument) MarshalJSON() ([]byte, error) {
    type Alias Answergenerationdocument

    if AnswergenerationdocumentMarshalled {
        return []byte("{}"), nil
    }
    AnswergenerationdocumentMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

