package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkerrordetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkerrordetailDud struct { 
    


    


    

}

// Bulkerrordetail
type Bulkerrordetail struct { 
    // FieldName - The name of the field which experienced an error.
    FieldName string `json:"fieldName"`


    // Value - The field value from the request which caused the error.
    Value string `json:"value"`


    // Message - A field-specific error message describing why this operation was rejected.
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Bulkerrordetail) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkerrordetail) MarshalJSON() ([]byte, error) {
    type Alias Bulkerrordetail

    if BulkerrordetailMarshalled {
        return []byte("{}"), nil
    }
    BulkerrordetailMarshalled = true

    return json.Marshal(&struct {
        
        FieldName string `json:"fieldName"`
        
        Value string `json:"value"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

