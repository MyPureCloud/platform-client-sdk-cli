package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LimitcountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LimitcountDud struct { 
    


    


    


    


    

}

// Limitcount
type Limitcount struct { 
    // Name - The name of the limit.
    Name string `json:"name"`


    // EstimatedCount - The total used count of the limit.
    EstimatedCount int `json:"estimatedCount"`


    // Max - The maximum value of the limit.
    Max int `json:"max"`


    // EntityId - The entity which makes this count unique. The context of what the entity is would be dependant on the limit. May not be applicable for all limits.
    EntityId string `json:"entityId"`


    // UserId - The user which makes this count unique. May not be applicable for all limits.
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Limitcount) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Limitcount) MarshalJSON() ([]byte, error) {
    type Alias Limitcount

    if LimitcountMarshalled {
        return []byte("{}"), nil
    }
    LimitcountMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        EstimatedCount int `json:"estimatedCount"`
        
        Max int `json:"max"`
        
        EntityId string `json:"entityId"`
        
        UserId string `json:"userId"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

