package migrate

import (
	"fmt"
	"os"

	"github.com/grcatterall/go-web-server/pkg/utils"
)

func Migrate() {
	content, err := os.ReadFile("../migrations/20240529105239.sql")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	conn := utils.DbConnection()

	res, err := conn.Query(string(content))

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
