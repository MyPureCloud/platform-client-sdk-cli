package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServicecontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServicecontextDud struct { 
    

}

// Servicecontext
type Servicecontext struct { 
    // Name - Unused field for the purpose of ensuring a Swagger definition is created for a class with only @JsonIgnore members.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Servicecontext) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Servicecontext) MarshalJSON() ([]byte, error) {
    type Alias Servicecontext

    if ServicecontextMarshalled {
        return []byte("{}"), nil
    }
    ServicecontextMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

