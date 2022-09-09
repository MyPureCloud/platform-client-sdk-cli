package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserexternalidentifierMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserexternalidentifierDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Userexternalidentifier - Defines a link between an External Identifier and Authority pair to a Entity Type and Entity Identifier pair. Represents the two way, one to one mapping of an External Authority or System of Record's identifier to a PureCloud entity. e.g. (ExternalId='05001',Authority='XyzCRM') to (entityType=user,entityId='8eb03b33-3acb-4bc1-a244-50b9b9f19495')
type Userexternalidentifier struct { 
    // AuthorityName - Authority or System of Record which owns the External Identifier
    AuthorityName string `json:"authorityName"`


    // ExternalKey - External Key
    ExternalKey string `json:"externalKey"`


    

}

// String returns a JSON representation of the model
func (o *Userexternalidentifier) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userexternalidentifier) MarshalJSON() ([]byte, error) {
    type Alias Userexternalidentifier

    if UserexternalidentifierMarshalled {
        return []byte("{}"), nil
    }
    UserexternalidentifierMarshalled = true

    return json.Marshal(&struct {
        
        AuthorityName string `json:"authorityName"`
        
        ExternalKey string `json:"externalKey"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

