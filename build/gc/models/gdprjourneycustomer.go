package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GdprjourneycustomerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GdprjourneycustomerDud struct { 
    


    

}

// Gdprjourneycustomer
type Gdprjourneycustomer struct { 
    // VarType - The type of the customerId within the Journey System (e.g. cookie). Required if `id` is defined.
    VarType string `json:"type"`


    // Id - An ID of a customer within the Journey System at a point-in-time. Required if `type` is defined.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Gdprjourneycustomer) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Gdprjourneycustomer) MarshalJSON() ([]byte, error) {
    type Alias Gdprjourneycustomer

    if GdprjourneycustomerMarshalled {
        return []byte("{}"), nil
    }
    GdprjourneycustomerMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

