// Generated by: setup
// TypeWriter: set
// Directive: +test on Thing

package main

// Set is a modification of https://github.com/deckarep/golang-set
// The MIT License (MIT)
// Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)

// The primary type that represents a set
type ThingSet map[Thing]struct{}

// Creates and returns a reference to an empty set.
func NewThingSet() ThingSet {
	return make(ThingSet)
}

// Creates and returns a reference to a set from an existing slice
func NewThingSetFromSlice(s []Thing) ThingSet {
	a := NewThingSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}

// Adds an item to the current set if it doesn't already exist in the set.
func (set ThingSet) Add(i Thing) bool {
	_, found := set[i]
	set[i] = struct{}{}
	return !found //False if it existed already
}

// Determines if a given item is already in the set.
func (set ThingSet) Contains(i Thing) bool {
	_, found := set[i]
	return found
}

// Determines if the given items are all in the set
func (set ThingSet) ContainsAll(i ...Thing) bool {
	allSet := NewThingSetFromSlice(i)
	if allSet.IsSubset(set) {
		return true
	}
	return false
}

// Determines if every item in the other set is in this set.
func (set ThingSet) IsSubset(other ThingSet) bool {
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Determines if every item of this set is in the other set.
func (set ThingSet) IsSuperset(other ThingSet) bool {
	return other.IsSubset(set)
}

// Returns a new set with all items in both sets.
func (set ThingSet) Union(other ThingSet) ThingSet {
	unionedSet := NewThingSet()

	for elem := range set {
		unionedSet.Add(elem)
	}
	for elem := range other {
		unionedSet.Add(elem)
	}
	return unionedSet
}

// Returns a new set with items that exist only in both sets.
func (set ThingSet) Intersect(other ThingSet) ThingSet {
	intersection := NewThingSet()
	// loop over smaller set
	if set.Cardinality() < other.Cardinality() {
		for elem := range set {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range other {
			if set.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return intersection
}

// Returns a new set with items in the current set but not in the other set
func (set ThingSet) Difference(other ThingSet) ThingSet {
	differencedSet := NewThingSet()
	for elem := range set {
		if !other.Contains(elem) {
			differencedSet.Add(elem)
		}
	}
	return differencedSet
}

// Returns a new set with items in the current set or the other set but not in both.
func (set ThingSet) SymmetricDifference(other ThingSet) ThingSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Clears the entire set to be the empty set.
func (set *ThingSet) Clear() {
	*set = make(ThingSet)
}

// Allows the removal of a single item in the set.
func (set ThingSet) Remove(i Thing) {
	delete(set, i)
}

// Cardinality returns how many items are currently in the set.
func (set ThingSet) Cardinality() int {
	return len(set)
}

// Iter() returns a channel of type Thing that you can range over.
func (set ThingSet) Iter() <-chan Thing {
	ch := make(chan Thing)
	go func() {
		for elem := range set {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

// Equal determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set ThingSet) Equal(other ThingSet) bool {
	if set.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Returns a clone of the set.
// Does NOT clone the underlying elements.
func (set ThingSet) Clone() ThingSet {
	clonedSet := NewThingSet()
	for elem := range set {
		clonedSet.Add(elem)
	}
	return clonedSet
}
