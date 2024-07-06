package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestionlistingDud struct { 
    


    


    


    

}

// Suggestionlisting
type Suggestionlisting struct { 
    // Entities
    Entities []Suggestion `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Suggestionlisting) String() string {
     o.Entities = []Suggestion{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestionlisting) MarshalJSON() ([]byte, error) {
    type Alias Suggestionlisting

    if SuggestionlistingMarshalled {
        return []byte("{}"), nil
    }
    SuggestionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Suggestion `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Suggestion{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

