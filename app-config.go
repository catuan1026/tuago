package tuago

type AppConfiger struct {
	ActiveEnv     string            `yaml:"activeEnv"` //当前激活的环境
	Databases     []DatabaseInfo    `yaml:"databases"`
	Redis         []RedisInfo       `yaml:"redis"`
	Routes        []RouteInfo       `yaml:"routes"`
	EnvProperties map[string]string `yaml:"envProperties"`
}

type RouteInfo struct {
	Path    string   `yaml:"path"`
	Methods []string `yaml:"methods"`
	Timeout int      `yaml:"timeout"` // 单位秒
	Role    string   `yaml:"role"`    // 角色
}

type LogInfo struct {
	Level    string `yaml:"level"`
	FilePath string `yaml:"file_path"`
}

type DatabaseInfo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
	Debug    int8   `yaml:"debug"`
}

type RedisInfo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}
