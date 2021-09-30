package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GeolocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GeolocationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Geolocation
type Geolocation struct { 
    


    // Name
    Name string `json:"name"`


    // VarType - A string used to describe the type of client the geolocation is being updated from e.g. ios, android, web, etc.
    VarType string `json:"type"`


    // Primary - A boolean used to tell whether or not to set this geolocation client as the primary on a PATCH
    Primary bool `json:"primary"`


    // Latitude
    Latitude float64 `json:"latitude"`


    // Longitude
    Longitude float64 `json:"longitude"`


    // Country
    Country string `json:"country"`


    // Region
    Region string `json:"region"`


    // City
    City string `json:"city"`


    // Locations
    Locations []Locationdefinition `json:"locations"`


    

}

// String returns a JSON representation of the model
func (o *Geolocation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Locations = []Locationdefinition{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Geolocation) MarshalJSON() ([]byte, error) {
    type Alias Geolocation

    if GeolocationMarshalled {
        return []byte("{}"), nil
    }
    GeolocationMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Primary bool `json:"primary"`
        
        Latitude float64 `json:"latitude"`
        
        Longitude float64 `json:"longitude"`
        
        Country string `json:"country"`
        
        Region string `json:"region"`
        
        City string `json:"city"`
        
        Locations []Locationdefinition `json:"locations"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Locations: []Locationdefinition{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

