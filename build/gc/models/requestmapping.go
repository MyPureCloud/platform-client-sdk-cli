package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequestmappingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequestmappingDud struct { 
    


    


    


    

}

// Requestmapping
type Requestmapping struct { 
    // Name - Name of the Integration Action Attribute to supply the value for
    Name string `json:"name"`


    // AttributeType - Type of the value supplied
    AttributeType string `json:"attributeType"`


    // MappingType - Method of finding value to use with Attribute
    MappingType string `json:"mappingType"`


    // Value - Value to supply for the specified Attribute
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Requestmapping) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requestmapping) MarshalJSON() ([]byte, error) {
    type Alias Requestmapping

    if RequestmappingMarshalled {
        return []byte("{}"), nil
    }
    RequestmappingMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        AttributeType string `json:"attributeType"`
        
        MappingType string `json:"mappingType"`
        
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

