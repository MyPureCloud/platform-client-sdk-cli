package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TesttopicphrasetopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TesttopicphrasetopicDud struct { 
    


    


    


    

}

// Testtopicphrasetopic
type Testtopicphrasetopic struct { 
    // Phrase - The topic phrase to test
    Phrase Testtopicphrasephrase `json:"phrase"`


    // Strictness - The topic strictness, default value is 72
    Strictness string `json:"strictness"`


    // Dialect - The topic dialect, default value is en-US
    Dialect string `json:"dialect"`


    // Participants - The topic participants, default value is both
    Participants string `json:"participants"`

}

// String returns a JSON representation of the model
func (o *Testtopicphrasetopic) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testtopicphrasetopic) MarshalJSON() ([]byte, error) {
    type Alias Testtopicphrasetopic

    if TesttopicphrasetopicMarshalled {
        return []byte("{}"), nil
    }
    TesttopicphrasetopicMarshalled = true

    return json.Marshal(&struct {
        
        Phrase Testtopicphrasephrase `json:"phrase"`
        
        Strictness string `json:"strictness"`
        
        Dialect string `json:"dialect"`
        
        Participants string `json:"participants"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

