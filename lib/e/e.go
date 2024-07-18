package e

import "fmt"

// Функция для обработки ошибок и их обёртки в нужный формат
func Wrap(msg string, err error) error {
	return fmt.Errorf("[ERR] %s: %v", msg, err)
}
