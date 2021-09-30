package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OverallbestpointsitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OverallbestpointsitemDud struct { 
    GranularityType string `json:"granularityType"`


    Users []Userreference `json:"users"`


    Count int `json:"count"`


    Points int `json:"points"`


    DateStartWorkday time.Time `json:"dateStartWorkday"`


    DateEndWorkday time.Time `json:"dateEndWorkday"`

}

// Overallbestpointsitem
type Overallbestpointsitem struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Overallbestpointsitem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Overallbestpointsitem) MarshalJSON() ([]byte, error) {
    type Alias Overallbestpointsitem

    if OverallbestpointsitemMarshalled {
        return []byte("{}"), nil
    }
    OverallbestpointsitemMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

