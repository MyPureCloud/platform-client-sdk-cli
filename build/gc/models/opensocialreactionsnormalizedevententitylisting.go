package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialreactionsnormalizedevententitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialreactionsnormalizedevententitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Opensocialreactionsnormalizedevententitylisting
type Opensocialreactionsnormalizedevententitylisting struct { 
    // Entities
    Entities []Opensocialmediareactionsnormalizedevent `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Opensocialreactionsnormalizedevententitylisting) String() string {
     o.Entities = []Opensocialmediareactionsnormalizedevent{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialreactionsnormalizedevententitylisting) MarshalJSON() ([]byte, error) {
    type Alias Opensocialreactionsnormalizedevententitylisting

    if OpensocialreactionsnormalizedevententitylistingMarshalled {
        return []byte("{}"), nil
    }
    OpensocialreactionsnormalizedevententitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Opensocialmediareactionsnormalizedevent `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Opensocialmediareactionsnormalizedevent{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

