package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowoutcomeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowoutcomeDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowoutcome
type Flowoutcome struct { 
    // Id - The flow outcome identifier
    Id string `json:"id"`


    // Name - The flow outcome name.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description
    Description string `json:"description"`


    // CurrentOperation
    CurrentOperation Operation `json:"currentOperation"`


    

}

// String returns a JSON representation of the model
func (o *Flowoutcome) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowoutcome) MarshalJSON() ([]byte, error) {
    type Alias Flowoutcome

    if FlowoutcomeMarshalled {
        return []byte("{}"), nil
    }
    FlowoutcomeMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        CurrentOperation Operation `json:"currentOperation"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

