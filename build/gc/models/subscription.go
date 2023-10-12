package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SubscriptionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SubscriptionDud struct { 
    


    

}

// Subscription
type Subscription struct { 
    // Id - The subscription id
    Id string `json:"id"`


    // Topic - Notification topic
    Topic string `json:"topic"`

}

// String returns a JSON representation of the model
func (o *Subscription) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Subscription) MarshalJSON() ([]byte, error) {
    type Alias Subscription

    if SubscriptionMarshalled {
        return []byte("{}"), nil
    }
    SubscriptionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Topic string `json:"topic"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

