package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestprizesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestprizesDud struct { 
    


    


    


    


    ImageUrl string `json:"imageUrl"`


    

}

// Contestprizes
type Contestprizes struct { 
    // Tier - The Contest Prize tier
    Tier int `json:"tier"`


    // Name - The Contest Prize name
    Name string `json:"name"`


    // Description - The Contest Prize description
    Description string `json:"description"`


    // ImageId - The Contest Prize image id
    ImageId string `json:"imageId"`


    


    // WinnersCount - The Contest Prize winner Count
    WinnersCount int `json:"winnersCount"`

}

// String returns a JSON representation of the model
func (o *Contestprizes) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestprizes) MarshalJSON() ([]byte, error) {
    type Alias Contestprizes

    if ContestprizesMarshalled {
        return []byte("{}"), nil
    }
    ContestprizesMarshalled = true

    return json.Marshal(&struct {
        
        Tier int `json:"tier"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        ImageId string `json:"imageId"`
        
        WinnersCount int `json:"winnersCount"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

