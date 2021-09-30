package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KeyrotationscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KeyrotationscheduleDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Keyrotationschedule
type Keyrotationschedule struct { 
    


    // Name
    Name string `json:"name"`


    // Period - Value to set schedule to
    Period string `json:"period"`


    

}

// String returns a JSON representation of the model
func (o *Keyrotationschedule) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Keyrotationschedule) MarshalJSON() ([]byte, error) {
    type Alias Keyrotationschedule

    if KeyrotationscheduleMarshalled {
        return []byte("{}"), nil
    }
    KeyrotationscheduleMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Period string `json:"period"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

