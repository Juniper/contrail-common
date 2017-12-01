package db

import ("fmt"
        "testing"

        "github.com/Juniper/contrail-common/go/pkg/util"
        "github.com/Juniper/contrail-common/go/pkg/generated/models")

func TestQosConfig(t *testing.T) {
    t.Parallel()
    server, err := util.NewTestServer("TestQosConfig", tableDefs)
    if err != nil {
        t.Fatal(err)
    }
    defer server.Close()
    model := models.MakeQosConfig()
    db := server.DB
    tx, err := db.Begin()
    if err != nil {
        t.Fatal(err)
    }
    err = CreateQosConfig(tx, model)
    if err != nil {
        t.Fatal(err)
    }
    tx.Commit()

    tx2, err := db.Begin()
    if err != nil {
        t.Fatal(err)
    }
    models, err := ListQosConfig(tx2)
    if err != nil {
        t.Fatal(err)
    }
    tx2.Commit()
    fmt.Println(models)
}