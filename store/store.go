package store

type Store interface{
	Service() Service
}