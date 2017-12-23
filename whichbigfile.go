package whichbigfile

import (
    "fmt"
    "os/exec"
    )
    
func WhichBigFile(directory string) {
      lsCmd := exec.Command("bash", "-c", "find "+directory+" -type f -printf '%s %p\n' | sort -rn | head -n 5")
      lsOut, _:= lsCmd.Output()
      fmt.Println("Listing.. \n\nSize | File\n")
      fmt.Println(string(lsOut))
}
