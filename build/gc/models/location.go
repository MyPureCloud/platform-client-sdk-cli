package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocationDud struct { 
    


    FloorplanId string `json:"floorplanId"`


    


    


    

}

// Location
type Location struct { 
    // Id - Unique identifier for the location
    Id string `json:"id"`


    


    // Coordinates - Users coordinates on the floorplan. Only used when floorplanImage is set
    Coordinates map[string]float64 `json:"coordinates"`


    // Notes - Optional description on the users location
    Notes string `json:"notes"`


    // LocationDefinition
    LocationDefinition Locationdefinition `json:"locationDefinition"`

}

// String returns a JSON representation of the model
func (o *Location) String() string {
    
     o.Coordinates = map[string]float64{"": 0.0} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Location) MarshalJSON() ([]byte, error) {
    type Alias Location

    if LocationMarshalled {
        return []byte("{}"), nil
    }
    LocationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Coordinates map[string]float64 `json:"coordinates"`
        
        Notes string `json:"notes"`
        
        LocationDefinition Locationdefinition `json:"locationDefinition"`
        *Alias
    }{

        


        


        
        Coordinates: map[string]float64{"": 0.0},
        


        


        

        Alias: (*Alias)(u),
    })
}

