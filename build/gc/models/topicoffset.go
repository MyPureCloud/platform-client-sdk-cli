package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicoffsetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicoffsetDud struct { 
    WordCount int `json:"wordCount"`


    CharacterCount int `json:"characterCount"`

}

// Topicoffset
type Topicoffset struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Topicoffset) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topicoffset) MarshalJSON() ([]byte, error) {
    type Alias Topicoffset

    if TopicoffsetMarshalled {
        return []byte("{}"), nil
    }
    TopicoffsetMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

