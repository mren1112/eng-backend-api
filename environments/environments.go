package environments

import (

	"time"
 
)



func TimeZoneInit() {
	ict := time.Now().Local().Location()
	time.Local = ict
}