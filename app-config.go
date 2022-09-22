package tuago

type AppConfiger struct {
	Services  []ServiceInfo     `yaml:"services" json:"services"`
	Databases []DatabaseInfo    `yaml:"databases" json:"databases"`
	Redis     []RedisInfo       `yaml:"redis,omitempty" json:"redis,omitempty"`
	Routes    []RouteInfo       `yaml:"routes,omitempty" json:"routes,omitempty"`
	Session   *SessionInfo      `yaml:"session,omitempty" json:"session,omitempty"`
	Env       map[string]string `yaml:"env,omitempty" json:"env,omitempty"`
}

type RouteInfo struct {
	Path    string   `yaml:"path" json:"path"`
	Methods []string `yaml:"methods,omitempty" json:"methods,omitempty"`
	Timeout int      `yaml:"timeout,omitempty" json:"timeout,omitempty"` // 单位秒
	Role    string   `yaml:"role" json:"role"`                           // 角色
}

type LogInfo struct {
	Level    string `yaml:"level" json:"level"`
	FilePath string `yaml:"file_path" json:"filePath"`
}

type DatabaseInfo struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"database" json:"database"`
	Charset  string `yaml:"charset" json:"charset"`
	Debug    int8   `yaml:"debug" json:"debug"`
}

type RedisInfo struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Database int    `yaml:"database" json:"database"`
}

type ServiceInfo struct {
	Port     int    `yaml:"port" json:"port"`               // 服务端口
	Protocol string `yaml:"protocol" json:"protocol"`       // 服务协议 http,grpc,https
	SslPem   string `yaml:"sslPem" json:"sslPem,omitempty"` // ssl证书
	SslKey   string `yaml:"sslKey" json:"sslKey,omitempty"` // ssl私钥
}

type SessionInfo struct {
	ExpiresIn int    `yaml:"expiresIn" json:"expiresIn"` // 单位秒
	Storage   string `yaml:"storage" json:"storage"`     // 存储方式 redis,mysql
}
