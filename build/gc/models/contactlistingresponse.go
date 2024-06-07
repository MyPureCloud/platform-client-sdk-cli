package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistingresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistingresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Contactlistingresponse
type Contactlistingresponse struct { 
    // Entities
    Entities []Dialercontact `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // ContactsCount
    ContactsCount int `json:"contactsCount"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Contactlistingresponse) String() string {
     o.Entities = []Dialercontact{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistingresponse) MarshalJSON() ([]byte, error) {
    type Alias Contactlistingresponse

    if ContactlistingresponseMarshalled {
        return []byte("{}"), nil
    }
    ContactlistingresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Dialercontact `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        ContactsCount int `json:"contactsCount"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Dialercontact{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

