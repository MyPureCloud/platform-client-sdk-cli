package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillinginvoiceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillinginvoiceDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Billinginvoice
type Billinginvoice struct { 
    


    // BillToCustomer - The bill to customer.
    BillToCustomer Customer `json:"billToCustomer"`


    // ShipToCustomer - The ship to customer.
    ShipToCustomer Customer `json:"shipToCustomer"`


    // SoldToCustomer - The sold to customer.
    SoldToCustomer Customer `json:"soldToCustomer"`


    // DateInvoiced - Date when the invoice was issued. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateInvoiced time.Time `json:"dateInvoiced"`


    // BillToAddress - Details of the bill to address.
    BillToAddress Invoiceaddress `json:"billToAddress"`


    // ShipToAddress - Details of the ship to address.
    ShipToAddress Invoiceaddress `json:"shipToAddress"`


    // CurrencyIsoCode - Contains the ISO code for any currency allowed by the organization.
    CurrencyIsoCode string `json:"currencyIsoCode"`


    // PaymentStatus - Status of the payment.
    PaymentStatus string `json:"paymentStatus"`


    // PaymentTerms - Payment terms.
    PaymentTerms string `json:"paymentTerms"`


    // PaymentLink - Payment link.
    PaymentLink string `json:"paymentLink"`


    // CustomerPoNumber - Purchase Order Number.
    CustomerPoNumber string `json:"customerPoNumber"`


    // CustomerInvoiceType - Type of the invoice.
    CustomerInvoiceType string `json:"customerInvoiceType"`


    // Amount - Amount.
    Amount float32 `json:"amount"`

}

// String returns a JSON representation of the model
func (o *Billinginvoice) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billinginvoice) MarshalJSON() ([]byte, error) {
    type Alias Billinginvoice

    if BillinginvoiceMarshalled {
        return []byte("{}"), nil
    }
    BillinginvoiceMarshalled = true

    return json.Marshal(&struct {
        
        BillToCustomer Customer `json:"billToCustomer"`
        
        ShipToCustomer Customer `json:"shipToCustomer"`
        
        SoldToCustomer Customer `json:"soldToCustomer"`
        
        DateInvoiced time.Time `json:"dateInvoiced"`
        
        BillToAddress Invoiceaddress `json:"billToAddress"`
        
        ShipToAddress Invoiceaddress `json:"shipToAddress"`
        
        CurrencyIsoCode string `json:"currencyIsoCode"`
        
        PaymentStatus string `json:"paymentStatus"`
        
        PaymentTerms string `json:"paymentTerms"`
        
        PaymentLink string `json:"paymentLink"`
        
        CustomerPoNumber string `json:"customerPoNumber"`
        
        CustomerInvoiceType string `json:"customerInvoiceType"`
        
        Amount float32 `json:"amount"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

