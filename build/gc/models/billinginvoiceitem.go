package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillinginvoiceitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillinginvoiceitemDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    

}

// Billinginvoiceitem
type Billinginvoiceitem struct { 
    


    // Product - Represents the details of a product.
    Product Billingproduct `json:"product"`


    // Description - Line Item Description.
    Description string `json:"description"`


    // DateTransacted - Invoice transaction date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateTransacted time.Time `json:"dateTransacted"`


    // DateStart - Invoice start date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`


    // DateEnd - Invoice end date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEnd time.Time `json:"dateEnd"`


    // Organization - A Genesys Cloud Organization.
    Organization Namedentity `json:"organization"`


    // Quantity - Quantity of the item.
    Quantity int `json:"quantity"`


    // UnitOfMeasure - Unit of Measure.
    UnitOfMeasure string `json:"unitOfMeasure"`


    // Amount - Amount.
    Amount float32 `json:"amount"`

}

// String returns a JSON representation of the model
func (o *Billinginvoiceitem) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billinginvoiceitem) MarshalJSON() ([]byte, error) {
    type Alias Billinginvoiceitem

    if BillinginvoiceitemMarshalled {
        return []byte("{}"), nil
    }
    BillinginvoiceitemMarshalled = true

    return json.Marshal(&struct {
        
        Product Billingproduct `json:"product"`
        
        Description string `json:"description"`
        
        DateTransacted time.Time `json:"dateTransacted"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        Organization Namedentity `json:"organization"`
        
        Quantity int `json:"quantity"`
        
        UnitOfMeasure string `json:"unitOfMeasure"`
        
        Amount float32 `json:"amount"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

