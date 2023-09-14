package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConfusionintentinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConfusionintentinfoDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    

}

// Confusionintentinfo
type Confusionintentinfo struct { 
    


    


    // UtteranceCount - Number of utterances in this intent which are similar to parent utterance.
    UtteranceCount int `json:"utteranceCount"`

}

// String returns a JSON representation of the model
func (o *Confusionintentinfo) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Confusionintentinfo) MarshalJSON() ([]byte, error) {
    type Alias Confusionintentinfo

    if ConfusionintentinfoMarshalled {
        return []byte("{}"), nil
    }
    ConfusionintentinfoMarshalled = true

    return json.Marshal(&struct {
        
        UtteranceCount int `json:"utteranceCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

