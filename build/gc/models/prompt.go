package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PromptMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PromptDud struct { 
    


    


    


    Resources []Promptasset `json:"resources"`


    CurrentOperation Operation `json:"currentOperation"`


    SelfUri string `json:"selfUri"`

}

// Prompt
type Prompt struct { 
    // Id - The prompt identifier
    Id string `json:"id"`


    // Name - The prompt name.
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Prompt) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Prompt) MarshalJSON() ([]byte, error) {
    type Alias Prompt

    if PromptMarshalled {
        return []byte("{}"), nil
    }
    PromptMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

