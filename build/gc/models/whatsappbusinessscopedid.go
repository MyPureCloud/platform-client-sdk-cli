package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappbusinessscopedidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappbusinessscopedidDud struct { 
    


    


    

}

// Whatsappbusinessscopedid - A WhatsAppBusinessScopedId record
type Whatsappbusinessscopedid struct { 
    // Bsuid - The Whatsapp Business Scoped Id associated with a WhatsApp User.
    Bsuid string `json:"bsuid"`


    // Parent - True if this is a parent BSUID, false if it is a regular BSUID.
    Parent bool `json:"parent"`


    // PortfolioId - The portfolio id associated with the WhatsApp bsuid.
    PortfolioId string `json:"portfolioId"`

}

// String returns a JSON representation of the model
func (o *Whatsappbusinessscopedid) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappbusinessscopedid) MarshalJSON() ([]byte, error) {
    type Alias Whatsappbusinessscopedid

    if WhatsappbusinessscopedidMarshalled {
        return []byte("{}"), nil
    }
    WhatsappbusinessscopedidMarshalled = true

    return json.Marshal(&struct {
        
        Bsuid string `json:"bsuid"`
        
        Parent bool `json:"parent"`
        
        PortfolioId string `json:"portfolioId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

