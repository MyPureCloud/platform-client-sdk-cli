package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesdatatypepropertyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesdatatypepropertyDud struct { 
    


    


    


    


    


    


    


    


    

}

// Capabilitiesdatatypeproperty
type Capabilitiesdatatypeproperty struct { 
    // Availability
    Availability Capabilitiesavailability `json:"availability"`


    // DataTypeId
    DataTypeId string `json:"dataTypeId"`


    // Description
    Description string `json:"description"`


    // EnumGroup
    EnumGroup string `json:"enumGroup"`


    // Facets
    Facets Capabilitiesdatatypefacetsmodel `json:"facets"`


    // Id
    Id string `json:"id"`


    // IsCollection
    IsCollection bool `json:"isCollection"`


    // IsComputed
    IsComputed bool `json:"isComputed"`


    // Name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypeproperty) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesdatatypeproperty) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesdatatypeproperty

    if CapabilitiesdatatypepropertyMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesdatatypepropertyMarshalled = true

    return json.Marshal(&struct {
        
        Availability Capabilitiesavailability `json:"availability"`
        
        DataTypeId string `json:"dataTypeId"`
        
        Description string `json:"description"`
        
        EnumGroup string `json:"enumGroup"`
        
        Facets Capabilitiesdatatypefacetsmodel `json:"facets"`
        
        Id string `json:"id"`
        
        IsCollection bool `json:"isCollection"`
        
        IsComputed bool `json:"isComputed"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

