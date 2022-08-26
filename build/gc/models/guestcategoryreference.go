package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuestcategoryreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuestcategoryreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Guestcategoryreference
type Guestcategoryreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Guestcategoryreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guestcategoryreference) MarshalJSON() ([]byte, error) {
    type Alias Guestcategoryreference

    if GuestcategoryreferenceMarshalled {
        return []byte("{}"), nil
    }
    GuestcategoryreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

