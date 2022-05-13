package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotflowlaunchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotflowlaunchresponseDud struct { 
    

}

// Textbotflowlaunchresponse - Information related to a successful launch of a bot flow. The ID will be used in subsequent turn requests of the bot flow.
type Textbotflowlaunchresponse struct { 
    // Id - The session ID of the bot flow, used to send to subsequent turn requests
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Textbotflowlaunchresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotflowlaunchresponse) MarshalJSON() ([]byte, error) {
    type Alias Textbotflowlaunchresponse

    if TextbotflowlaunchresponseMarshalled {
        return []byte("{}"), nil
    }
    TextbotflowlaunchresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

