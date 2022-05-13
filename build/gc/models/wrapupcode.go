package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WrapupcodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WrapupcodeDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Wrapupcode
type Wrapupcode struct { 
    


    // Name - The wrap-up code name.
    Name string `json:"name"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy
    ModifiedBy string `json:"modifiedBy"`


    // CreatedBy
    CreatedBy string `json:"createdBy"`


    

}

// String returns a JSON representation of the model
func (o *Wrapupcode) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wrapupcode) MarshalJSON() ([]byte, error) {
    type Alias Wrapupcode

    if WrapupcodeMarshalled {
        return []byte("{}"), nil
    }
    WrapupcodeMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

