package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomerintentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomerintentresponseDud struct { 
    Id string `json:"id"`


    


    


    


    Category Addressableentityref `json:"category"`


    DateCreated time.Time `json:"dateCreated"`


    SelfUri string `json:"selfUri"`

}

// Customerintentresponse
type Customerintentresponse struct { 
    


    // Name - Name of the customer intent
    Name string `json:"name"`


    // Description - Description of the customer intent
    Description string `json:"description"`


    // ExpiryTime - Expiry time in hours of the customer intent
    ExpiryTime int `json:"expiryTime"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Customerintentresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customerintentresponse) MarshalJSON() ([]byte, error) {
    type Alias Customerintentresponse

    if CustomerintentresponseMarshalled {
        return []byte("{}"), nil
    }
    CustomerintentresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        ExpiryTime int `json:"expiryTime"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

