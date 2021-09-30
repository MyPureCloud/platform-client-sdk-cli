package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmbusinessunitreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmbusinessunitreferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Wfmbusinessunitreference
type Wfmbusinessunitreference struct { 
    // Id - The ID of the business unit
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Wfmbusinessunitreference) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmbusinessunitreference) MarshalJSON() ([]byte, error) {
    type Alias Wfmbusinessunitreference

    if WfmbusinessunitreferenceMarshalled {
        return []byte("{}"), nil
    }
    WfmbusinessunitreferenceMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

