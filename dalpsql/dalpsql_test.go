package dalpsql_test

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/broersa/mybroker/dalpsql"
)

func TestGetFreeNwkAddr(t *testing.T) {
	c := os.Getenv("MYBROKER_DB")
	s, err := sql.Open("postgres", c)
	d := dalpsql.New(s)
	a, err := d.GetFreeNwkAddr()
	if err != nil {
		t.Error(err)
	}
	if a == nil {
		t.Error("There must be a free nwkaddr")
	}
}
