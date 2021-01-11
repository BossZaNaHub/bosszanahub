package bzgorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host string
	Username string
	Password string
	DB string
	Port string
	SslMode bool
	Timezone string
}

type defaultClient struct {
	Config Config
	DB *gorm.DB
}

func NewClient(config Config) Client {
	return &defaultClient{
		Config: config,
	}
}

func (c *defaultClient) InitMysql() (*gorm.DB, error) {
	dsn := c.Config.Username + ":" + c.Config.Password + "@tcp(" + c.Config.Host + ":" + defaultPort("MYSQL",c.Config.Port) +")/"+ c.Config.DB + "?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *defaultClient) InitPsql() (*gorm.DB, error) {
	dsn := "host=" + c.Config.Host +" user=" + c.Config.Username + " password=" + c.Config.Password + "DB.name=" + c.Config.DB + "port=" + defaultPort("POSTGRESQL",c.Config.Port) + "sslmode=" + openSSL(c.Config.SslMode) +  "TimeZone=" + defaultTimeZone(c.Config.Timezone)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func openSSL(mode bool) string {
	if mode {
		return "enable"
	}
	return "disable"
}

func defaultTimeZone(zone string) string {
	if zone == "" {
		return "Asia/Bangkok"
	}

	return zone
}

func defaultPort(driverName, port string) string {
	if port != "" {
		return port
	}

	switch driverName {
	case "MYSQL":
		return "3306"
	case "POSTGRESQL":
		return "5432"
	default:
		return ""
	}
}