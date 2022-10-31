package pgconfig

type PostgreSQLConfig struct {
	DataSourceName               string
	MaxOpenConnections           int
	MaxIdleConnections           int
	MaxConnectionLifetimeMinutes int
}
