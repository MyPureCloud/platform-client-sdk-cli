package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoryrequestDud struct { 
    


    


    


    

}

// Categoryrequest
type Categoryrequest struct { 
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
func (o *Categoryrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categoryrequest) MarshalJSON() ([]byte, error) {
    type Alias Categoryrequest

    if CategoryrequestMarshalled {
        return []byte("{}"), nil
    }
    CategoryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        InteractionType string `json:"interactionType"`
        
        Criteria Operand `json:"criteria"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

