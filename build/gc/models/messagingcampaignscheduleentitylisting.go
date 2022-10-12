package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingcampaignscheduleentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingcampaignscheduleentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Messagingcampaignscheduleentitylisting
type Messagingcampaignscheduleentitylisting struct { 
    // Entities
    Entities []Messagingcampaignschedule `json:"entities"`


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
func (o *Messagingcampaignscheduleentitylisting) String() string {
     o.Entities = []Messagingcampaignschedule{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingcampaignscheduleentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Messagingcampaignscheduleentitylisting

    if MessagingcampaignscheduleentitylistingMarshalled {
        return []byte("{}"), nil
    }
    MessagingcampaignscheduleentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Messagingcampaignschedule `json:"entities"`
        
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

        
        Entities: []Messagingcampaignschedule{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

