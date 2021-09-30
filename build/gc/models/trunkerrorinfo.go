package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkerrorinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkerrorinfoDud struct { 
    


    


    

}

// Trunkerrorinfo
type Trunkerrorinfo struct { 
    // Text
    Text string `json:"text"`


    // Code
    Code string `json:"code"`


    // Details
    Details Trunkerrorinfodetails `json:"details"`

}

// String returns a JSON representation of the model
func (o *Trunkerrorinfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkerrorinfo) MarshalJSON() ([]byte, error) {
    type Alias Trunkerrorinfo

    if TrunkerrorinfoMarshalled {
        return []byte("{}"), nil
    }
    TrunkerrorinfoMarshalled = true

    return json.Marshal(&struct { 
        Text string `json:"text"`
        
        Code string `json:"code"`
        
        Details Trunkerrorinfodetails `json:"details"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

