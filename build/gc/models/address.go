package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddressDud struct { 
    


    


    


    


    

}

// Address
type Address struct { 
    // Name - This will be nameRaw if present, or a locality lookup of the address field otherwise.
    Name string `json:"name"`


    // NameRaw - The name as close to the bits on the wire as possible.
    NameRaw string `json:"nameRaw"`


    // AddressNormalized - The normalized address. This field is acquired from the Address Normalization Table.  The addressRaw could have gone through some transformations, such as only using the numeric portion, before being run through the Address Normalization Table.
    AddressNormalized string `json:"addressNormalized"`


    // AddressRaw - The address as close to the bits on the wire as possible.
    AddressRaw string `json:"addressRaw"`


    // AddressDisplayable - The displayable address. This field is acquired from the Address Normalization Table.  The addressRaw could have gone through some transformations, such as only using the numeric portion, before being run through the Address Normalization Table.
    AddressDisplayable string `json:"addressDisplayable"`

}

// String returns a JSON representation of the model
func (o *Address) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Address) MarshalJSON() ([]byte, error) {
    type Alias Address

    if AddressMarshalled {
        return []byte("{}"), nil
    }
    AddressMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        NameRaw string `json:"nameRaw"`
        
        AddressNormalized string `json:"addressNormalized"`
        
        AddressRaw string `json:"addressRaw"`
        
        AddressDisplayable string `json:"addressDisplayable"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

