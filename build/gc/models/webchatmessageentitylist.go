package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatmessageentitylistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatmessageentitylistDud struct { 
    


    


    


    


    

}

// Webchatmessageentitylist
type Webchatmessageentitylist struct { 
    // PageSize
    PageSize int `json:"pageSize"`


    // Entities
    Entities []Webchatmessage `json:"entities"`


    // PreviousPage
    PreviousPage string `json:"previousPage"`


    // Next
    Next string `json:"next"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Webchatmessageentitylist) String() string {
    
     o.Entities = []Webchatmessage{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatmessageentitylist) MarshalJSON() ([]byte, error) {
    type Alias Webchatmessageentitylist

    if WebchatmessageentitylistMarshalled {
        return []byte("{}"), nil
    }
    WebchatmessageentitylistMarshalled = true

    return json.Marshal(&struct {
        
        PageSize int `json:"pageSize"`
        
        Entities []Webchatmessage `json:"entities"`
        
        PreviousPage string `json:"previousPage"`
        
        Next string `json:"next"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Webchatmessage{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

