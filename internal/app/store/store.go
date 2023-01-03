package store

// here we declare interface for stores
type Store interface {
	User() UserRepository
}
