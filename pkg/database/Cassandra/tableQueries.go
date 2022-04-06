package cassandra

var tableQueries = [1]string
	`CREATE TABLE IF NOT EXISTS client_database.clients(id bnt, project_id bigint, is_paid_client boolean, created_at date, paid_time date, status boolean, PRIMARY KEY ((project_id),created_at)) WITH CLUSTERING ORDER BY (created_at DESC);`,
}

func GetTableQueries(1]string {
	return tableQueries
}
