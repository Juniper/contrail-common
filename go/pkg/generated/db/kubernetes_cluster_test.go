package db

import ("fmt"
        "testing"

        "github.com/Juniper/contrail-common/go/pkg/util"
        "github.com/Juniper/contrail-common/go/pkg/generated/models")

func TestKubernetesCluster(t *testing.T) {
    t.Parallel()
    server, err := util.NewTestServer("TestKubernetesCluster", tableDefs)
    if err != nil {
        t.Fatal(err)
    }
    defer server.Close()
    model := models.MakeKubernetesCluster()
    db := server.DB
    tx, err := db.Begin()
    if err != nil {
        t.Fatal(err)
    }
    err = CreateKubernetesCluster(tx, model)
    if err != nil {
        t.Fatal(err)
    }
    tx.Commit()

    tx2, err := db.Begin()
    if err != nil {
        t.Fatal(err)
    }
    models, err := ListKubernetesCluster(tx2)
    if err != nil {
        t.Fatal(err)
    }
    tx2.Commit()
    fmt.Println(models)
}