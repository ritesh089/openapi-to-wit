package main

// WITType defines the basic types in WebAssembly
type WITType string

const (
	WITTypeI32    WITType = "i32"
	WITTypeI64    WITType = "i64"
	WITTypeF32    WITType = "f32"
	WITTypeF64    WITType = "f64"
	WITTypeString WITType = "string"
	WITTypeList   WITType = "list"
	WITTypeStruct WITType = "struct"
)

// Parameter represents a parameter in a WIT function signature
type Parameter struct {
	Name string       // Parameter name
	Type WITType      // Parameter type (basic or complex)
	Ref  *ComplexType // Optional: Reference to a complex type (for structs/lists)
}

// ComplexType represents a user-defined complex type (e.g., struct or list)
type ComplexType struct {
	Name     string      // The name of the complex type (struct, list, etc.)
	Fields   []Parameter // Fields of the struct or items of the list
	ItemType WITType     // For lists, the type of the list items
	IsStruct bool        // Whether it's a struct
	IsList   bool        // Whether it's a list
}

// FunctionSignature represents the signature of a function in WIT
type FunctionSignature struct {
	Name       string       // Function name
	Parameters []Parameter  // Function parameters
	ReturnType WITType      // Function return type (could be complex)
	ReturnRef  *ComplexType // Optional: Reference to a complex return type
}

// WITModule represents a WIT module containing functions
type WITModule struct {
	ModuleName string              // Module name
	Functions  []FunctionSignature // List of functions in the module
}
