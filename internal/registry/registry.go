package registry

type Key struct {
	Year int
	Day  int
}

type registry struct {
	solvers map[Key]Solver
}

func NewRegistry() Registry {
	return &registry{
		solvers: make(map[Key]Solver),
	}
}

func (r *registry) Register(year int, day int, solver Solver) error {
	if solver == nil {
		return &RegistrationError{Year: year, Day: day, Err: ErrNilSolver}
	}
	if r.solvers == nil {
		r.solvers = make(map[Key]Solver)
	}

	key := Key{year, day}
	if _, exists := r.solvers[key]; exists {
		return &RegistrationError{Year: year, Day: day, Err: ErrSolverAlreadyRegistered}
	}

	r.solvers[key] = solver
	return nil
}

func (r *registry) GetSolver(year int, day int) (Solver, bool) {
	if r.solvers == nil {
		return nil, false
	}

	key := Key{year, day}
	solver, ok := r.solvers[key]
	return solver, ok
}
