package internal

type GameServer struct {
	Name     string
	Template *Template
}

func GenerateServer(t *Template, env, cli map[string]string) (*GameServer, error) {
	gs := new(GameServer)
	gs.Template = t
	for k, v := range env {
		gs.Template.Env[k] = v
	}
	for k, v := range cli {
		gs.Template.Cli[k] = v
	}
	return gs, nil
}
