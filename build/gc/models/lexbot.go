package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LexbotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LexbotDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Lexbot
type Lexbot struct { 
    


    // Name
    Name string `json:"name"`


    // Description - A description of the Lex bot
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Lexbot) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexbot) MarshalJSON() ([]byte, error) {
    type Alias Lexbot

    if LexbotMarshalled {
        return []byte("{}"), nil
    }
    LexbotMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

