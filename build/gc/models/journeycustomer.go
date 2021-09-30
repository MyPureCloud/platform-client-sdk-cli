package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneycustomerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneycustomerDud struct { 
    


    

}

// Journeycustomer
type Journeycustomer struct { 
    // Id - An ID of a customer within the Journey System at a point-in-time.  Note that a customer entity can have multiple customerIds based on the stitching process.  Depending on the context within the PureCloud conversation, this may or may not be mutable.
    Id string `json:"id"`


    // IdType - The type of the customerId within the Journey System (e.g. cookie).
    IdType string `json:"idType"`

}

// String returns a JSON representation of the model
func (o *Journeycustomer) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeycustomer) MarshalJSON() ([]byte, error) {
    type Alias Journeycustomer

    if JourneycustomerMarshalled {
        return []byte("{}"), nil
    }
    JourneycustomerMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        IdType string `json:"idType"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

