package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LicensedefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LicensedefinitionDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Licensedefinition
type Licensedefinition struct { 
    


    // Description
    Description string `json:"description"`


    // Permissions
    Permissions Permissions `json:"permissions"`


    // Prerequisites
    Prerequisites []Addressablelicensedefinition `json:"prerequisites"`


    // Comprises
    Comprises []Licensedefinition `json:"comprises"`


    

}

// String returns a JSON representation of the model
func (o *Licensedefinition) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.Prerequisites = []Addressablelicensedefinition{{}} 
    
    
    
     o.Comprises = []Licensedefinition{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Licensedefinition) MarshalJSON() ([]byte, error) {
    type Alias Licensedefinition

    if LicensedefinitionMarshalled {
        return []byte("{}"), nil
    }
    LicensedefinitionMarshalled = true

    return json.Marshal(&struct { 
        
        
        Description string `json:"description"`
        
        Permissions Permissions `json:"permissions"`
        
        Prerequisites []Addressablelicensedefinition `json:"prerequisites"`
        
        Comprises []Licensedefinition `json:"comprises"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Prerequisites: []Addressablelicensedefinition{{}},
        

        

        
        Comprises: []Licensedefinition{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

