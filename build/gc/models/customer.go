package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomerDud struct { 
    Id string `json:"id"`


    


    

}

// Customer
type Customer struct { 
    


    // Number - The unique number of the customer.
    Number string `json:"number"`


    // Name - The name of the customer.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Customer) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customer) MarshalJSON() ([]byte, error) {
    type Alias Customer

    if CustomerMarshalled {
        return []byte("{}"), nil
    }
    CustomerMarshalled = true

    return json.Marshal(&struct {
        
        Number string `json:"number"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

