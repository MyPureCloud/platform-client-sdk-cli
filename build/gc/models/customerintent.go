package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomerintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomerintentDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Customerintent
type Customerintent struct { 
    


    // Name - Name of the customer intent
    Name string `json:"name"`


    // Description - Description of the customer intent
    Description string `json:"description"`


    // ExpiryTime - Expiry time in hours of the customer intent
    ExpiryTime int `json:"expiryTime"`


    // CategoryId - CategoryId of the customer intent
    CategoryId string `json:"categoryId"`


    

}

// String returns a JSON representation of the model
func (o *Customerintent) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customerintent) MarshalJSON() ([]byte, error) {
    type Alias Customerintent

    if CustomerintentMarshalled {
        return []byte("{}"), nil
    }
    CustomerintentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        ExpiryTime int `json:"expiryTime"`
        
        CategoryId string `json:"categoryId"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

