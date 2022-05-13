package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImporterrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImporterrorDud struct { 
    


    

}

// Importerror
type Importerror struct { 
    // Message
    Message string `json:"message"`


    // Line
    Line int `json:"line"`

}

// String returns a JSON representation of the model
func (o *Importerror) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Importerror) MarshalJSON() ([]byte, error) {
    type Alias Importerror

    if ImporterrorMarshalled {
        return []byte("{}"), nil
    }
    ImporterrorMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        Line int `json:"line"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

