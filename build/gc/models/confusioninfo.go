package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConfusioninfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConfusioninfoDud struct { 
    

}

// Confusioninfo
type Confusioninfo struct { 
    // Intents - Confusion details between this utterance and other intents.
    Intents []Confusionintentinfo `json:"intents"`

}

// String returns a JSON representation of the model
func (o *Confusioninfo) String() string {
     o.Intents = []Confusionintentinfo{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Confusioninfo) MarshalJSON() ([]byte, error) {
    type Alias Confusioninfo

    if ConfusioninfoMarshalled {
        return []byte("{}"), nil
    }
    ConfusioninfoMarshalled = true

    return json.Marshal(&struct {
        
        Intents []Confusionintentinfo `json:"intents"`
        *Alias
    }{

        
        Intents: []Confusionintentinfo{{}},
        

        Alias: (*Alias)(u),
    })
}

