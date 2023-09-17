package panel

type Panel interface {
	GetSize() (int, int)
	GetBuffer() []string
	GetConfig() *Config
	ChangeHandler(h func())
}
