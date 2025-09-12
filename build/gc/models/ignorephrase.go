package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IgnorephraseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IgnorephraseDud struct { 
    


    


    

}

// Ignorephrase
type Ignorephrase struct { 
    // Text - Text of the phrase to be ignored
    Text string `json:"text"`


    // Participant - Type of participant
    Participant string `json:"participant"`


    // MediaType - Media Type for the entity (Optional)
    MediaType string `json:"mediaType"`

}

// String returns a JSON representation of the model
func (o *Ignorephrase) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ignorephrase) MarshalJSON() ([]byte, error) {
    type Alias Ignorephrase

    if IgnorephraseMarshalled {
        return []byte("{}"), nil
    }
    IgnorephraseMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Participant string `json:"participant"`
        
        MediaType string `json:"mediaType"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

