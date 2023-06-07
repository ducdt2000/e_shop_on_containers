package seed_work

type IRepository interface {
	GetUnitOfWork() IUnitOfWork
}
