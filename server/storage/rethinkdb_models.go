package storage

import (
	"github.com/docker/notary/storage/rethinkdb"
)

// These consts are the index names we've defined for RethinkDB
const (
	rdbSHA256Idx         = "sha256"
	rdbGunRoleIdx        = "gun_role"
	rdbGunRoleSHA256Idx  = "gun_role_sha256"
	rdbGunRoleVersionIdx = "gun_role_version"
)

var (
	// TUFFilesRethinkTable is the table definition of notary server's TUF metadata files
	TUFFilesRethinkTable = rethinkdb.Table{
		Name:       RDBTUFFile{}.TableName(),
		PrimaryKey: "gun_role_version",
		SecondaryIndexes: map[string][]string{
			rdbSHA256Idx:         nil,
			"gun":                nil,
			"timestamp_checksum": nil,
			rdbGunRoleIdx:        {"gun", "role"},
			rdbGunRoleSHA256Idx:  {"gun", "role", "sha256"},
		},
		// this configuration guarantees linearizability of individual atomic operations on individual documents
		Config: map[string]string{
			"write_acks": "majority",
		},
		JSONUnmarshaller: rdbTUFFileFromJSON,
	}

	// ChangeRethinkTable is the table definition for changefeed objects
	ChangeRethinkTable = rethinkdb.Table{
		Name:       Change{}.TableName(),
		PrimaryKey: "id",
		SecondaryIndexes: map[string][]string{
			"rdb_created_at_id":     {"created_at", "id"},
			"rdb_gun_created_at_id": {"gun", "created_at", "id"},
		},
		Config: map[string]string{
			"write_acks": "majority",
		},
		JSONUnmarshaller: rdbChangeFromJSON,
	}
)
