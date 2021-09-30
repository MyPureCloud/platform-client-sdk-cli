package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitlistitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitlistitemDud struct { 
    Id string `json:"id"`


    


    Authorized bool `json:"authorized"`


    


    SelfUri string `json:"selfUri"`

}

// Businessunitlistitem
type Businessunitlistitem struct { 
    


    // Name
    Name string `json:"name"`


    


    // Division - The division to which this entity belongs.
    Division Divisionreference `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Businessunitlistitem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunitlistitem) MarshalJSON() ([]byte, error) {
    type Alias Businessunitlistitem

    if BusinessunitlistitemMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitlistitemMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        Division Divisionreference `json:"division"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

