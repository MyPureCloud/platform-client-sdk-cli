package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingcontractlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingcontractlistingDud struct { 
    


    


    


    

}

// Billingcontractlisting
type Billingcontractlisting struct { 
    // Entities
    Entities []Billingcontract `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Billingcontractlisting) String() string {
     o.Entities = []Billingcontract{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingcontractlisting) MarshalJSON() ([]byte, error) {
    type Alias Billingcontractlisting

    if BillingcontractlistingMarshalled {
        return []byte("{}"), nil
    }
    BillingcontractlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Billingcontract `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Billingcontract{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

