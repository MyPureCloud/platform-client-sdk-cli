package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SourceDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Source
type Source struct { 
    


    // Name - The name of the source
    Name string `json:"name"`


    // Description - The description of the source
    Description string `json:"description"`


    // VarType - The type of source
    VarType string `json:"type"`


    // Deactivated
    Deactivated bool `json:"deactivated"`


    

}

// String returns a JSON representation of the model
func (o *Source) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Source) MarshalJSON() ([]byte, error) {
    type Alias Source

    if SourceMarshalled {
        return []byte("{}"), nil
    }
    SourceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        VarType string `json:"type"`
        
        Deactivated bool `json:"deactivated"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

