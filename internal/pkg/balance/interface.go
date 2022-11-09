package balance

// Database интерфейс отображающий функционал класса, отвечающий за взаимодействие с бд
type Database interface {
	Close() error
	Get(userid string) (float64, error)
	Change(userid string, new_balance float64) error
	Create(userid string, new_balance float64) error
}
