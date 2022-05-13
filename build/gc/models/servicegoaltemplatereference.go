package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServicegoaltemplatereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServicegoaltemplatereferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Servicegoaltemplatereference
type Servicegoaltemplatereference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Servicegoaltemplatereference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Servicegoaltemplatereference) MarshalJSON() ([]byte, error) {
    type Alias Servicegoaltemplatereference

    if ServicegoaltemplatereferenceMarshalled {
        return []byte("{}"), nil
    }
    ServicegoaltemplatereferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

