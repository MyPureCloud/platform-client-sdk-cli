package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultnotenoteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultnotenoteDud struct { 
    


    


    


    

}

// Bulkresponseresultnotenote
type Bulkresponseresultnotenote struct { 
    // Id
    Id string `json:"id"`


    // Success
    Success bool `json:"success"`


    // Entity
    Entity Note `json:"entity"`


    // VarError
    VarError Bulkerrornote `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultnotenote) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultnotenote) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultnotenote

    if BulkresponseresultnotenoteMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultnotenoteMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity Note `json:"entity"`
        
        VarError Bulkerrornote `json:"error"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

