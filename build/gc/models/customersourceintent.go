package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomersourceintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomersourceintentDud struct { 
    


    

}

// Customersourceintent
type Customersourceintent struct { 
    // SourceIntent - Source intent data
    SourceIntent Sourceintent `json:"sourceIntent"`


    // CustomerIntent - Customer intent source intent is assigned to
    CustomerIntent Domainentityref `json:"customerIntent"`

}

// String returns a JSON representation of the model
func (o *Customersourceintent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customersourceintent) MarshalJSON() ([]byte, error) {
    type Alias Customersourceintent

    if CustomersourceintentMarshalled {
        return []byte("{}"), nil
    }
    CustomersourceintentMarshalled = true

    return json.Marshal(&struct {
        
        SourceIntent Sourceintent `json:"sourceIntent"`
        
        CustomerIntent Domainentityref `json:"customerIntent"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

