package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PropertychangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PropertychangeDud struct { 
    


    


    

}

// Propertychange
type Propertychange struct { 
    // Property - The property that was changed
    Property string `json:"property"`


    // OldValues - Previous values for the property.
    OldValues []string `json:"oldValues"`


    // NewValues - New values for the property.
    NewValues []string `json:"newValues"`

}

// String returns a JSON representation of the model
func (o *Propertychange) String() string {
    
    
    
    
    
    
     o.OldValues = []string{""} 
    
    
    
     o.NewValues = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Propertychange) MarshalJSON() ([]byte, error) {
    type Alias Propertychange

    if PropertychangeMarshalled {
        return []byte("{}"), nil
    }
    PropertychangeMarshalled = true

    return json.Marshal(&struct { 
        Property string `json:"property"`
        
        OldValues []string `json:"oldValues"`
        
        NewValues []string `json:"newValues"`
        
        *Alias
    }{
        

        

        

        
        OldValues: []string{""},
        

        

        
        NewValues: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

