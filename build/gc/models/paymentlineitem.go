package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PaymentlineitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PaymentlineitemDud struct { 
    


    

}

// Paymentlineitem
type Paymentlineitem struct { 
    // Name - The display name for the item.
    Name string `json:"name"`


    // Price - The price of the item.
    Price float64 `json:"price"`

}

// String returns a JSON representation of the model
func (o *Paymentlineitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Paymentlineitem) MarshalJSON() ([]byte, error) {
    type Alias Paymentlineitem

    if PaymentlineitemMarshalled {
        return []byte("{}"), nil
    }
    PaymentlineitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Price float64 `json:"price"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

