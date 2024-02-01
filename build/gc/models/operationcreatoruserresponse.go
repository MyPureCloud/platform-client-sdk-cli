package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OperationcreatoruserresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OperationcreatoruserresponseDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Operationcreatoruserresponse
type Operationcreatoruserresponse struct { 
    


    // VarType - Type of the operation creator entity: User or OAuthClient
    VarType string `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Operationcreatoruserresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Operationcreatoruserresponse) MarshalJSON() ([]byte, error) {
    type Alias Operationcreatoruserresponse

    if OperationcreatoruserresponseMarshalled {
        return []byte("{}"), nil
    }
    OperationcreatoruserresponseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

