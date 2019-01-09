package config

// Config monta o retorno padrão
type Config struct {
	PageAccessToken string
	VerifyToken     string
	AppSecret       string
	GraphURL        string
}

// Configuracoes centraliza as opções da api
func Configuracoes() Config {
	r := Config{
		GraphURL: "https://graph.facebook.com/v2.6/" /*os.Getenv("GRAPHURL")*/}

	return r
}
