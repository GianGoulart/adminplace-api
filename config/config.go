package config

// Config monta o retorno padrão
type Config struct {
	GraphURL  string
	Community string
}

// Configuracoes centraliza as opções da api
func Configuracoes() Config {
	r := Config{
		GraphURL:  "https://graph.facebook.com/v2.6/", /*os.Getenv("GRAPHURL")*/
		Community: "394761294206762"}

	return r
}
