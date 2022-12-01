package mysql

type MysqlConfig struct {
	Username     string
	Password     string
	DatabaseName string
	Host         string
	Port         string
}

func (c *MysqlConfig) AsDSN() string {
	return c.Username + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DatabaseName + "?parseTime=true"
}
