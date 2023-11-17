package logger

type Config struct {
	JSON  bool   `yaml:"json"`
	Level string `yaml:"level"`
	// true 启用日志采样，其中相同的日志消息和级别将受到限制。
	// 我们有两层采样
	// 1. 全局采样 - 一秒钟内，它将记录第一个 SampleInitial，然后记录每个 SampleInterval 消息。
	// 2. 每个参与者/轨道采样 - 与 Logger.WithItemSampler() 一起使用。这将用于节流
	// 特定参与者/轨道的日志。
	Sample bool `yaml:"sample,omitempty"`

	// 每个服务器的全局采样
	// 采样时，会记录前N条日志
	SampleInitial int `yaml:"sample_initial,omitempty"`
	// 采样时，每第 M 个日志都会被记录
	SampleInterval int `yaml:"sample_interval,omitempty"`

	// 参与者/轨道级别采样
	ItemSampleSeconds  int `yaml:"item_sample_seconds,omitempty"`
	ItemSampleInitial  int `yaml:"item_sample_initial,omitempty"`
	ItemSampleInterval int `yaml:"item_sample_interval,omitempty"`
}
