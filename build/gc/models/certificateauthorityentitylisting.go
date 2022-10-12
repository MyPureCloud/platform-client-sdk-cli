package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CertificateauthorityentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CertificateauthorityentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Certificateauthorityentitylisting
type Certificateauthorityentitylisting struct { 
    // Entities
    Entities []Domaincertificateauthority `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Certificateauthorityentitylisting) String() string {
     o.Entities = []Domaincertificateauthority{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Certificateauthorityentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Certificateauthorityentitylisting

    if CertificateauthorityentitylistingMarshalled {
        return []byte("{}"), nil
    }
    CertificateauthorityentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Domaincertificateauthority `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Domaincertificateauthority{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

