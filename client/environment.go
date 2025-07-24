package client

type Env string

const (
	Environment Env = "https://petstore3.swagger.io/api/v3"
	MockServer  Env = "http://127.0.0.1:8082/v1/mock/demo/pets/0.2.0"
)

// String returns the environment as a string
func (e Env) String() string {
	return string(e)
}
