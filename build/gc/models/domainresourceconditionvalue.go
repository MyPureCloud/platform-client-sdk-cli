package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainresourceconditionvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainresourceconditionvalueDud struct { 
    


    


    


    

}

// Domainresourceconditionvalue
type Domainresourceconditionvalue struct { 
    // User
    User User `json:"user"`


    // Queue
    Queue Queue `json:"queue"`


    // Value
    Value string `json:"value"`


    // VarType
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Domainresourceconditionvalue) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainresourceconditionvalue) MarshalJSON() ([]byte, error) {
    type Alias Domainresourceconditionvalue

    if DomainresourceconditionvalueMarshalled {
        return []byte("{}"), nil
    }
    DomainresourceconditionvalueMarshalled = true

    return json.Marshal(&struct {
        
        User User `json:"user"`
        
        Queue Queue `json:"queue"`
        
        Value string `json:"value"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

