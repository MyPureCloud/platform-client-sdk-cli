package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicphraseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicphraseDud struct { 
    


    


    

}

// Topicphrase
type Topicphrase struct { 
    // Id
    Id string `json:"id"`


    // Text
    Text string `json:"text"`


    // UtteranceCount
    UtteranceCount int `json:"utteranceCount"`

}

// String returns a JSON representation of the model
func (o *Topicphrase) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topicphrase) MarshalJSON() ([]byte, error) {
    type Alias Topicphrase

    if TopicphraseMarshalled {
        return []byte("{}"), nil
    }
    TopicphraseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        UtteranceCount int `json:"utteranceCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

