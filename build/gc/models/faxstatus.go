package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FaxstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FaxstatusDud struct { 
    


    


    


    


    


    


    


    

}

// Faxstatus
type Faxstatus struct { 
    // Direction - The fax direction, either \"send\" or \"receive\".
    Direction string `json:"direction"`


    // ExpectedPages - Total number of expected pages, if known.
    ExpectedPages int `json:"expectedPages"`


    // ActivePage - Active page of the transmission.
    ActivePage int `json:"activePage"`


    // LinesTransmitted - Number of lines that have completed transmission.
    LinesTransmitted int `json:"linesTransmitted"`


    // BytesTransmitted - Number of bytes that have competed transmission.
    BytesTransmitted int `json:"bytesTransmitted"`


    // BaudRate - Current signaling rate of transmission, baud rate.
    BaudRate int `json:"baudRate"`


    // PageErrors - Number of page errors.
    PageErrors int `json:"pageErrors"`


    // LineErrors - Number of line errors.
    LineErrors int `json:"lineErrors"`

}

// String returns a JSON representation of the model
func (o *Faxstatus) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Faxstatus) MarshalJSON() ([]byte, error) {
    type Alias Faxstatus

    if FaxstatusMarshalled {
        return []byte("{}"), nil
    }
    FaxstatusMarshalled = true

    return json.Marshal(&struct { 
        Direction string `json:"direction"`
        
        ExpectedPages int `json:"expectedPages"`
        
        ActivePage int `json:"activePage"`
        
        LinesTransmitted int `json:"linesTransmitted"`
        
        BytesTransmitted int `json:"bytesTransmitted"`
        
        BaudRate int `json:"baudRate"`
        
        PageErrors int `json:"pageErrors"`
        
        LineErrors int `json:"lineErrors"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

