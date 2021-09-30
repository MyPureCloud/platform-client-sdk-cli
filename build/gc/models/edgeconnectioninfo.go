package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeconnectioninfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeconnectioninfoDud struct { 
    Id string `json:"id"`


    


    InterfaceName string `json:"interfaceName"`


    InterfaceIpAddress string `json:"interfaceIpAddress"`


    ConnectionErrors []string `json:"connectionErrors"`


    


    SelfUri string `json:"selfUri"`

}

// Edgeconnectioninfo
type Edgeconnectioninfo struct { 
    


    // Name
    Name string `json:"name"`


    


    


    


    // Site
    Site Addressableentityref `json:"site"`


    

}

// String returns a JSON representation of the model
func (o *Edgeconnectioninfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeconnectioninfo) MarshalJSON() ([]byte, error) {
    type Alias Edgeconnectioninfo

    if EdgeconnectioninfoMarshalled {
        return []byte("{}"), nil
    }
    EdgeconnectioninfoMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        
        
        Site Addressableentityref `json:"site"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

