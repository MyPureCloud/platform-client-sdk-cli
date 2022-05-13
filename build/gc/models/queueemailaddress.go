package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueemailaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueemailaddressDud struct { 
    


    

}

// Queueemailaddress
type Queueemailaddress struct { 
    // Domain
    Domain Domainentityref `json:"domain"`


    // Route
    Route Inboundroute `json:"route"`

}

// String returns a JSON representation of the model
func (o *Queueemailaddress) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueemailaddress) MarshalJSON() ([]byte, error) {
    type Alias Queueemailaddress

    if QueueemailaddressMarshalled {
        return []byte("{}"), nil
    }
    QueueemailaddressMarshalled = true

    return json.Marshal(&struct {
        
        Domain Domainentityref `json:"domain"`
        
        Route Inboundroute `json:"route"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

