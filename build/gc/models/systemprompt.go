package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SystempromptMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SystempromptDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Systemprompt
type Systemprompt struct { 
    // Id - The system prompt identifier
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // Resources
    Resources []Systempromptasset `json:"resources"`


    

}

// String returns a JSON representation of the model
func (o *Systemprompt) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Resources = []Systempromptasset{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Systemprompt) MarshalJSON() ([]byte, error) {
    type Alias Systemprompt

    if SystempromptMarshalled {
        return []byte("{}"), nil
    }
    SystempromptMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Resources []Systempromptasset `json:"resources"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Resources: []Systempromptasset{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

