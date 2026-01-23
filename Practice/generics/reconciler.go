package main

import (
	"sync"
	"time"
)

type Resource interface {
	// Metadata access
	GetName() string
	GetNamespace() string
	GetUID() string
	GetGeneration() int64 // Increments when spec changes
	SetGeneration(int64)

	// Deletion handling
	GetDeletionTimestamp() *time.Time
	SetDeletionTimestamp(*time.Time)
	IsBeingDeleted() bool

	// Finalizers (for cleanup before deletion)
	GetFinalizers() []string
	SetFinalizers([]string)
	HasFinalizer(string) bool
	AddFinalizer(string)
	RemoveFinalizer(string)

	// Status tracking
	GetObservedGeneration() int64 // Last generation we reconciled
	SetObservedGeneration(int64)
	GetConditions() []Condition
	SetCondition(Condition)

	// Deep copy for cache updates
	DeepCopy() Resource
}

type Reconciler[T Resource] struct {
	// In-memory cache of resources (key: namespace/name)
	cache   map[string]T
	cacheMu sync.RWMutex

	// The actual reconciliation logic (passed in by user)
	reconcileFunc func(T) error

	// Work queue for processing
	queue *WorkQueue[ReconcileRequest]

	// For graceful shutdown
	stopCh chan struct{}
}

// Create a new reconciler
func NewReconciler[T Resource](reconcileFunc func(T) error) *Reconciler[T] {

}

// Start the reconciliation loop (runs in background)
func (r *Reconciler[T]) Start() {

}

// Stop the reconciler
func (r *Reconciler[T]) Stop() {

}

// Add/Update a resource (triggers reconciliation)
func (r *Reconciler[T]) CreateOrUpdate(resource T) {

}

// Delete a resource (triggers finalizer handling)
func (r *Reconciler[T]) Delete(namespace, name string) {

}

// The main reconciliation loop
func (r *Reconciler[T]) processNextItem() {

}

// Handle a single reconciliation request
func (r *Reconciler[T]) reconcile(req ReconcileRequest) error {

}

type ReconcileRequest struct {
	Namespace string
	Name      string
}

func (r ReconcileRequest) NamespacedName() string {
	if r.Namespace == "" {
		return r.Name // Cluster-scoped resource
	}
	return r.Namespace + "/" + r.Name
}

// Optional: Implement String() for better logging
func (r ReconcileRequest) String() string {
	return r.NamespacedName()
}

type WorkQueue[T comparable] struct {
	queue      []T
	dirty      map[T]struct{} // Items already in queue (deduplication)
	processing map[T]struct{} // Currently being processed
	mu         sync.Mutex
	cond       *sync.Cond // Signal when items added

	shuttingDown bool
}

type Condition struct {
	Type               string
	Status             bool
	Reason             string
	Message            string
	LastTransitionTime time.Time
}
