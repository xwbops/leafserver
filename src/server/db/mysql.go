package db

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"server/db/ent"
	"time"
)

func NewEntClient(url string) *ent.Client {
	client, err := ent.Open(
		DBDriverName,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	if conf.Database.IsDebug {
		client = client.Debug()
	}

	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Add a global hook that runs on all types and all operations.
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				log.Infof("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})

	return client
}
