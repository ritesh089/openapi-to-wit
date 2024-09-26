package main

import (
	"fmt"
	"strings"
)

// GenerateWITType converts a parameter to its WIT representation
func GenerateWITType(param Parameter) string {
	if param.Ref != nil {
		if param.Ref.IsStruct {
			return param.Ref.Name // Use struct name in WIT
		} else if param.Ref.IsList {
			return fmt.Sprintf("list<%s>", param.Ref.ItemType)
		}
	}
	return string(param.Type)
}

// GenerateComplexTypes generates WIT type definitions for complex types (e.g., structs)
func GenerateComplexTypes(complexTypes map[string]ComplexType) string {
	var witBuilder strings.Builder
	for _, ctype := range complexTypes {
		if ctype.IsStruct {
			witBuilder.WriteString(fmt.Sprintf("type %s = record {\n", ctype.Name))
			for _, field := range ctype.Fields {
				witBuilder.WriteString(fmt.Sprintf("    %s: %s,\n", field.Name, GenerateWITType(field)))
			}
			witBuilder.WriteString("}\n\n")
		} else if ctype.IsList {
			witBuilder.WriteString(fmt.Sprintf("type %s = list<%s>\n\n", ctype.Name, ctype.ItemType))
		}
	}
	return witBuilder.String()
}

// GenerateWIT generates WIT for the provided module
func GenerateWIT(module WITModule, complexTypes map[string]ComplexType) string {
	var witBuilder strings.Builder

	// Generate complex type definitions
	witBuilder.WriteString(GenerateComplexTypes(complexTypes))

	// Generate world and function definitions
	witBuilder.WriteString(fmt.Sprintf("world %s {\n", module.ModuleName))
	for _, fn := range module.Functions {
		witBuilder.WriteString(fmt.Sprintf("    %s: func(", fn.Name))
		for i, param := range fn.Parameters {
			if i > 0 {
				witBuilder.WriteString(", ")
			}
			witBuilder.WriteString(fmt.Sprintf("%s: %s", param.Name, GenerateWITType(param)))
		}
		witBuilder.WriteString(")")

		// Handle return types
		if fn.ReturnType == WITTypeStruct && fn.ReturnRef != nil {
			witBuilder.WriteString(fmt.Sprintf(" -> %s", fn.ReturnRef.Name))
		} else if fn.ReturnType != "" {
			witBuilder.WriteString(fmt.Sprintf(" -> %s", fn.ReturnType))
		}

		witBuilder.WriteString("\n")
	}
	witBuilder.WriteString("}\n")

	return witBuilder.String()
}
