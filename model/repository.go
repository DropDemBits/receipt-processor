package model

import "sync"

// Generic repository pattern to abstract shared CRUD functionality.
// This is thread-safe, and is expected to be used as a singleton instance.
type Repository[K comparable, V any] struct {
	values map[K]V
	// Protects [values]
	values_lock sync.Mutex
}

// Add an object to the repository.
// Returns false if key is already in the repository, true otherwise.
func (r *Repository[K, V]) Add(key K, value V) bool {
	r.values_lock.Lock()
	defer r.values_lock.Unlock()

	if _, exist := r.values[key]; exist {
		return false
	}

	r.values[key] = value
	return true
}

// Gets a single object from the repository.
// Returns the value and true if the key exists, false otherwise.
func (r *Repository[K, V]) GetSingle(key K) (V, bool) {
	r.values_lock.Lock()
	defer r.values_lock.Unlock()

	value, exist := r.values[key]

	return value, exist
}

// Updates a single object in the repository.
// Returns if the key existed, false otherwise.
func (r *Repository[K, V]) Update(key K, value V) bool {
	// Equivalent functionality for this in-memory database repository.
	return r.Add(key, value)
}

// Deletes an object from the repository.
// Returns true if there is a key to delete, false otherwise.
func (r *Repository[K, V]) Delete(key K, value V) bool {
	r.values_lock.Lock()
	defer r.values_lock.Unlock()

	if _, exist := r.values[key]; !exist {
		return false
	}

	delete(r.values, key)
	return true
}
