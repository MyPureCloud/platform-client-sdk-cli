package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Lexv2botMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Lexv2botDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Lexv2bot
type Lexv2bot struct { 
    


    // Name
    Name string `json:"name"`


    // BotId - Lex V2 bot Id
    BotId string `json:"botId"`


    // Region - A region of the Lex V2 bot
    Region string `json:"region"`


    // Description - A description of the Lex V2 bot
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Lexv2bot) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexv2bot) MarshalJSON() ([]byte, error) {
    type Alias Lexv2bot

    if Lexv2botMarshalled {
        return []byte("{}"), nil
    }
    Lexv2botMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        BotId string `json:"botId"`
        
        Region string `json:"region"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

