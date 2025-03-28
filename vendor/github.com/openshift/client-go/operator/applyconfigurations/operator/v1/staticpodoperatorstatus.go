// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// StaticPodOperatorStatusApplyConfiguration represents a declarative configuration of the StaticPodOperatorStatus type for use
// with apply.
type StaticPodOperatorStatusApplyConfiguration struct {
	OperatorStatusApplyConfiguration `json:",inline"`
	LatestAvailableRevisionReason    *string                        `json:"latestAvailableRevisionReason,omitempty"`
	NodeStatuses                     []NodeStatusApplyConfiguration `json:"nodeStatuses,omitempty"`
}

// StaticPodOperatorStatusApplyConfiguration constructs a declarative configuration of the StaticPodOperatorStatus type for use with
// apply.
func StaticPodOperatorStatus() *StaticPodOperatorStatusApplyConfiguration {
	return &StaticPodOperatorStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *StaticPodOperatorStatusApplyConfiguration) WithObservedGeneration(value int64) *StaticPodOperatorStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *StaticPodOperatorStatusApplyConfiguration) WithConditions(values ...*OperatorConditionApplyConfiguration) *StaticPodOperatorStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.OperatorStatusApplyConfiguration.Conditions = append(b.OperatorStatusApplyConfiguration.Conditions, *values[i])
	}
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *StaticPodOperatorStatusApplyConfiguration) WithVersion(value string) *StaticPodOperatorStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.Version = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *StaticPodOperatorStatusApplyConfiguration) WithReadyReplicas(value int32) *StaticPodOperatorStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.ReadyReplicas = &value
	return b
}

// WithLatestAvailableRevision sets the LatestAvailableRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevision field is set to the value of the last call.
func (b *StaticPodOperatorStatusApplyConfiguration) WithLatestAvailableRevision(value int32) *StaticPodOperatorStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.LatestAvailableRevision = &value
	return b
}

// WithGenerations adds the given value to the Generations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generations field.
func (b *StaticPodOperatorStatusApplyConfiguration) WithGenerations(values ...*GenerationStatusApplyConfiguration) *StaticPodOperatorStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerations")
		}
		b.OperatorStatusApplyConfiguration.Generations = append(b.OperatorStatusApplyConfiguration.Generations, *values[i])
	}
	return b
}

// WithLatestAvailableRevisionReason sets the LatestAvailableRevisionReason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevisionReason field is set to the value of the last call.
func (b *StaticPodOperatorStatusApplyConfiguration) WithLatestAvailableRevisionReason(value string) *StaticPodOperatorStatusApplyConfiguration {
	b.LatestAvailableRevisionReason = &value
	return b
}

// WithNodeStatuses adds the given value to the NodeStatuses field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NodeStatuses field.
func (b *StaticPodOperatorStatusApplyConfiguration) WithNodeStatuses(values ...*NodeStatusApplyConfiguration) *StaticPodOperatorStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNodeStatuses")
		}
		b.NodeStatuses = append(b.NodeStatuses, *values[i])
	}
	return b
}
