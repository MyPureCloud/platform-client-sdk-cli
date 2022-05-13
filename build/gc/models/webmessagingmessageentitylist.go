package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingmessageentitylistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingmessageentitylistDud struct { 
    


    


    


    


    

}

// Webmessagingmessageentitylist
type Webmessagingmessageentitylist struct { 
    // Entities
    Entities []Webmessagingmessage `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Webmessagingmessageentitylist) String() string {
     o.Entities = []Webmessagingmessage{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingmessageentitylist) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingmessageentitylist

    if WebmessagingmessageentitylistMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingmessageentitylistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Webmessagingmessage `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Webmessagingmessage{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

