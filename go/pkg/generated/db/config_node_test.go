package db

import (
	"fmt"
	"testing"

	"github.com/Juniper/contrail-common/go/pkg/generated/models"
	"github.com/Juniper/contrail-common/go/pkg/util"
)

func TestConfigNode(t *testing.T) {
	t.Parallel()
	server, err := util.NewTestServer("TestConfigNode", tableDefs)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	model := models.MakeConfigNode()
	model.UUID = "dummy_uuid"
	db := server.DB
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = CreateConfigNode(tx, model)
	if err != nil {
		t.Fatal(err)
	}
	tx.Commit()

	tx2, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	models, err := ListConfigNode(tx2, nil, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	tx2.Commit()
	if len(models) != 1 {
		t.Fatal("List failed")
	}

	tx3, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	model2, err := ShowConfigNode(tx3, model.UUID)
	if err != nil {
		t.Fatal(err)
	}
	tx3.Commit()
	if model2 == nil {
		t.Fatal("show failed")
	}

	tx4, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DeleteConfigNode(tx4, model.UUID)
	if err != nil {
		t.Fatal(err)
	}
	tx4.Commit()
	if model2 == nil {
		t.Fatal("delete failed")
	}

	tx5, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	models, err = ListConfigNode(tx5, nil, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	tx5.Commit()
	if len(models) != 0 {
		t.Fatal("delete failed")
	}
	fmt.Println(models)
}
