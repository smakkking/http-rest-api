package sqlstore_test

import (
	"fmt"
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	// вызывается один раз во всем пакете

	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		fmt.Print("dfsgfdsh")
		databaseUrl = "host=localhost dbname=restapi_test user=postgres password=myPassword sslmode=disable"
	}

	os.Exit(m.Run())
}
