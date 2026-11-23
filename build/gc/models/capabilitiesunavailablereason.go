package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesunavailablereasonMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesunavailablereasonDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Capabilitiesunavailablereason
type Capabilitiesunavailablereason struct { 
    // CustomCondition
    CustomCondition string `json:"customCondition"`


    // Enabled
    Enabled bool `json:"enabled"`


    // ExceptionEncountered
    ExceptionEncountered bool `json:"exceptionEncountered"`


    // Feature - The capability isn't available because a feature is not available
    Feature string `json:"feature"`


    // FeaturesAllOf - The capability isn't available because all of a given set of features are not available
    FeaturesAllOf []string `json:"featuresAllOf"`


    // FeaturesOneOf - The capability isn't available because a feature isn't available
    FeaturesOneOf []string `json:"featuresOneOf"`


    // ProductsAllOf - The capability isn't available because all of a given set of products are not available
    ProductsAllOf []string `json:"productsAllOf"`


    // ProductsOneOf - The capability isn't available because a product is not available
    ProductsOneOf []string `json:"productsOneOf"`


    // UnavailableFeature
    UnavailableFeature string `json:"unavailableFeature"`


    // UnavailableProduct
    UnavailableProduct string `json:"unavailableProduct"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesunavailablereason) String() string {
    
    
    
    
     o.FeaturesAllOf = []string{""} 
     o.FeaturesOneOf = []string{""} 
     o.ProductsAllOf = []string{""} 
     o.ProductsOneOf = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesunavailablereason) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesunavailablereason

    if CapabilitiesunavailablereasonMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesunavailablereasonMarshalled = true

    return json.Marshal(&struct {
        
        CustomCondition string `json:"customCondition"`
        
        Enabled bool `json:"enabled"`
        
        ExceptionEncountered bool `json:"exceptionEncountered"`
        
        Feature string `json:"feature"`
        
        FeaturesAllOf []string `json:"featuresAllOf"`
        
        FeaturesOneOf []string `json:"featuresOneOf"`
        
        ProductsAllOf []string `json:"productsAllOf"`
        
        ProductsOneOf []string `json:"productsOneOf"`
        
        UnavailableFeature string `json:"unavailableFeature"`
        
        UnavailableProduct string `json:"unavailableProduct"`
        *Alias
    }{

        


        


        


        


        
        FeaturesAllOf: []string{""},
        


        
        FeaturesOneOf: []string{""},
        


        
        ProductsAllOf: []string{""},
        


        
        ProductsOneOf: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

