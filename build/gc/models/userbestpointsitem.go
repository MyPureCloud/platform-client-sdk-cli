package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserbestpointsitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserbestpointsitemDud struct { 
    GranularityType string `json:"granularityType"`


    Points int `json:"points"`


    DateStartWorkday time.Time `json:"dateStartWorkday"`


    DateEndWorkday time.Time `json:"dateEndWorkday"`


    Rank int `json:"rank"`

}

// Userbestpointsitem
type Userbestpointsitem struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Userbestpointsitem) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userbestpointsitem) MarshalJSON() ([]byte, error) {
    type Alias Userbestpointsitem

    if UserbestpointsitemMarshalled {
        return []byte("{}"), nil
    }
    UserbestpointsitemMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

