package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowoutcomedivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowoutcomedivisionviewDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowoutcomedivisionview
type Flowoutcomedivisionview struct { 
    // Id - The flow outcome identifier
    Id string `json:"id"`


    // Name - The flow outcome name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Flowoutcomedivisionview) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowoutcomedivisionview) MarshalJSON() ([]byte, error) {
    type Alias Flowoutcomedivisionview

    if FlowoutcomedivisionviewMarshalled {
        return []byte("{}"), nil
    }
    FlowoutcomedivisionviewMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

