package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowprojectMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowprojectDud struct { 
    


    

}

// Dialogflowproject
type Dialogflowproject struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Dialogflowproject) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowproject) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowproject

    if DialogflowprojectMarshalled {
        return []byte("{}"), nil
    }
    DialogflowprojectMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

