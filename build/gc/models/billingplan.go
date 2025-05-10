package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingplanDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    


    


    

}

// Billingplan
type Billingplan struct { 
    


    


    // Organizations - List of organizations for the plan.
    Organizations []Namedentity `json:"organizations"`


    // Product - Represents the details of a product.
    Product Billingproduct `json:"product"`


    // Items - List of items for the plan.
    Items []Billingplanitem `json:"items"`

}

// String returns a JSON representation of the model
func (o *Billingplan) String() string {
     o.Organizations = []Namedentity{{}} 
    
     o.Items = []Billingplanitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingplan) MarshalJSON() ([]byte, error) {
    type Alias Billingplan

    if BillingplanMarshalled {
        return []byte("{}"), nil
    }
    BillingplanMarshalled = true

    return json.Marshal(&struct {
        
        Organizations []Namedentity `json:"organizations"`
        
        Product Billingproduct `json:"product"`
        
        Items []Billingplanitem `json:"items"`
        *Alias
    }{

        


        


        
        Organizations: []Namedentity{{}},
        


        


        
        Items: []Billingplanitem{{}},
        

        Alias: (*Alias)(u),
    })
}

