package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkerrornoteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkerrornoteDud struct { 
    


    


    


    


    


    

}

// Bulkerrornote
type Bulkerrornote struct { 
    // Code
    Code string `json:"code"`


    // Message
    Message string `json:"message"`


    // Status
    Status int `json:"status"`


    // Retryable
    Retryable bool `json:"retryable"`


    // Entity
    Entity Note `json:"entity"`


    // Details
    Details []Bulkerrordetail `json:"details"`

}

// String returns a JSON representation of the model
func (o *Bulkerrornote) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Details = []Bulkerrordetail{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkerrornote) MarshalJSON() ([]byte, error) {
    type Alias Bulkerrornote

    if BulkerrornoteMarshalled {
        return []byte("{}"), nil
    }
    BulkerrornoteMarshalled = true

    return json.Marshal(&struct { 
        Code string `json:"code"`
        
        Message string `json:"message"`
        
        Status int `json:"status"`
        
        Retryable bool `json:"retryable"`
        
        Entity Note `json:"entity"`
        
        Details []Bulkerrordetail `json:"details"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        Details: []Bulkerrordetail{{}},
        

        
        Alias: (*Alias)(u),
    })
}

