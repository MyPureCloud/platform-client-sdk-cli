package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PosttextresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PosttextresponseDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Posttextresponse
type Posttextresponse struct { 
    // BotState - The state of the bot after completion of the request
    BotState string `json:"botState"`


    // ReplyMessages - The list of messages to respond with, if any
    ReplyMessages []Posttextmessage `json:"replyMessages"`


    // IntentName - The name of the intent the bot is either processing or has processed, this will be blank if no intent could be detected.
    IntentName string `json:"intentName"`


    // Slots - Data parameters detected and filled by the bot.
    Slots map[string]string `json:"slots"`


    // BotCorrelationId - The optional ID specified in the request
    BotCorrelationId string `json:"botCorrelationId"`


    // AmazonLex - Raw data response from AWS (if called)
    AmazonLex map[string]interface{} `json:"amazonLex"`


    // GoogleDialogFlow - Raw data response from Google Dialogflow (if called)
    GoogleDialogFlow map[string]interface{} `json:"googleDialogFlow"`


    // GenesysDialogEngine - Raw data response from Genesys' Dialogengine (if called)
    GenesysDialogEngine map[string]interface{} `json:"genesysDialogEngine"`


    // GenesysBotConnector - Raw data response from Genesys' BotConnector (if called)
    GenesysBotConnector map[string]interface{} `json:"genesysBotConnector"`


    // NuanceMixDlg - Raw data response from Nuance Mix Dlg (if called)
    NuanceMixDlg map[string]interface{} `json:"nuanceMixDlg"`

}

// String returns a JSON representation of the model
func (o *Posttextresponse) String() string {
    
    
    
    
    
    
     o.ReplyMessages = []Posttextmessage{{}} 
    
    
    
    
    
    
    
     o.Slots = map[string]string{"": ""} 
    
    
    
    
    
    
    
     o.AmazonLex = map[string]interface{}{"": Interface{}} 
    
    
    
     o.GoogleDialogFlow = map[string]interface{}{"": Interface{}} 
    
    
    
     o.GenesysDialogEngine = map[string]interface{}{"": Interface{}} 
    
    
    
     o.GenesysBotConnector = map[string]interface{}{"": Interface{}} 
    
    
    
     o.NuanceMixDlg = map[string]interface{}{"": Interface{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Posttextresponse) MarshalJSON() ([]byte, error) {
    type Alias Posttextresponse

    if PosttextresponseMarshalled {
        return []byte("{}"), nil
    }
    PosttextresponseMarshalled = true

    return json.Marshal(&struct { 
        BotState string `json:"botState"`
        
        ReplyMessages []Posttextmessage `json:"replyMessages"`
        
        IntentName string `json:"intentName"`
        
        Slots map[string]string `json:"slots"`
        
        BotCorrelationId string `json:"botCorrelationId"`
        
        AmazonLex map[string]interface{} `json:"amazonLex"`
        
        GoogleDialogFlow map[string]interface{} `json:"googleDialogFlow"`
        
        GenesysDialogEngine map[string]interface{} `json:"genesysDialogEngine"`
        
        GenesysBotConnector map[string]interface{} `json:"genesysBotConnector"`
        
        NuanceMixDlg map[string]interface{} `json:"nuanceMixDlg"`
        
        *Alias
    }{
        

        

        

        
        ReplyMessages: []Posttextmessage{{}},
        

        

        

        

        
        Slots: map[string]string{"": ""},
        

        

        

        

        
        AmazonLex: map[string]interface{}{"": Interface{}},
        

        

        
        GoogleDialogFlow: map[string]interface{}{"": Interface{}},
        

        

        
        GenesysDialogEngine: map[string]interface{}{"": Interface{}},
        

        

        
        GenesysBotConnector: map[string]interface{}{"": Interface{}},
        

        

        
        NuanceMixDlg: map[string]interface{}{"": Interface{}},
        

        
        Alias: (*Alias)(u),
    })
}

