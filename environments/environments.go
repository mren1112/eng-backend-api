package environments

import (
<<<<<<< HEAD

=======
>>>>>>> d59ef4c7f3f3949a5d269274d97355325461518d
	"strings"
	"time"

	"github.com/spf13/viper"
<<<<<<< HEAD
 
)


func EnvironmentInit() {
	// ไฟล์ที่จะจัดเก็บตัว Connection string Database
	//viper.SetConfigName("config")
	// ภาษาที่จะใช้ในการ Config
	//viper.SetConfigType("yaml")
=======
)

func EnvironmentInit() {
	// ไฟล์ที่จะจัดเก็บตัว Connection string Database
	viper.SetConfigName("config")
	// ภาษาที่จะใช้ในการ Config
	viper.SetConfigType("yaml")
>>>>>>> d59ef4c7f3f3949a5d269274d97355325461518d
	// แล้วเข้ามาที่ environment folder
	viper.AddConfigPath("./environments")
	// ที่อยู่ของ file config เริ่มค้นหาจาก root ด้านนอกสุด
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// เรียก file config.yaml มาใช้
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func TimeZoneInit() {
	ict := time.Now().Local().Location()
	time.Local = ict
}