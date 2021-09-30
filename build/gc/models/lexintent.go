package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LexintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LexintentDud struct { 
    


    


    


    

}

// Lexintent
type Lexintent struct { 
    // Name - The intent name
    Name string `json:"name"`


    // Description - A description of the intent
    Description string `json:"description"`


    // Slots - An object mapping slot names to Slot objects
    Slots map[string]Lexslot `json:"slots"`


    // Version - The intent version
    Version string `json:"version"`

}

// String returns a JSON representation of the model
func (o *Lexintent) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Slots = map[string]Lexslot{"": {}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexintent) MarshalJSON() ([]byte, error) {
    type Alias Lexintent

    if LexintentMarshalled {
        return []byte("{}"), nil
    }
    LexintentMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Slots map[string]Lexslot `json:"slots"`
        
        Version string `json:"version"`
        
        *Alias
    }{
        

        

        

        

        

        
        Slots: map[string]Lexslot{"": {}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

