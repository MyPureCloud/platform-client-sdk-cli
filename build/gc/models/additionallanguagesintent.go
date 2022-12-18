package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdditionallanguagesintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdditionallanguagesintentDud struct { 
    Id string `json:"id"`


    

}

// Additionallanguagesintent
type Additionallanguagesintent struct { 
    


    // Utterances - Utterances list for additional language
    Utterances []Nluutterance `json:"utterances"`

}

// String returns a JSON representation of the model
func (o *Additionallanguagesintent) String() string {
     o.Utterances = []Nluutterance{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Additionallanguagesintent) MarshalJSON() ([]byte, error) {
    type Alias Additionallanguagesintent

    if AdditionallanguagesintentMarshalled {
        return []byte("{}"), nil
    }
    AdditionallanguagesintentMarshalled = true

    return json.Marshal(&struct {
        
        Utterances []Nluutterance `json:"utterances"`
        *Alias
    }{

        


        
        Utterances: []Nluutterance{{}},
        

        Alias: (*Alias)(u),
    })
}

