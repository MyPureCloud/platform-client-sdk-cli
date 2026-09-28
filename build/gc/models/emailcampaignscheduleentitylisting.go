package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailcampaignscheduleentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailcampaignscheduleentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Emailcampaignscheduleentitylisting
type Emailcampaignscheduleentitylisting struct { 
    // Entities
    Entities []Emailcampaignschedule `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Emailcampaignscheduleentitylisting) String() string {
     o.Entities = []Emailcampaignschedule{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailcampaignscheduleentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Emailcampaignscheduleentitylisting

    if EmailcampaignscheduleentitylistingMarshalled {
        return []byte("{}"), nil
    }
    EmailcampaignscheduleentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Emailcampaignschedule `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        LastUri string `json:"lastUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Emailcampaignschedule{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

