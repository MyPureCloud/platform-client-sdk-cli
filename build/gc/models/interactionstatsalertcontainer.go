package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InteractionstatsalertcontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InteractionstatsalertcontainerDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Interactionstatsalertcontainer
type Interactionstatsalertcontainer struct { 
    // Entities
    Entities []Interactionstatsalert `json:"entities"`


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


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Interactionstatsalertcontainer) String() string {
     o.Entities = []Interactionstatsalert{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Interactionstatsalertcontainer) MarshalJSON() ([]byte, error) {
    type Alias Interactionstatsalertcontainer

    if InteractionstatsalertcontainerMarshalled {
        return []byte("{}"), nil
    }
    InteractionstatsalertcontainerMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Interactionstatsalert `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Interactionstatsalert{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

