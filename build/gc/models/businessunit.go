package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Businessunit
type Businessunit struct { 
    


    // Name
    Name string `json:"name"`


    // Settings - Settings for this business unit
    Settings Businessunitsettings `json:"settings"`


    // Division - The division to which this entity belongs.
    Division Divisionreference `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Businessunit) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunit) MarshalJSON() ([]byte, error) {
    type Alias Businessunit

    if BusinessunitMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Settings Businessunitsettings `json:"settings"`
        
        Division Divisionreference `json:"division"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

