/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * builtin.go
 *
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"io/ioutil"
	r "reflect"
)

func builtinAppend(env *Env, args ...ast.Expr) (r.Value, []r.Value) {
	n := len(args)
	if n < 1 {
		return env.Errorf("builtin append() expects at least one argument, found %d", n)
	}
	elems := env.evalExprs(args)
	slice := elems[0]
	return r.Append(slice, elems[1:]...), nil
}

func callCap(arg interface{}) int {
	return r.ValueOf(arg).Cap()
}

func callClose(channel interface{}) {
	r.ValueOf(channel).Close()
}

func callCopy(dst, src interface{}) int {
	return r.Copy(r.ValueOf(dst), r.ValueOf(src))
}

func callDelete(m interface{}, key interface{}) {
	r.ValueOf(m).SetMapIndex(r.ValueOf(key), Nil)
}

func callLen(arg interface{}) int {
	return r.ValueOf(arg).Len()
}

func builtinMake(env *Env, args ...ast.Expr) (r.Value, []r.Value) {
	n := len(args)
	if n < 1 || n > 3 {
		return env.Errorf("builtin make() expects one, two or three arguments, found %d", n)
	}
	t := env.evalType(args[0])
	values := env.evalExprs(args[1:])
	n--
	ret := Nil
	switch t.Kind() {
	case r.Chan:
		buffer := 0
		if n > 0 {
			buffer = int(values[0].Int())
		}
		ret = r.MakeChan(t, buffer)
	case r.Map:
		ret = r.MakeMap(t)
	case r.Slice:
		length := 0
		if n > 0 {
			length = int(values[0].Int())
		}
		capacity := length
		if n > 1 {
			capacity = int(values[1].Int())
		}
		ret = r.MakeSlice(t, length, capacity)
	}
	return ret, nil
}

func builtinNew(env *Env, args ...ast.Expr) (r.Value, []r.Value) {
	n := len(args)
	if n != 1 {
		return env.Errorf("builtin new() expects exactly one argument, found %d", n)
	}
	t := env.evalType(args[0])
	return r.New(t), nil
}

func callPanic(arg interface{}) {
	panic(Panic{arg})
}

func callReadFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		callPanic(err)
	}
	return string(bytes)
}

func callReadDir(dirname string) []string {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		callPanic(err)
	}
	n := len(files)
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = files[i].Name()
	}
	return names
}

func callRecover() interface{} {
	return recover()
}

func callSlice(args ...interface{}) []interface{} {
	return args
}

func (env *Env) addBuiltins() {
	binds := env.Binds

	binds["append"] = r.ValueOf(Builtin{builtinAppend})
	binds["cap"] = r.ValueOf(callCap)
	binds["close"] = r.ValueOf(callClose)
	binds["copy"] = r.ValueOf(callCopy)
	binds["DeepEqual"] = r.ValueOf(r.DeepEqual)
	binds["delete"] = r.ValueOf(callDelete)
	binds["Eval"] = r.ValueOf(func(node ast.Node) interface{} {
		return toInterface(env.Eval1(node))
	})
	binds["EvalN"] = r.ValueOf(func(node ast.Node) []interface{} {
		return toInterfaces(packValues(env.Eval(node)))
	})
	binds["len"] = r.ValueOf(callLen)
	binds["new"] = r.ValueOf(Builtin{builtinNew})
	binds["MacroExpand"] = r.ValueOf(func(in ast.Node) (out ast.Node, expanded bool) {
		return env.MacroExpand(in)
	})
	binds["MacroExpand1"] = r.ValueOf(func(in ast.Node) (out ast.Node, expanded bool) {
		return env.MacroExpand1(in)
	})
	binds["MacroExpandCodewalk"] = r.ValueOf(func(in ast.Node) (out ast.Node, expanded bool) {
		return env.MacroExpandCodewalk(in)
	})
	binds["make"] = r.ValueOf(Builtin{builtinMake})
	binds["nil"] = Nil
	binds["panic"] = r.ValueOf(callPanic)
	binds["Parse"] = r.ValueOf(func(src interface{}) []ast.Node {
		return env.Parse(src)
	})
	binds["Parse1"] = r.ValueOf(func(src interface{}) ast.Node {
		return env.Parse1(src)
	})
	binds["println"] = r.ValueOf(func(args ...interface{}) {
		values := toValues(args)
		env.FprintValues(env.Stdout, values...)
	})
	binds["ReadDir"] = r.ValueOf(callReadDir)
	binds["ReadFile"] = r.ValueOf(callReadFile)
	// binds["recover"] = r.ValueOf(callRecover) // does not work! recover() works only inside a deferred function (but not any function called by it)
	binds["Slice"] = r.ValueOf(callSlice)
	binds["String"] = r.ValueOf(func(args ...interface{}) string {
		return env.toString("", args...)
	})

	types := env.Types
	types["bool"] = r.TypeOf(false)
	types["byte"] = r.TypeOf(byte(0))
	types["complex64"] = r.TypeOf(complex(float32(0), float32(0)))
	types["complex128"] = r.TypeOf(complex(float64(0), float64(0)))
	types["float32"] = r.TypeOf(float32(0))
	types["float64"] = r.TypeOf(float64(0))
	types["int"] = r.TypeOf(int(0))
	types["int8"] = r.TypeOf(int8(0))
	types["int16"] = r.TypeOf(int16(0))
	types["int32"] = r.TypeOf(int32(0))
	types["int64"] = r.TypeOf(int64(0))
	types["rune"] = r.TypeOf(rune(0))
	types["string"] = r.TypeOf("")
	types["uint"] = r.TypeOf(uint(0))
	types["uint8"] = r.TypeOf(uint8(0))
	types["uint16"] = r.TypeOf(uint16(0))
	types["uint32"] = r.TypeOf(uint32(0))
	types["uint64"] = r.TypeOf(uint64(0))
	types["uintptr"] = r.TypeOf(uintptr(0))
}

func (env *Env) addInterpretedBuiltins() {
	if false {
		line := "func not(flag bool) bool { if flag { return false } else { return true } }"
		env.EvalList(env.Parse(line))
	}
	if false {
		// Factorial(1000000): eval() elapsed time: 1.233714899 s
		line := "func Factorial(n int) int { t := 1; for i := 2; i <= n; i=i+1 { t = t * i }; t }"
		env.EvalList(env.Parse(line))
	}
}
