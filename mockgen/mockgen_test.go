package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/mockgen/model"
)

func TestMakeArgString(t *testing.T) {
	testCases := []struct {
		argNames  []string
		argTypes  []string
		argString string
	}{
		{
			argNames:  nil,
			argTypes:  nil,
			argString: "",
		},
		{
			argNames:  []string{"arg0"},
			argTypes:  []string{"int"},
			argString: "arg0 int",
		},
		{
			argNames:  []string{"arg0", "arg1"},
			argTypes:  []string{"int", "bool"},
			argString: "arg0 int, arg1 bool",
		},
		{
			argNames:  []string{"arg0", "arg1"},
			argTypes:  []string{"int", "int"},
			argString: "arg0, arg1 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2"},
			argTypes:  []string{"bool", "int", "int"},
			argString: "arg0 bool, arg1, arg2 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2"},
			argTypes:  []string{"int", "bool", "int"},
			argString: "arg0 int, arg1 bool, arg2 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2"},
			argTypes:  []string{"int", "int", "bool"},
			argString: "arg0, arg1 int, arg2 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2"},
			argTypes:  []string{"int", "int", "int"},
			argString: "arg0, arg1, arg2 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3"},
			argTypes:  []string{"bool", "int", "int", "int"},
			argString: "arg0 bool, arg1, arg2, arg3 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3"},
			argTypes:  []string{"int", "bool", "int", "int"},
			argString: "arg0 int, arg1 bool, arg2, arg3 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3"},
			argTypes:  []string{"int", "int", "bool", "int"},
			argString: "arg0, arg1 int, arg2 bool, arg3 int",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3"},
			argTypes:  []string{"int", "int", "int", "bool"},
			argString: "arg0, arg1, arg2 int, arg3 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"bool", "int", "int", "int", "bool"},
			argString: "arg0 bool, arg1, arg2, arg3 int, arg4 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"int", "bool", "int", "int", "bool"},
			argString: "arg0 int, arg1 bool, arg2, arg3 int, arg4 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"int", "int", "bool", "int", "bool"},
			argString: "arg0, arg1 int, arg2 bool, arg3 int, arg4 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"int", "int", "int", "bool", "bool"},
			argString: "arg0, arg1, arg2 int, arg3, arg4 bool",
		},
		{
			argNames:  []string{"arg0", "arg1", "arg2", "arg3", "arg4"},
			argTypes:  []string{"int", "int", "bool", "bool", "int"},
			argString: "arg0, arg1 int, arg2, arg3 bool, arg4 int",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			s := makeArgString(tc.argNames, tc.argTypes)
			if s != tc.argString {
				t.Errorf("result == %q, want %q", s, tc.argString)
			}
		})
	}
}

func TestGetArgNames(t *testing.T) {
	for _, testCase := range []struct {
		name     string
		method   *model.Method
		expected []string
	}{
		{
			name: "NamedArg",
			method: &model.Method{
				In: []*model.Parameter{
					{
						Name: "firstArg",
						Type: &model.NamedType{Type: "int"},
					},
					{
						Name: "secondArg",
						Type: &model.NamedType{Type: "string"},
					},
				},
			},
			expected: []string{"firstArg", "secondArg"},
		},
		{
			name: "NotNamedArg",
			method: &model.Method{
				In: []*model.Parameter{
					{
						Name: "",
						Type: &model.NamedType{Type: "int"},
					},
					{
						Name: "",
						Type: &model.NamedType{Type: "string"},
					},
				},
			},
			expected: []string{"arg0", "arg1"},
		},
		{
			name: "MixedNameArg",
			method: &model.Method{
				In: []*model.Parameter{
					{
						Name: "firstArg",
						Type: &model.NamedType{Type: "int"},
					},
					{
						Name: "_",
						Type: &model.NamedType{Type: "string"},
					},
				},
			},
			expected: []string{"firstArg", "arg1"},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			g := generator{}

			result := g.getArgNames(testCase.method)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Fatalf("expected %s, got %s", result, testCase.expected)
			}
		})
	}
}

func Test_createPackageMap(t *testing.T) {
	tests := []struct {
		name            string
		importPath      string
		wantPackageName string
		wantOK          bool
	}{
		{"golang package", "context", "context", true},
		{"third party", "golang.org/x/tools/present", "present", true},
		//{"modules", "rsc.io/quote/v3", "quote", true},
		{"fail", "this/should/not/work", "", false},
	}
	var importPaths []string
	for _, t := range tests {
		importPaths = append(importPaths, t.importPath)
	}
	packages := createPackageMap(importPaths)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPackageName, gotOk := packages[tt.importPath]
			if gotPackageName != tt.wantPackageName {
				t.Errorf("createPackageMap() gotPackageName = %v, wantPackageName = %v", gotPackageName, tt.wantPackageName)
			}
			if gotOk != tt.wantOK {
				t.Errorf("createPackageMap() gotOk = %v, wantOK = %v", gotOk, tt.wantOK)
			}
		})
	}
}
