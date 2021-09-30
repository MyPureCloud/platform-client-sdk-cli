package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneygeolocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneygeolocationDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Journeygeolocation
type Journeygeolocation struct { 
    // Country - Geolocation's ISO 3166-1 alpha-2 country code.
    Country string `json:"country"`


    // CountryName - Geolocation's country name.
    CountryName string `json:"countryName"`


    // Latitude - Geolocation's latitude.
    Latitude float64 `json:"latitude"`


    // Longitude - Geolocation's longitude.
    Longitude float64 `json:"longitude"`


    // Locality - Geolocation's locality or city.
    Locality string `json:"locality"`


    // PostalCode - Geolocation's postal code or ZIP code.
    PostalCode string `json:"postalCode"`


    // Region - Geolocation's ISO-3166-2 region code.
    Region string `json:"region"`


    // RegionName - Geolocation's region name.
    RegionName string `json:"regionName"`


    // Source - The source that was used to determine the geolocation information.
    Source string `json:"source"`


    // Timezone - Geolocation's timezone.
    Timezone string `json:"timezone"`

}

// String returns a JSON representation of the model
func (o *Journeygeolocation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeygeolocation) MarshalJSON() ([]byte, error) {
    type Alias Journeygeolocation

    if JourneygeolocationMarshalled {
        return []byte("{}"), nil
    }
    JourneygeolocationMarshalled = true

    return json.Marshal(&struct { 
        Country string `json:"country"`
        
        CountryName string `json:"countryName"`
        
        Latitude float64 `json:"latitude"`
        
        Longitude float64 `json:"longitude"`
        
        Locality string `json:"locality"`
        
        PostalCode string `json:"postalCode"`
        
        Region string `json:"region"`
        
        RegionName string `json:"regionName"`
        
        Source string `json:"source"`
        
        Timezone string `json:"timezone"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

