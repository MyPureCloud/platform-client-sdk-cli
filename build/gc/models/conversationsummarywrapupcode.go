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

}

// Conversationsummarywrapupcode
type Conversationsummarywrapupcode struct { 
    // Id - The id of the wrapup code.
    Id string `json:"id"`


    // Name - The name of the wrapup code.
    Name string `json:"name"`


    // Description - The description of the wrapup code.
    Description string `json:"description"`


    // Confidence - The AI confidence value.
    Confidence float32 `json:"confidence"`


    

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
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Confidence float32 `json:"confidence"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

