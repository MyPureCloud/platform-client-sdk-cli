package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesdatatypefacetsmodelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesdatatypefacetsmodelDud struct { 
    


    


    


    


    


    


    


    

}

// Capabilitiesdatatypefacetsmodel
type Capabilitiesdatatypefacetsmodel struct { 
    // MaxItems
    MaxItems int `json:"maxItems"`


    // MinItems
    MinItems int `json:"minItems"`


    // Minimum
    Minimum Capabilitiesdatatypefacetslimit `json:"minimum"`


    // MinimumExclusive
    MinimumExclusive bool `json:"minimumExclusive"`


    // Maximum
    Maximum Capabilitiesdatatypefacetslimit `json:"maximum"`


    // MaximumExclusive
    MaximumExclusive bool `json:"maximumExclusive"`


    // MinLength
    MinLength int `json:"minLength"`


    // MaxLength
    MaxLength int `json:"maxLength"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypefacetsmodel) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesdatatypefacetsmodel) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesdatatypefacetsmodel

    if CapabilitiesdatatypefacetsmodelMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesdatatypefacetsmodelMarshalled = true

    return json.Marshal(&struct {
        
        MaxItems int `json:"maxItems"`
        
        MinItems int `json:"minItems"`
        
        Minimum Capabilitiesdatatypefacetslimit `json:"minimum"`
        
        MinimumExclusive bool `json:"minimumExclusive"`
        
        Maximum Capabilitiesdatatypefacetslimit `json:"maximum"`
        
        MaximumExclusive bool `json:"maximumExclusive"`
        
        MinLength int `json:"minLength"`
        
        MaxLength int `json:"maxLength"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

