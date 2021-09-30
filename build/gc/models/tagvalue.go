package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TagvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TagvalueDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Tagvalue
type Tagvalue struct { 
    


    // Name - The workspace tag name.
    Name string `json:"name"`


    // InUse
    InUse bool `json:"inUse"`


    // Acl
    Acl []string `json:"acl"`


    

}

// String returns a JSON representation of the model
func (o *Tagvalue) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.Acl = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Tagvalue) MarshalJSON() ([]byte, error) {
    type Alias Tagvalue

    if TagvalueMarshalled {
        return []byte("{}"), nil
    }
    TagvalueMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        InUse bool `json:"inUse"`
        
        Acl []string `json:"acl"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Acl: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

