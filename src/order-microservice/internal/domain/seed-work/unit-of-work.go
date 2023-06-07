package seed_work

type IUnitOfWork interface {
	SaveChanges() int
	SaveEntities() bool
}
