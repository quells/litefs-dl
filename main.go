package main

import (
	"context"
	"log"
	"net/url"
	"os"

	"github.com/superfly/litefs"
	"github.com/superfly/litefs/lfsc"
	"github.com/superfly/ltx"
)

func main() {
	dbName := os.Args[1]
	dbFileName := os.Args[2]

	store := litefs.NewStore("/tmp/litefs", false)
	store.Leaser = litefs.NewStaticLeaser(false, "", "")
	if err := store.Open(); err != nil {
		log.Println("store.Open", err)
		return
	}
	defer store.Close()

	u, _ := url.Parse("https://litefs.fly.io")
	client := lfsc.NewBackupClient(store, *u)
	client.AuthToken = os.Getenv("LITEFS_CLOUD_TOKEN")
	if err := client.Open(); err != nil {
		log.Println("lfsc.NewBackupClient", err)
		return
	}

	ctx := context.Background()

	snapshot, err := client.FetchSnapshot(ctx, dbName)
	if err != nil {
		log.Println("client.FetchSnapshot", err)
		return
	}
	defer snapshot.Close()

	dec := ltx.NewDecoder(snapshot)
	defer dec.Close()

	_ = os.Remove(dbFileName)
	f, err := os.Create(dbFileName)
	if err != nil {
		log.Println("os.Open", err)
		return
	}
	defer f.Close()

	if err := dec.DecodeDatabaseTo(f); err != nil {
		log.Println("dec.DecodeDatabaseTo", err)
		return
	}
}
