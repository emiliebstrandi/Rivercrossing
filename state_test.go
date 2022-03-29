package state

import (
	"testing"
) 
func TestViewState(t *testing.T) {  
    wanted := "[kylling rev korn ---\\__/ _________________/---]"     
    state := CheckState();     
    if state != wanted {    
        t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)           
    }    
}  