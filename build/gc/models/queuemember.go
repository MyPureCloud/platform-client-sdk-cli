package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueuememberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueuememberDud struct { 
    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Queuemember
type Queuemember struct { 
    // Id - The queue member's id.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // User
    User User `json:"user"`


    // RingNumber
    RingNumber int `json:"ringNumber"`


    // Joined
    Joined bool `json:"joined"`


    // MemberBy
    MemberBy string `json:"memberBy"`


    // RoutingStatus
    RoutingStatus Routingstatus `json:"routingStatus"`


    

}

// String returns a JSON representation of the model
func (o *Queuemember) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queuemember) MarshalJSON() ([]byte, error) {
    type Alias Queuemember

    if QueuememberMarshalled {
        return []byte("{}"), nil
    }
    QueuememberMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        User User `json:"user"`
        
        RingNumber int `json:"ringNumber"`
        
        Joined bool `json:"joined"`
        
        MemberBy string `json:"memberBy"`
        
        RoutingStatus Routingstatus `json:"routingStatus"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

