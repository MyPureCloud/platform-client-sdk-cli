package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DestinationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DestinationDud struct { 
    


    


    


    

}

// Destination
type Destination struct { 
    // Address - Address or phone number.
    Address string `json:"address"`


    // Name - The name of the internal user.
    Name string `json:"name"`


    // UserId - The user ID.
    UserId string `json:"userId"`


    // QueueId - The queue ID.
    QueueId string `json:"queueId"`

}

// String returns a JSON representation of the model
func (o *Destination) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Destination) MarshalJSON() ([]byte, error) {
    type Alias Destination

    if DestinationMarshalled {
        return []byte("{}"), nil
    }
    DestinationMarshalled = true

    return json.Marshal(&struct {
        
        Address string `json:"address"`
        
        Name string `json:"name"`
        
        UserId string `json:"userId"`
        
        QueueId string `json:"queueId"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

