package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BasetopicentitiyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BasetopicentitiyDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Basetopicentitiy
type Basetopicentitiy struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Basetopicentitiy) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Basetopicentitiy) MarshalJSON() ([]byte, error) {
    type Alias Basetopicentitiy

    if BasetopicentitiyMarshalled {
        return []byte("{}"), nil
    }
    BasetopicentitiyMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

