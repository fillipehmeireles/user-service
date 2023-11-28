package customerrors

import "fmt"

type ConfigKeyNotFound struct {
	Key string
}

func (cknf *ConfigKeyNotFound) Error() string {
	return fmt.Sprintf("[Config] Config %s not found", cknf.Key)
}
