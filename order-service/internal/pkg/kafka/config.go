package kafka

type Config struct {
	// Поля User и Password будут использованы в будущем
	// для реализации механизма SASL аутентификации
	User     string
	Password string
	Hosts    []string
}
