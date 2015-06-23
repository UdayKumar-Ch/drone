package builtin

// DO NOT EDIT
// code generated by go:generate

import (
	"database/sql"
	"encoding/json"

	. "github.com/drone/drone/pkg/types"
)

var _ = json.Marshal

// generic database interface, matching both *sql.Db and *sql.Tx
type buildDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func getBuild(db buildDB, query string, args ...interface{}) (*Build, error) {
	row := db.QueryRow(query, args...)
	return scanBuild(row)
}

func getBuilds(db buildDB, query string, args ...interface{}) ([]*Build, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanBuilds(rows)
}

func createBuild(db buildDB, query string, v *Build) error {
	var v0 int64
	var v1 int
	var v2 string
	var v3 int64
	var v4 int64
	var v5 string
	var v6 string
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 int
	var v14 string
	var v15 string
	var v16 string
	var v17 string
	var v18 string
	var v19 string
	var v20 string
	var v21 string
	var v22 string
	v0 = v.RepoID
	v1 = v.Number
	v2 = v.Status
	v3 = v.Started
	v4 = v.Finished
	if v.Commit != nil {
		v5 = v.Commit.Sha
		v6 = v.Commit.Ref
		v7 = v.Commit.Branch
		v8 = v.Commit.Message
		v9 = v.Commit.Timestamp
		v10 = v.Commit.Remote
		if v.Commit.Author != nil {
			v11 = v.Commit.Author.Login
			v12 = v.Commit.Author.Email
		}
	}
	if v.PullRequest != nil {
		v13 = v.PullRequest.Number
		v14 = v.PullRequest.Title
		if v.PullRequest.Base != nil {
			v15 = v.PullRequest.Base.Sha
			v16 = v.PullRequest.Base.Ref
			v17 = v.PullRequest.Base.Branch
			v18 = v.PullRequest.Base.Message
			v19 = v.PullRequest.Base.Timestamp
			v20 = v.PullRequest.Base.Remote
			if v.PullRequest.Base.Author != nil {
				v21 = v.PullRequest.Base.Author.Login
				v22 = v.PullRequest.Base.Author.Email
			}
		}
	}

	res, err := db.Exec(query,
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
		&v9,
		&v10,
		&v11,
		&v12,
		&v13,
		&v14,
		&v15,
		&v16,
		&v17,
		&v18,
		&v19,
		&v20,
		&v21,
		&v22,
	)
	if err != nil {
		return err
	}

	v.ID, err = res.LastInsertId()
	return err
}

func updateBuild(db buildDB, query string, v *Build) error {
	var v0 int64
	var v1 int64
	var v2 int
	var v3 string
	var v4 int64
	var v5 int64
	var v6 string
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 int
	var v15 string
	var v16 string
	var v17 string
	var v18 string
	var v19 string
	var v20 string
	var v21 string
	var v22 string
	var v23 string
	v0 = v.ID
	v1 = v.RepoID
	v2 = v.Number
	v3 = v.Status
	v4 = v.Started
	v5 = v.Finished
	if v.Commit != nil {
		v6 = v.Commit.Sha
		v7 = v.Commit.Ref
		v8 = v.Commit.Branch
		v9 = v.Commit.Message
		v10 = v.Commit.Timestamp
		v11 = v.Commit.Remote
		if v.Commit.Author != nil {
			v12 = v.Commit.Author.Login
			v13 = v.Commit.Author.Email
		}
	}
	if v.PullRequest != nil {
		v14 = v.PullRequest.Number
		v15 = v.PullRequest.Title
		if v.PullRequest.Base != nil {
			v16 = v.PullRequest.Base.Sha
			v17 = v.PullRequest.Base.Ref
			v18 = v.PullRequest.Base.Branch
			v19 = v.PullRequest.Base.Message
			v20 = v.PullRequest.Base.Timestamp
			v21 = v.PullRequest.Base.Remote
			if v.PullRequest.Base.Author != nil {
				v22 = v.PullRequest.Base.Author.Login
				v23 = v.PullRequest.Base.Author.Email
			}
		}
	}

	_, err := db.Exec(query,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
		&v9,
		&v10,
		&v11,
		&v12,
		&v13,
		&v14,
		&v15,
		&v16,
		&v17,
		&v18,
		&v19,
		&v20,
		&v21,
		&v22,
		&v23,
		&v0,
	)
	return err
}

func scanBuild(row *sql.Row) (*Build, error) {
	var v0 int64
	var v1 int64
	var v2 int
	var v3 string
	var v4 int64
	var v5 int64
	var v6 string
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 int
	var v15 string
	var v16 string
	var v17 string
	var v18 string
	var v19 string
	var v20 string
	var v21 string
	var v22 string
	var v23 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
		&v9,
		&v10,
		&v11,
		&v12,
		&v13,
		&v14,
		&v15,
		&v16,
		&v17,
		&v18,
		&v19,
		&v20,
		&v21,
		&v22,
		&v23,
	)
	if err != nil {
		return nil, err
	}

	v := &Build{}
	v.ID = v0
	v.RepoID = v1
	v.Number = v2
	v.Status = v3
	v.Started = v4
	v.Finished = v5
	v.Commit = &Commit{}
	v.Commit.Sha = v6
	v.Commit.Ref = v7
	v.Commit.Branch = v8
	v.Commit.Message = v9
	v.Commit.Timestamp = v10
	v.Commit.Remote = v11
	v.Commit.Author = &Author{}
	v.Commit.Author.Login = v12
	v.Commit.Author.Email = v13
	v.PullRequest = &PullRequest{}
	v.PullRequest.Number = v14
	v.PullRequest.Title = v15
	v.PullRequest.Base = &Commit{}
	v.PullRequest.Base.Sha = v16
	v.PullRequest.Base.Ref = v17
	v.PullRequest.Base.Branch = v18
	v.PullRequest.Base.Message = v19
	v.PullRequest.Base.Timestamp = v20
	v.PullRequest.Base.Remote = v21
	v.PullRequest.Base.Author = &Author{}
	v.PullRequest.Base.Author.Login = v22
	v.PullRequest.Base.Author.Email = v23

	return v, nil
}

func scanBuilds(rows *sql.Rows) ([]*Build, error) {
	var err error
	var vv []*Build
	for rows.Next() {
		var v0 int64
		var v1 int64
		var v2 int
		var v3 string
		var v4 int64
		var v5 int64
		var v6 string
		var v7 string
		var v8 string
		var v9 string
		var v10 string
		var v11 string
		var v12 string
		var v13 string
		var v14 int
		var v15 string
		var v16 string
		var v17 string
		var v18 string
		var v19 string
		var v20 string
		var v21 string
		var v22 string
		var v23 string
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,
			&v7,
			&v8,
			&v9,
			&v10,
			&v11,
			&v12,
			&v13,
			&v14,
			&v15,
			&v16,
			&v17,
			&v18,
			&v19,
			&v20,
			&v21,
			&v22,
			&v23,
		)
		if err != nil {
			return vv, err
		}

		v := &Build{}
		v.ID = v0
		v.RepoID = v1
		v.Number = v2
		v.Status = v3
		v.Started = v4
		v.Finished = v5
		v.Commit = &Commit{}
		v.Commit.Sha = v6
		v.Commit.Ref = v7
		v.Commit.Branch = v8
		v.Commit.Message = v9
		v.Commit.Timestamp = v10
		v.Commit.Remote = v11
		v.Commit.Author = &Author{}
		v.Commit.Author.Login = v12
		v.Commit.Author.Email = v13
		v.PullRequest = &PullRequest{}
		v.PullRequest.Number = v14
		v.PullRequest.Title = v15
		v.PullRequest.Base = &Commit{}
		v.PullRequest.Base.Sha = v16
		v.PullRequest.Base.Ref = v17
		v.PullRequest.Base.Branch = v18
		v.PullRequest.Base.Message = v19
		v.PullRequest.Base.Timestamp = v20
		v.PullRequest.Base.Remote = v21
		v.PullRequest.Base.Author = &Author{}
		v.PullRequest.Base.Author.Login = v22
		v.PullRequest.Base.Author.Email = v23
		vv = append(vv, v)
	}
	return vv, rows.Err()
}

const stmtBuildSelectList = `
SELECT
 build_id
,build_repo_id
,build_number
,build_status
,build_started
,build_finished
,build_commit_sha
,build_commit_ref
,build_commit_branch
,build_commit_message
,build_commit_timestamp
,build_commit_remote
,build_commit_author_login
,build_commit_author_email
,build_pull_request_number
,build_pull_request_title
,build_pull_request_base_sha
,build_pull_request_base_ref
,build_pull_request_base_branch
,build_pull_request_base_message
,build_pull_request_base_timestamp
,build_pull_request_base_remote
,build_pull_request_base_author_login
,build_pull_request_base_author_email
FROM builds
`

const stmtBuildSelectRange = `
SELECT
 build_id
,build_repo_id
,build_number
,build_status
,build_started
,build_finished
,build_commit_sha
,build_commit_ref
,build_commit_branch
,build_commit_message
,build_commit_timestamp
,build_commit_remote
,build_commit_author_login
,build_commit_author_email
,build_pull_request_number
,build_pull_request_title
,build_pull_request_base_sha
,build_pull_request_base_ref
,build_pull_request_base_branch
,build_pull_request_base_message
,build_pull_request_base_timestamp
,build_pull_request_base_remote
,build_pull_request_base_author_login
,build_pull_request_base_author_email
FROM builds
LIMIT ? OFFSET ?
`

const stmtBuildSelect = `
SELECT
 build_id
,build_repo_id
,build_number
,build_status
,build_started
,build_finished
,build_commit_sha
,build_commit_ref
,build_commit_branch
,build_commit_message
,build_commit_timestamp
,build_commit_remote
,build_commit_author_login
,build_commit_author_email
,build_pull_request_number
,build_pull_request_title
,build_pull_request_base_sha
,build_pull_request_base_ref
,build_pull_request_base_branch
,build_pull_request_base_message
,build_pull_request_base_timestamp
,build_pull_request_base_remote
,build_pull_request_base_author_login
,build_pull_request_base_author_email
FROM builds
WHERE build_id = ?
`

const stmtBuildSelectBuildRepoId = `
SELECT
 build_id
,build_repo_id
,build_number
,build_status
,build_started
,build_finished
,build_commit_sha
,build_commit_ref
,build_commit_branch
,build_commit_message
,build_commit_timestamp
,build_commit_remote
,build_commit_author_login
,build_commit_author_email
,build_pull_request_number
,build_pull_request_title
,build_pull_request_base_sha
,build_pull_request_base_ref
,build_pull_request_base_branch
,build_pull_request_base_message
,build_pull_request_base_timestamp
,build_pull_request_base_remote
,build_pull_request_base_author_login
,build_pull_request_base_author_email
FROM builds
WHERE build_repo_id = ?
`

const stmtBuildSelectBuildNumber = `
SELECT
 build_id
,build_repo_id
,build_number
,build_status
,build_started
,build_finished
,build_commit_sha
,build_commit_ref
,build_commit_branch
,build_commit_message
,build_commit_timestamp
,build_commit_remote
,build_commit_author_login
,build_commit_author_email
,build_pull_request_number
,build_pull_request_title
,build_pull_request_base_sha
,build_pull_request_base_ref
,build_pull_request_base_branch
,build_pull_request_base_message
,build_pull_request_base_timestamp
,build_pull_request_base_remote
,build_pull_request_base_author_login
,build_pull_request_base_author_email
FROM builds
WHERE build_repo_id = ?
AND build_number = ?
`

const stmtBuildSelectCommitBranch = `
SELECT
 build_id
,build_repo_id
,build_number
,build_status
,build_started
,build_finished
,build_commit_sha
,build_commit_ref
,build_commit_branch
,build_commit_message
,build_commit_timestamp
,build_commit_remote
,build_commit_author_login
,build_commit_author_email
,build_pull_request_number
,build_pull_request_title
,build_pull_request_base_sha
,build_pull_request_base_ref
,build_pull_request_base_branch
,build_pull_request_base_message
,build_pull_request_base_timestamp
,build_pull_request_base_remote
,build_pull_request_base_author_login
,build_pull_request_base_author_email
FROM builds
WHERE build_branch = ?
AND build_branch = ?
`

const stmtBuildSelectCount = `
SELECT count(1)
FROM builds
`

const stmtBuildInsert = `
INSERT INTO builds (
 build_repo_id
,build_number
,build_status
,build_started
,build_finished
,build_commit_sha
,build_commit_ref
,build_commit_branch
,build_commit_message
,build_commit_timestamp
,build_commit_remote
,build_commit_author_login
,build_commit_author_email
,build_pull_request_number
,build_pull_request_title
,build_pull_request_base_sha
,build_pull_request_base_ref
,build_pull_request_base_branch
,build_pull_request_base_message
,build_pull_request_base_timestamp
,build_pull_request_base_remote
,build_pull_request_base_author_login
,build_pull_request_base_author_email
) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);
`

const stmtBuildUpdate = `
UPDATE builds SET
 build_repo_id = ?
,build_number = ?
,build_status = ?
,build_started = ?
,build_finished = ?
,build_commit_sha = ?
,build_commit_ref = ?
,build_commit_branch = ?
,build_commit_message = ?
,build_commit_timestamp = ?
,build_commit_remote = ?
,build_commit_author_login = ?
,build_commit_author_email = ?
,build_pull_request_number = ?
,build_pull_request_title = ?
,build_pull_request_base_sha = ?
,build_pull_request_base_ref = ?
,build_pull_request_base_branch = ?
,build_pull_request_base_message = ?
,build_pull_request_base_timestamp = ?
,build_pull_request_base_remote = ?
,build_pull_request_base_author_login = ?
,build_pull_request_base_author_email = ?
WHERE build_id = ?
`

const stmtBuildDelete = `
DELETE FROM builds
WHERE build_id = ?
`

const stmtBuildTable = `
CREATE TABLE IF NOT EXISTS builds (
 build_id                             INTEGER PRIMARY KEY AUTOINCREMENT
,build_repo_id                        INTEGER
,build_number                         INTEGER
,build_status                         VARCHAR(512)
,build_started                        INTEGER
,build_finished                       INTEGER
,build_commit_sha                     VARCHAR(512)
,build_commit_ref                     VARCHAR(512)
,build_commit_branch                  VARCHAR(512)
,build_commit_message                 VARCHAR(512)
,build_commit_timestamp               VARCHAR(512)
,build_commit_remote                  VARCHAR(512)
,build_commit_author_login            VARCHAR(512)
,build_commit_author_email            VARCHAR(512)
,build_pull_request_number            INTEGER
,build_pull_request_title             VARCHAR(512)
,build_pull_request_base_sha          VARCHAR(512)
,build_pull_request_base_ref          VARCHAR(512)
,build_pull_request_base_branch       VARCHAR(512)
,build_pull_request_base_message      VARCHAR(512)
,build_pull_request_base_timestamp    VARCHAR(512)
,build_pull_request_base_remote       VARCHAR(512)
,build_pull_request_base_author_login VARCHAR(512)
,build_pull_request_base_author_email VARCHAR(512)
);
`

const stmtBuildBuildRepoIdIndex = `
CREATE INDEX IF NOT EXISTS ix_build_repo_id ON builds (build_repo_id);
`

const stmtBuildBuildNumberIndex = `
CREATE UNIQUE INDEX IF NOT EXISTS ux_build_number ON builds (build_repo_id,build_number);
`

const stmtBuildCommitBranchIndex = `
CREATE INDEX IF NOT EXISTS ix_commit_branch ON builds (build_branch,build_branch);
`
