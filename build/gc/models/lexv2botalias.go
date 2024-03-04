package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Lexv2botaliasMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Lexv2botaliasDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Lexv2botalias
type Lexv2botalias struct { 
    


    // Name
    Name string `json:"name"`


    // Region - The Lex V2 bot region
    Region string `json:"region"`


    // AliasId - The Lex V2 bot alias Id
    AliasId string `json:"aliasId"`


    // Bot - The Lex V2 bot this is an alias for
    Bot Lexv2bot `json:"bot"`


    // BotVersion - The version of the Lex V2 bot this alias points at
    BotVersion string `json:"botVersion"`


    // Status - The status of the Lex V2 bot alias
    Status string `json:"status"`


    // Language - The target language of the Lex V2 bot
    Language string `json:"language"`


    // Intents - An array of Intents associated with this bot alias
    Intents []Lexv2intent `json:"intents"`


    

}

// String returns a JSON representation of the model
func (o *Lexv2botalias) String() string {
    
    
    
    
    
    
    
     o.Intents = []Lexv2intent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexv2botalias) MarshalJSON() ([]byte, error) {
    type Alias Lexv2botalias

    if Lexv2botaliasMarshalled {
        return []byte("{}"), nil
    }
    Lexv2botaliasMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Region string `json:"region"`
        
        AliasId string `json:"aliasId"`
        
        Bot Lexv2bot `json:"bot"`
        
        BotVersion string `json:"botVersion"`
        
        Status string `json:"status"`
        
        Language string `json:"language"`
        
        Intents []Lexv2intent `json:"intents"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Intents: []Lexv2intent{{}},
        


        

        Alias: (*Alias)(u),
    })
}

