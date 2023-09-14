package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConfusionutteranceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConfusionutteranceDud struct { 
    Id string `json:"id"`


    


    

}

// Confusionutterance
type Confusionutterance struct { 
    


    // Text - Utterance Text.
    Text string `json:"text"`


    // Similarity - Utterance's similarity score, 0 being dissimilar and 1 being very similar.
    Similarity float32 `json:"similarity"`

}

// String returns a JSON representation of the model
func (o *Confusionutterance) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Confusionutterance) MarshalJSON() ([]byte, error) {
    type Alias Confusionutterance

    if ConfusionutteranceMarshalled {
        return []byte("{}"), nil
    }
    ConfusionutteranceMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Similarity float32 `json:"similarity"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

