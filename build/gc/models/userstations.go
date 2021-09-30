package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserstationsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserstationsDud struct { 
    AssociatedStation Userstation `json:"associatedStation"`


    EffectiveStation Userstation `json:"effectiveStation"`


    DefaultStation Userstation `json:"defaultStation"`


    LastAssociatedStation Userstation `json:"lastAssociatedStation"`

}

// Userstations
type Userstations struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Userstations) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userstations) MarshalJSON() ([]byte, error) {
    type Alias Userstations

    if UserstationsMarshalled {
        return []byte("{}"), nil
    }
    UserstationsMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

