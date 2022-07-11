package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingcampaignentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingcampaignentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Messagingcampaignentitylisting
type Messagingcampaignentitylisting struct { 
    // Entities
    Entities []Messagingcampaign `json:"entities"`


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


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Messagingcampaignentitylisting) String() string {
     o.Entities = []Messagingcampaign{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingcampaignentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Messagingcampaignentitylisting

    if MessagingcampaignentitylistingMarshalled {
        return []byte("{}"), nil
    }
    MessagingcampaignentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Messagingcampaign `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Messagingcampaign{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

