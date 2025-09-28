package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ApplepayMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ApplepayDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Applepay
type Applepay struct { 
    // StoreName - The name of the store.
    StoreName string `json:"storeName"`


    // MerchantId - The stores merchant identifier.
    MerchantId string `json:"merchantId"`


    // DomainName - The domain name associated with the merchant account.
    DomainName string `json:"domainName"`


    // PaymentCapabilities - The payment capabilities supported by the merchant.
    PaymentCapabilities []string `json:"paymentCapabilities"`


    // SupportedPaymentNetworks - The payment networks supported by the merchant.
    SupportedPaymentNetworks []string `json:"supportedPaymentNetworks"`


    // PaymentCertificateCredentialId - The Genesys credentialId the payment certificates are stored under.
    PaymentCertificateCredentialId string `json:"paymentCertificateCredentialId"`


    // PaymentGatewayUrl - The url used to process payments.
    PaymentGatewayUrl string `json:"paymentGatewayUrl"`


    // FallbackUrl - The url opened in a web browser if the customers device is unable to make payments using Apple Pay.
    FallbackUrl string `json:"fallbackUrl"`


    // ShippingMethodUpdateUrl - The url called when the customer changes the shipping method.
    ShippingMethodUpdateUrl string `json:"shippingMethodUpdateUrl"`


    // ShippingContactUpdateUrl - The url called when the customer changes their shipping address information.
    ShippingContactUpdateUrl string `json:"shippingContactUpdateUrl"`


    // PaymentMethodUpdateUrl - The url called when the customer changes their payment method.
    PaymentMethodUpdateUrl string `json:"paymentMethodUpdateUrl"`


    // OrderTrackingUrl - The url called after completing the order to update the order information in your system
    OrderTrackingUrl string `json:"orderTrackingUrl"`

}

// String returns a JSON representation of the model
func (o *Applepay) String() string {
    
    
    
     o.PaymentCapabilities = []string{""} 
     o.SupportedPaymentNetworks = []string{""} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Applepay) MarshalJSON() ([]byte, error) {
    type Alias Applepay

    if ApplepayMarshalled {
        return []byte("{}"), nil
    }
    ApplepayMarshalled = true

    return json.Marshal(&struct {
        
        StoreName string `json:"storeName"`
        
        MerchantId string `json:"merchantId"`
        
        DomainName string `json:"domainName"`
        
        PaymentCapabilities []string `json:"paymentCapabilities"`
        
        SupportedPaymentNetworks []string `json:"supportedPaymentNetworks"`
        
        PaymentCertificateCredentialId string `json:"paymentCertificateCredentialId"`
        
        PaymentGatewayUrl string `json:"paymentGatewayUrl"`
        
        FallbackUrl string `json:"fallbackUrl"`
        
        ShippingMethodUpdateUrl string `json:"shippingMethodUpdateUrl"`
        
        ShippingContactUpdateUrl string `json:"shippingContactUpdateUrl"`
        
        PaymentMethodUpdateUrl string `json:"paymentMethodUpdateUrl"`
        
        OrderTrackingUrl string `json:"orderTrackingUrl"`
        *Alias
    }{

        


        


        


        
        PaymentCapabilities: []string{""},
        


        
        SupportedPaymentNetworks: []string{""},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

