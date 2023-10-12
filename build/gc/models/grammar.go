package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GrammarMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GrammarDud struct { 
    Id string `json:"id"`


    


    


    Languages []Grammarlanguage `json:"languages"`


    SelfUri string `json:"selfUri"`

}

// Grammar
type Grammar struct { 
    


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    


    

}

// String returns a JSON representation of the model
func (o *Grammar) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Grammar) MarshalJSON() ([]byte, error) {
    type Alias Grammar

    if GrammarMarshalled {
        return []byte("{}"), nil
    }
    GrammarMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

