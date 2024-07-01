package configs

type Configs struct {
	MySQL MySQL
	App   Fiber
}

type Fiber struct {
	Host string
	Port string
}

type MySQL struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}
