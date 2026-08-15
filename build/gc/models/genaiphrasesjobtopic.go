package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GenaiphrasesjobtopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GenaiphrasesjobtopicDud struct { 
    


    


    


    

}

// Genaiphrasesjobtopic
type Genaiphrasesjobtopic struct { 
    // Name - The topic name
    Name string `json:"name"`


    // Description - The topic description
    Description string `json:"description"`


    // Dialect - The topic dialect
    Dialect string `json:"dialect"`


    // Phrases - Existing phrases for the topic
    Phrases []string `json:"phrases"`

}

// String returns a JSON representation of the model
func (o *Genaiphrasesjobtopic) String() string {
    
    
    
     o.Phrases = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Genaiphrasesjobtopic) MarshalJSON() ([]byte, error) {
    type Alias Genaiphrasesjobtopic

    if GenaiphrasesjobtopicMarshalled {
        return []byte("{}"), nil
    }
    GenaiphrasesjobtopicMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Dialect string `json:"dialect"`
        
        Phrases []string `json:"phrases"`
        *Alias
    }{

        


        


        


        
        Phrases: []string{""},
        

        Alias: (*Alias)(u),
    })
}

