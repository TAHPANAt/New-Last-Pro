
package main

import (
	"project.com/sa-68-project/configs"
//	"project.com/sa-68-project/controllers"
//	"github.com/gin-gonic/gin"
)

const PORT = "8088"

func main() {
	configs.ConnectionDB()
	configs.SetupDatabase()
}

