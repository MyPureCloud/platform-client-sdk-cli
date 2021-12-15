package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotsearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotsearchresponseDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Botsearchresponse
type Botsearchresponse struct { 
    // Id - The id of the bot
    Id string `json:"id"`


    // Name - The name of the bot
    Name string `json:"name"`


    // BotType - The provider of the bot
    BotType string `json:"botType"`


    // Description - The description of the bot
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Botsearchresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botsearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Botsearchresponse

    if BotsearchresponseMarshalled {
        return []byte("{}"), nil
    }
    BotsearchresponseMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        BotType string `json:"botType"`
        
        Description string `json:"description"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

