package panel

type Panel interface {
	GetSize() (int, int)
	GetBuffer() []*Line
	GetConfig() *Config
	ChangeHandler(h func())
}
