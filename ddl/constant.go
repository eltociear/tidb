// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ddl

import (
	"github.com/pingcap/tidb/meta"
)

const (
	// JobTable stores the information of DDL jobs.
	JobTable = "tidb_ddl_job"
	// ReorgTable stores the information of DDL reorganization.
	ReorgTable = "tidb_ddl_reorg"
	// HistoryTable stores the history DDL jobs.
	HistoryTable = "tidb_ddl_history"

	// JobTableID is the table ID of `tidb_ddl_job`.
	JobTableID = meta.MaxInt48 - 1
	// ReorgTableID is the table ID of `tidb_ddl_reorg`.
	ReorgTableID = meta.MaxInt48 - 2
	// HistoryTableID is the table ID of `tidb_ddl_history`.
	HistoryTableID = meta.MaxInt48 - 3

	// JobTableSQL is the CREATE TABLE SQL of `tidb_ddl_job`.
	JobTableSQL = "create table " + JobTable + "(job_id bigint not null, reorg int, schema_ids text(65535), table_ids text(65535), job_meta longblob, type int, processing int, primary key(job_id))"
	// ReorgTableSQL is the CREATE TABLE SQL of `tidb_ddl_reorg`.
	ReorgTableSQL = "create table " + ReorgTable + "(job_id bigint not null, ele_id bigint, ele_type blob, start_key blob, end_key blob, physical_id bigint, reorg_meta longblob, unique key(job_id, ele_id, ele_type(20)))"
	// HistoryTableSQL is the CREATE TABLE SQL of `tidb_ddl_history`.
	HistoryTableSQL = "create table " + HistoryTable + "(job_id bigint not null, job_meta longblob, db_name char(64), table_name char(64), schema_ids text(65535), table_ids text(65535), create_time datetime, primary key(job_id))"
)
