package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentlocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentlocationDud struct { 
    


    


    


    


    

}

// Contentlocation - Location object.
type Contentlocation struct { 
    // Url - URL of the Location.
    Url string `json:"url"`


    // Address - Location postal address.
    Address string `json:"address"`


    // Text - Location name.
    Text string `json:"text"`


    // Latitude - Latitude of the location.
    Latitude float64 `json:"latitude"`


    // Longitude - Longitude of the location.
    Longitude float64 `json:"longitude"`

}

// String returns a JSON representation of the model
func (o *Contentlocation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentlocation) MarshalJSON() ([]byte, error) {
    type Alias Contentlocation

    if ContentlocationMarshalled {
        return []byte("{}"), nil
    }
    ContentlocationMarshalled = true

    return json.Marshal(&struct { 
        Url string `json:"url"`
        
        Address string `json:"address"`
        
        Text string `json:"text"`
        
        Latitude float64 `json:"latitude"`
        
        Longitude float64 `json:"longitude"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

