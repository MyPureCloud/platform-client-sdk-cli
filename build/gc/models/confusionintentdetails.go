package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConfusionintentdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConfusionintentdetailsDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Confusionintentdetails
type Confusionintentdetails struct { 
    


    // Name
    Name string `json:"name"`


    // UtteranceCount - Number of utterances in this intent which are similar to parent utterance.
    UtteranceCount int `json:"utteranceCount"`


    // Utterances - List of utterance which are similar to parent utterance.
    Utterances []Confusionutterance `json:"utterances"`


    

}

// String returns a JSON representation of the model
func (o *Confusionintentdetails) String() string {
    
    
     o.Utterances = []Confusionutterance{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Confusionintentdetails) MarshalJSON() ([]byte, error) {
    type Alias Confusionintentdetails

    if ConfusionintentdetailsMarshalled {
        return []byte("{}"), nil
    }
    ConfusionintentdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        UtteranceCount int `json:"utteranceCount"`
        
        Utterances []Confusionutterance `json:"utterances"`
        *Alias
    }{

        


        


        


        
        Utterances: []Confusionutterance{{}},
        


        

        Alias: (*Alias)(u),
    })
}

