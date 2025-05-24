package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentpaymentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentpaymentrequestDud struct { 
    PaymentPlatform string `json:"paymentPlatform"`


    


    


    


    


    


    


    


    

}

// Conversationcontentpaymentrequest - Payment Request object used to request payment from a customer.
type Conversationcontentpaymentrequest struct { 
    


    // CountryCode - The merchant's two-letter ISO 3166 country code.
    CountryCode string `json:"countryCode"`


    // CurrencyCode - The three-letter ISO 4217 currency code for the payment.
    CurrencyCode string `json:"currencyCode"`


    // OrderTotal - The total price of the order.
    OrderTotal float64 `json:"orderTotal"`


    // LineItems - The items that make up the order.
    LineItems []Conversationcontentlineitem `json:"lineItems"`


    // ShippingOptions - The available shipping options.
    ShippingOptions []Conversationcontentlineitem `json:"shippingOptions"`


    // RequiredContactFields - Contact fields required to complete the order.
    RequiredContactFields []Conversationcontentrequiredcontactfield `json:"requiredContactFields"`


    // ReceivedMessage - The message prompt to complete a payment transaction.
    ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`


    // ReplyMessage - The reply message after the user has completed the payment transaction.
    ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentpaymentrequest) String() string {
    
    
    
     o.LineItems = []Conversationcontentlineitem{{}} 
     o.ShippingOptions = []Conversationcontentlineitem{{}} 
     o.RequiredContactFields = []Conversationcontentrequiredcontactfield{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentpaymentrequest) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentpaymentrequest

    if ConversationcontentpaymentrequestMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentpaymentrequestMarshalled = true

    return json.Marshal(&struct {
        
        CountryCode string `json:"countryCode"`
        
        CurrencyCode string `json:"currencyCode"`
        
        OrderTotal float64 `json:"orderTotal"`
        
        LineItems []Conversationcontentlineitem `json:"lineItems"`
        
        ShippingOptions []Conversationcontentlineitem `json:"shippingOptions"`
        
        RequiredContactFields []Conversationcontentrequiredcontactfield `json:"requiredContactFields"`
        
        ReceivedMessage Conversationcontentreceivedreplymessage `json:"receivedMessage"`
        
        ReplyMessage Conversationcontentreceivedreplymessage `json:"replyMessage"`
        *Alias
    }{

        


        


        


        


        
        LineItems: []Conversationcontentlineitem{{}},
        


        
        ShippingOptions: []Conversationcontentlineitem{{}},
        


        
        RequiredContactFields: []Conversationcontentrequiredcontactfield{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

