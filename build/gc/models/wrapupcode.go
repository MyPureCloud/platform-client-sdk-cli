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


    // Division - The division to which this entity belongs.
    Division Starrabledivision `json:"division"`


    // Description - The wrap-up code description.
    Description string `json:"description"`


    // DateCreated - Date when the wrap-up code was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date when the wrap-up code was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy - The ID of the user that created the wrap-up code.
    CreatedBy string `json:"createdBy"`


    // ModifiedBy - The ID of the user that modified the wrap-up code.
    ModifiedBy string `json:"modifiedBy"`


    

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
        
        Division Starrabledivision `json:"division"`
        
        Description string `json:"description"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy string `json:"createdBy"`
        
        ModifiedBy string `json:"modifiedBy"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

