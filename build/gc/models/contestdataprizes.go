package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestdataprizesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestdataprizesDud struct { 
    


    


    

}

// Contestdataprizes
type Contestdataprizes struct { 
    // Tier - Tier of the prize
    Tier int `json:"tier"`


    // Name - Name of the prize
    Name string `json:"name"`


    // ImageId - Id of the prize image
    ImageId string `json:"imageId"`

}

// String returns a JSON representation of the model
func (o *Contestdataprizes) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestdataprizes) MarshalJSON() ([]byte, error) {
    type Alias Contestdataprizes

    if ContestdataprizesMarshalled {
        return []byte("{}"), nil
    }
    ContestdataprizesMarshalled = true

    return json.Marshal(&struct {
        
        Tier int `json:"tier"`
        
        Name string `json:"name"`
        
        ImageId string `json:"imageId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

