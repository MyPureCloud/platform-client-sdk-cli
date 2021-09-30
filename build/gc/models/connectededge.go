package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnectededgeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnectededgeDud struct { 
    Id string `json:"id"`


    


    InterfaceName string `json:"interfaceName"`


    InterfaceIpAddress string `json:"interfaceIpAddress"`


    


    SelfUri string `json:"selfUri"`

}

// Connectededge
type Connectededge struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // EdgeConnectionList
    EdgeConnectionList []Edgeconnectioninfo `json:"edgeConnectionList"`


    

}

// String returns a JSON representation of the model
func (o *Connectededge) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.EdgeConnectionList = []Edgeconnectioninfo{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connectededge) MarshalJSON() ([]byte, error) {
    type Alias Connectededge

    if ConnectededgeMarshalled {
        return []byte("{}"), nil
    }
    ConnectededgeMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        EdgeConnectionList []Edgeconnectioninfo `json:"edgeConnectionList"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        EdgeConnectionList: []Edgeconnectioninfo{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

