package goscheme

var scope *Scope

func init() {
	scope = NewScope()
	scope.AddEnv()
}

// Env describes an environment
type Env map[string]Value

// Scope describes a scope
type Scope struct {
	envs []*Env
}

// NewScope creates a new, empty scope
func NewScope() *Scope {
	scope := &Scope{}
	scope.envs = make([]*Env, 0)
	return scope
}

// Dup duplicates an existing scope.
func (s *Scope) Dup() *Scope {
	scope := &Scope{}
	scope.envs = make([]*Env, len(s.envs))
	copy(scope.envs, s.envs)
	return scope
}

// Env returns the environment contained within a scope.
func (s *Scope) Env() *Env {
	if len(s.envs) > 0 {
		return s.envs[len(s.envs)-1]
	}
	return nil
}

// AddEnv adds a new environment to the scope.
func (s *Scope) AddEnv() *Env {
	env := make(Env)
	s.envs = append(s.envs, &env)
	return &env
}

// DropEnv removes an environment from the scope.
func (s *Scope) DropEnv() *Env {
	s.envs[len(s.envs)-1] = nil
	s.envs = s.envs[:len(s.envs)-1]
	return s.Env()
}

// Create adds a key/value pair to the environment within a scope.
func (s *Scope) Create(key string, value Value) Value {
	env := *s.Env()
	env[key] = value
	return value
}

// Set sets the value for a given key within a scope's enviornment.
func (s *Scope) Set(key string, value Value) Value {
	for i := len(s.envs) - 1; i >= 0; i-- {
		env := *s.envs[i]
		if _, ok := env[key]; ok {
			env[key] = value
			return value
		}
	}
	return s.Create(key, value)
}

// Get retrieves the value for a given key within a scope's enviornment.
func (s *Scope) Get(key string) (val Value, ok bool) {
	for i := len(s.envs) - 1; i >= 0; i-- {
		env := *s.envs[i]
		if val, ok = env[key]; ok {
			break
		}
	}
	return
}
