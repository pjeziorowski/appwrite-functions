package create_project

type CreateProjectInput struct {
	Name string
}

type CreateProjectOutput struct {
	Id   int32
	Name string
	Url  string
}

type Mutation struct {
	CreateProject *CreateProjectOutput
}

type CreateProjectArgs struct {
	Input CreateProjectInput
}
