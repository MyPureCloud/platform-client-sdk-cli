package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowmilestonedivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowmilestonedivisionviewDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowmilestonedivisionview
type Flowmilestonedivisionview struct { 
    // Id - The flow milestone identifier
    Id string `json:"id"`


    // Name - The flow milestone name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Flowmilestonedivisionview) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowmilestonedivisionview) MarshalJSON() ([]byte, error) {
    type Alias Flowmilestonedivisionview

    if FlowmilestonedivisionviewMarshalled {
        return []byte("{}"), nil
    }
    FlowmilestonedivisionviewMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

