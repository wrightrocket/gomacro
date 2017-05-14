/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * type.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xtype

import (
	"go/types"
	"reflect"
)

func SameType(t, u Type) bool {
	return t == u || t != nil && u != nil && types.IdenticalIgnoreTags(t.GoType(), u.GoType())
}

func maketype(gtype types.Type, rtype reflect.Type) Type {
	kind := gtypeToKind(gtype)
	return &xtype{kind, gtype, rtype, nil}
}

// GoType returns the go/types.Type corresponding to the type.
func (t *xtype) GoType() types.Type {
	return t.gtype
}

// ReflectType returns a best-effort reflect.Type that approximates the type.
// It may be inexact for the following reasons:
// 1) missing reflect.NamedOf(): no way to programmatically create named types, or to access the underlying type of a named type
// 2) missing reflect.InterfaceOf(): interface types created at runtime will be approximated by structs
// 3) missing reflect.MethodOf(): method types created at runtime will be approximated by functions
//    whose first parameter is the receiver
// 4) reflect.StructOf() does not support embedded or unexported fields
// 5) go/reflect lacks the ability to create self-referencing types:
//    references to the type itself will be replaced by interface{}.
//
// Examples:
//    after invoking at runtime type2.NewStruct() and type2.NewNamed()
//    to create the following type:
//        type List struct { Elem int; Rest *List }
//    ReflectType will return a reflect.Type equivalent to:
//        struct { Elem int; Rest interface{} }
//    i.e. the type name will be missing due to limitation 1 above,
//    and the field 'Rest' will have type interface{} instead of *List due to limitation 5.
func (t *xtype) ReflectType() reflect.Type {
	return t.rtype
}

// Named returns whether the type is named.
// It returns false for unnamed types.
func (t *xtype) Named() bool {
	switch t.gtype.(type) {
	case *types.Basic, *types.Named:
		return true
	default:
		return false
	}
}

// Name returns the type's name within its package.
// It returns an empty string for unnamed types.
func (t *xtype) Name() string {
	switch gtype := t.gtype.(type) {
	case *types.Basic:
		return gtype.Name()
	case *types.Named:
		return gtype.Obj().Name()
	default:
		return ""
	}
}

// Pkg returns a named type's package, that is, the package where it was defined.
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// Pkg will return nil.
func (t *xtype) Pkg() *Package {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		return (*Package)(gtype.Obj().Pkg())
	default:
		return nil
	}
}

// PkgName returns a named type's package name, that is,
// the default name that the package provides when imported.
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// the package name will be the empty string.
func (t *xtype) PkgName() string {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		return gtype.Obj().Pkg().Name()
	default:
		return ""
	}
}

// PkgPath returns a named type's package path, that is, the import path
// that uniquely identifies the package, such as "encoding/base64".
// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int),
// the package path will be the empty string.
func (t *xtype) PkgPath() string {
	switch gtype := t.gtype.(type) {
	case *types.Named:
		return gtype.Obj().Pkg().Path()
	default:
		return ""
	}
}

// Size returns the number of bytes needed to store
// a value of the given type; it is analogous to unsafe.Sizeof.
func (t *xtype) Size() uintptr {
	return t.rtype.Size()
}

// String returns a string representation of a type.
func (t *xtype) String() string {
	if t == nil {
		return "invalid type"
	}
	return t.gtype.String()
}

/*
// Underlying returns the underlying type of a type.
func (t *xtype) Underlying() Type {
	return Type{t.underlying}
}
*/

func (t *xtype) underlying() types.Type {
	return t.gtype.Underlying()
}

// Kind returns the specific kind of the type.
func (t *xtype) Kind() reflect.Kind {
	if t == nil {
		return reflect.Invalid
	}
	return t.kind
}

// Implements reports whether the type implements the interface type u.
// It panics if u's Kind is not Interface
func (t *xtype) Implements(u Type) bool {
	if u.Kind() != reflect.Interface {
		panic("type2: non-interface type passed to Type.Implements")
	}
	return types.Implements(t.gtype, u.GoType().(*types.Interface))
}

// AssignableTo reports whether a value of the type is assignable to type u.
func (t *xtype) AssignableTo(u Type) bool {
	return types.AssignableTo(t.gtype, u.GoType())
}

// ConvertibleTo reports whether a value of the type is convertible to type u.
func (t *xtype) ConvertibleTo(u Type) bool {
	return types.ConvertibleTo(t.gtype, u.GoType())
}

// Comparable reports whether values of this type are comparable.
func (t *xtype) Comparable() bool {
	return types.Comparable(t.gtype)
}

// Zero returns a Value representing the zero value for the specified type.
func Zero(t Type) reflect.Value {
	return reflect.Zero(t.ReflectType())
}