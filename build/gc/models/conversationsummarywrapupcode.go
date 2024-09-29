package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationsummarywrapupcodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationsummarywrapupcodeDud struct { 
    


    


    SelfUri string `json:"selfUri"`


    Id string `json:"id"`


    Confidence float32 `json:"confidence"`

}

// Conversationsummarywrapupcode
type Conversationsummarywrapupcode struct { 
    // Name - The name of the wrapup code.
    Name string `json:"name"`


    // Description - The description of the wrapup code.
    Description string `json:"description"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Conversationsummarywrapupcode) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationsummarywrapupcode) MarshalJSON() ([]byte, error) {
    type Alias Conversationsummarywrapupcode

    if ConversationsummarywrapupcodeMarshalled {
        return []byte("{}"), nil
    }
    ConversationsummarywrapupcodeMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

