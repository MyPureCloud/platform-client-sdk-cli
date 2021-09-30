package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GeolocationsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GeolocationsettingsDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Geolocationsettings
type Geolocationsettings struct { 
    


    // Name
    Name string `json:"name"`


    // Enabled
    Enabled bool `json:"enabled"`


    // MapboxKey
    MapboxKey string `json:"mapboxKey"`


    

}

// String returns a JSON representation of the model
func (o *Geolocationsettings) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Geolocationsettings) MarshalJSON() ([]byte, error) {
    type Alias Geolocationsettings

    if GeolocationsettingsMarshalled {
        return []byte("{}"), nil
    }
    GeolocationsettingsMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Enabled bool `json:"enabled"`
        
        MapboxKey string `json:"mapboxKey"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

