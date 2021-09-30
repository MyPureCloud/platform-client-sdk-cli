package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatenotificationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatenotificationresponseDud struct { 
    


    

}

// Updatenotificationresponse
type Updatenotificationresponse struct { 
    // MutableGroupId - The mutableGroupId of the notification
    MutableGroupId string `json:"mutableGroupId"`


    // Id - The id of the notification for mapping the potentially new mutableGroupId
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Updatenotificationresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatenotificationresponse) MarshalJSON() ([]byte, error) {
    type Alias Updatenotificationresponse

    if UpdatenotificationresponseMarshalled {
        return []byte("{}"), nil
    }
    UpdatenotificationresponseMarshalled = true

    return json.Marshal(&struct { 
        MutableGroupId string `json:"mutableGroupId"`
        
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

