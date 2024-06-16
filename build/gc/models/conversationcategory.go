package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcategoryDud struct { 
    


    


    


    


    

}

// Conversationcategory
type Conversationcategory struct { 
    // Id - The id of the category
    Id string `json:"id"`


    // Name - The name of the category
    Name string `json:"name"`


    // Description - The description of the category
    Description string `json:"description"`


    // InteractionType - The type of interaction the category will apply to
    InteractionType string `json:"interactionType"`


    // Criteria - A collection of conditions joined together by logical operation to provide more refined filtering of conversations
    Criteria Operand `json:"criteria"`

}

// String returns a JSON representation of the model
func (o *Conversationcategory) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcategory) MarshalJSON() ([]byte, error) {
    type Alias Conversationcategory

    if ConversationcategoryMarshalled {
        return []byte("{}"), nil
    }
    ConversationcategoryMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        InteractionType string `json:"interactionType"`
        
        Criteria Operand `json:"criteria"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

