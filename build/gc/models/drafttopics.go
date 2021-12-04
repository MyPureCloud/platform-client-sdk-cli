package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DrafttopicsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DrafttopicsDud struct { 
    


    


    Miner Miner `json:"miner"`


    ConversationCount int `json:"conversationCount"`


    ConversationPercent float32 `json:"conversationPercent"`


    UtteranceCount int `json:"utteranceCount"`


    PhraseCount int `json:"phraseCount"`


    


    SelfUri string `json:"selfUri"`

}

// Drafttopics
type Drafttopics struct { 
    // Id - Id for a topic.
    Id string `json:"id"`


    // Name - Topic name.
    Name string `json:"name"`


    


    


    


    


    


    // Phrases - The phrases that are extracted for a topic.
    Phrases []string `json:"phrases"`


    

}

// String returns a JSON representation of the model
func (o *Drafttopics) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Phrases = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Drafttopics) MarshalJSON() ([]byte, error) {
    type Alias Drafttopics

    if DrafttopicsMarshalled {
        return []byte("{}"), nil
    }
    DrafttopicsMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        
        
        
        
        
        
        
        
        
        
        Phrases []string `json:"phrases"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Phrases: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

