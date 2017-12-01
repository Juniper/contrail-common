package db

import ("fmt"
        "testing"

        "github.com/Juniper/contrail-common/go/pkg/util"
        "github.com/Juniper/contrail-common/go/pkg/generated/models")

func TestE2ServiceProvider(t *testing.T) {
    t.Parallel()
    server, err := util.NewTestServer("TestE2ServiceProvider", tableDefs)
    if err != nil {
        t.Fatal(err)
    }
    defer server.Close()
    model := models.MakeE2ServiceProvider()
    db := server.DB
    tx, err := db.Begin()
    if err != nil {
        t.Fatal(err)
    }
    err = CreateE2ServiceProvider(tx, model)
    if err != nil {
        t.Fatal(err)
    }
    tx.Commit()

    tx2, err := db.Begin()
    if err != nil {
        t.Fatal(err)
    }
    models, err := ListE2ServiceProvider(tx2)
    if err != nil {
        t.Fatal(err)
    }
    tx2.Commit()
    fmt.Println(models)
}