package config

type ServerCfg struct {
	Addr string
	Env  string
	Db   DbCfg
}

type DbCfg struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}
