package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportedcontentreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportedcontentreferenceDud struct { 
    


    Name string `json:"name"`


    SelfUri string `json:"selfUri"`


    MediaTypes Mediatypes `json:"mediaTypes"`

}

// Supportedcontentreference - Reference to supported content profile associated with the integration
type Supportedcontentreference struct { 
    // Id - The SupportedContent unique identifier associated with this integration
    Id string `json:"id"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Supportedcontentreference) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportedcontentreference) MarshalJSON() ([]byte, error) {
    type Alias Supportedcontentreference

    if SupportedcontentreferenceMarshalled {
        return []byte("{}"), nil
    }
    SupportedcontentreferenceMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

