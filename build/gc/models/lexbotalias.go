package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LexbotaliasMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LexbotaliasDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Lexbotalias
type Lexbotalias struct { 
    


    // Name
    Name string `json:"name"`


    // Bot - The Lex bot this is an alias for
    Bot Lexbot `json:"bot"`


    // BotVersion - The version of the Lex bot this alias points at
    BotVersion string `json:"botVersion"`


    // Status - The status of the Lex bot alias
    Status string `json:"status"`


    // FailureReason - If the status is FAILED, Amazon Lex explains why it failed to build the bot
    FailureReason string `json:"failureReason"`


    // Language - The target language of the Lex bot
    Language string `json:"language"`


    // Intents - An array of Intents associated with this bot alias
    Intents []Lexintent `json:"intents"`


    

}

// String returns a JSON representation of the model
func (o *Lexbotalias) String() string {
    
    
    
    
    
    
     o.Intents = []Lexintent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexbotalias) MarshalJSON() ([]byte, error) {
    type Alias Lexbotalias

    if LexbotaliasMarshalled {
        return []byte("{}"), nil
    }
    LexbotaliasMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Bot Lexbot `json:"bot"`
        
        BotVersion string `json:"botVersion"`
        
        Status string `json:"status"`
        
        FailureReason string `json:"failureReason"`
        
        Language string `json:"language"`
        
        Intents []Lexintent `json:"intents"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Intents: []Lexintent{{}},
        


        

        Alias: (*Alias)(u),
    })
}

