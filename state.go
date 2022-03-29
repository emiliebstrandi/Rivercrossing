package state

import (
	"testing"
)

var state = "[kylling rev korn ---\\ \\__E__/ _________________/---]"
  

func PutInBoat(){
	state = "[kylling rev korn ---\\ \\__E__/ _________________/---]"    
}

func CrossRiver(){
	state = "[kylling rev korn ---\\ \\ _________________\\__E__/---]"
}

func CheckState() string{
	return state
}  