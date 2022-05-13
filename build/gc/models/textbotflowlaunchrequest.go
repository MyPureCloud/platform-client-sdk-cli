package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotflowlaunchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotflowlaunchrequestDud struct { 
    


    


    


    


    

}

// Textbotflowlaunchrequest - Settings for launching an instance of a bot flow.
type Textbotflowlaunchrequest struct { 
    // Flow - Specifies which Bot Flow to launch.
    Flow Textbotflow `json:"flow"`


    // ExternalSessionId - The ID of the external session that is associated with the bot flow.
    ExternalSessionId string `json:"externalSessionId"`


    // ConversationId - A conversation ID to associate with the bot flow, if available.
    ConversationId string `json:"conversationId"`


    // InputData - Input values to the flow. Valid values are defined by the flow's input JSON schema.
    InputData Textbotinputoutputdata `json:"inputData"`


    // Channel - Channel information relevant to the bot flow.
    Channel Textbotchannel `json:"channel"`

}

// String returns a JSON representation of the model
func (o *Textbotflowlaunchrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotflowlaunchrequest) MarshalJSON() ([]byte, error) {
    type Alias Textbotflowlaunchrequest

    if TextbotflowlaunchrequestMarshalled {
        return []byte("{}"), nil
    }
    TextbotflowlaunchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Flow Textbotflow `json:"flow"`
        
        ExternalSessionId string `json:"externalSessionId"`
        
        ConversationId string `json:"conversationId"`
        
        InputData Textbotinputoutputdata `json:"inputData"`
        
        Channel Textbotchannel `json:"channel"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

