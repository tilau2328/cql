package ddl

import (
	. "github.com/samber/lo"
	. "github.com/tilau2328/cql/package/domain/model"
	. "github.com/tilau2328/cql/package/shared/cmd/pretty"
	"os"
)

func PrintKeySpaces(ks []KeySpace) {
	NewTable(
		Header("#", "KeyspaceName", "DurableWrites", "Replication"),
		Rows(Map(ks, func(v KeySpace, i int) []any {
			return []any{i, v.KeySpaceKey, v.Durable, v.Replication}
		})...),
	).Write(os.Stdout)
}

func PrintTables(t []Table) {
	NewTable(
		Header("#", "Id", "KeyspaceName", "TableName", "Comment"),
		Rows(Map(t, func(v Table, i int) []any {
			return []any{i, v.Id, v.KeySpace, v.Name, v.Comment}
		})...),
	).Write(os.Stdout)
}
