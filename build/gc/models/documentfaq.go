package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentfaqMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentfaqDud struct { 
    


    


    

}

// Documentfaq
type Documentfaq struct { 
    // Question - The question for this FAQ
    Question string `json:"question"`


    // Answer - The answer for this FAQ
    Answer string `json:"answer"`


    // Alternatives - List of Alternative questions related to the answer which helps in improving the likelihood of a match to user query
    Alternatives []string `json:"alternatives"`

}

// String returns a JSON representation of the model
func (o *Documentfaq) String() string {
    
    
     o.Alternatives = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentfaq) MarshalJSON() ([]byte, error) {
    type Alias Documentfaq

    if DocumentfaqMarshalled {
        return []byte("{}"), nil
    }
    DocumentfaqMarshalled = true

    return json.Marshal(&struct {
        
        Question string `json:"question"`
        
        Answer string `json:"answer"`
        
        Alternatives []string `json:"alternatives"`
        *Alias
    }{

        


        


        
        Alternatives: []string{""},
        

        Alias: (*Alias)(u),
    })
}

