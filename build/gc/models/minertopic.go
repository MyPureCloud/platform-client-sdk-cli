package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinertopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinertopicDud struct { 
    Id string `json:"id"`


    


    Miner Miner `json:"miner"`


    ConversationCount int `json:"conversationCount"`


    ConversationPercent float32 `json:"conversationPercent"`


    UtteranceCount int `json:"utteranceCount"`


    PhraseCount int `json:"phraseCount"`


    


    SelfUri string `json:"selfUri"`

}

// Minertopic
type Minertopic struct { 
    


    // Name - Topic name.
    Name string `json:"name"`


    


    


    


    


    


    // Phrases - Phrases associated with a topic.
    Phrases []Topicphrase `json:"phrases"`


    

}

// String returns a JSON representation of the model
func (o *Minertopic) String() string {
    
     o.Phrases = []Topicphrase{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minertopic) MarshalJSON() ([]byte, error) {
    type Alias Minertopic

    if MinertopicMarshalled {
        return []byte("{}"), nil
    }
    MinertopicMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Phrases []Topicphrase `json:"phrases"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Phrases: []Topicphrase{{}},
        


        

        Alias: (*Alias)(u),
    })
}

