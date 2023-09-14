package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConfusiondetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConfusiondetailsDud struct { 
    

}

// Confusiondetails
type Confusiondetails struct { 
    // Intents - Confusion details between this utterance and other intents.
    Intents []Confusionintentdetails `json:"intents"`

}

// String returns a JSON representation of the model
func (o *Confusiondetails) String() string {
     o.Intents = []Confusionintentdetails{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Confusiondetails) MarshalJSON() ([]byte, error) {
    type Alias Confusiondetails

    if ConfusiondetailsMarshalled {
        return []byte("{}"), nil
    }
    ConfusiondetailsMarshalled = true

    return json.Marshal(&struct {
        
        Intents []Confusionintentdetails `json:"intents"`
        *Alias
    }{

        
        Intents: []Confusionintentdetails{{}},
        

        Alias: (*Alias)(u),
    })
}

