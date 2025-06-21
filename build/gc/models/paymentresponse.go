package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PaymentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PaymentresponseDud struct { 
    


    


    

}

// Paymentresponse
type Paymentresponse struct { 
    // OriginatingMessageId - Reference to the ID of the original payment request message this response is for.
    OriginatingMessageId string `json:"originatingMessageId"`


    // PaymentStatus - The status of the payment transaction.
    PaymentStatus string `json:"paymentStatus"`


    // FailureReason - The reason the payment request failed.
    FailureReason string `json:"failureReason"`

}

// String returns a JSON representation of the model
func (o *Paymentresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Paymentresponse) MarshalJSON() ([]byte, error) {
    type Alias Paymentresponse

    if PaymentresponseMarshalled {
        return []byte("{}"), nil
    }
    PaymentresponseMarshalled = true

    return json.Marshal(&struct {
        
        OriginatingMessageId string `json:"originatingMessageId"`
        
        PaymentStatus string `json:"paymentStatus"`
        
        FailureReason string `json:"failureReason"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

