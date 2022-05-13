package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DetailDud struct { 
    


    


    


    

}

// Detail
type Detail struct { 
    // ErrorCode
    ErrorCode string `json:"errorCode"`


    // FieldName
    FieldName string `json:"fieldName"`


    // EntityId
    EntityId string `json:"entityId"`


    // EntityName
    EntityName string `json:"entityName"`

}

// String returns a JSON representation of the model
func (o *Detail) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Detail) MarshalJSON() ([]byte, error) {
    type Alias Detail

    if DetailMarshalled {
        return []byte("{}"), nil
    }
    DetailMarshalled = true

    return json.Marshal(&struct {
        
        ErrorCode string `json:"errorCode"`
        
        FieldName string `json:"fieldName"`
        
        EntityId string `json:"entityId"`
        
        EntityName string `json:"entityName"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

