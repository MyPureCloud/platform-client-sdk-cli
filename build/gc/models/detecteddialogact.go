package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DetecteddialogactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DetecteddialogactDud struct { 
    Name string `json:"name"`


    Probability float64 `json:"probability"`

}

// Detecteddialogact
type Detecteddialogact struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Detecteddialogact) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Detecteddialogact) MarshalJSON() ([]byte, error) {
    type Alias Detecteddialogact

    if DetecteddialogactMarshalled {
        return []byte("{}"), nil
    }
    DetecteddialogactMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

