package errorf

import (
	"fmt"
	"os"
	"strings"
	"bytes"
)

type T struct{}

func (t *T) Errorf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	
	if err.Error() == "" {
		return
	}
	
	var message bytes.Buffer
	for _, line := range strings.Split(err.Error(), "\n") {
	    line = strings.TrimLeft(line, "\t")
	    if strings.HasPrefix(line, "Error:") || message.Len() > 0 {
		line = strings.ReplaceAll(line, "Error:", "\t")
		line = strings.TrimLeft(line, "            	")
		message.WriteString(line)
		message.WriteByte('\n')
	    }   
	}
 
	fmt.Fprint(os.Stderr, strings.TrimRight(message.String(), "\n")) 
}
