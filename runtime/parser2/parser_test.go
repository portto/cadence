/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package parser2

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence/runtime/ast"
)

func TestParseExpression(t *testing.T) {

	expectedSimpleExpression := &ast.BinaryExpression{
		Operation: ast.OperationPlus,
		Left: &ast.IntegerExpression{
			Value: big.NewInt(1),
			Base:  10,
		},
		Right: &ast.BinaryExpression{
			Operation: ast.OperationMul,
			Left: &ast.IntegerExpression{
				Value: big.NewInt(2),
				Base:  10,
			},
			Right: &ast.IntegerExpression{
				Value: big.NewInt(3),
				Base:  10,
			},
		},
	}

	t.Run("simple, no spaces", func(t *testing.T) {
		result, errors := Parse("1+2*3")
		require.Empty(t, errors)

		assert.Equal(t,
			expectedSimpleExpression,
			result,
		)
	})

	t.Run("simple, spaces", func(t *testing.T) {
		result, errors := Parse("  1   +   2  *   3 ")
		require.Empty(t, errors)

		assert.Equal(t,
			expectedSimpleExpression,
			result,
		)
	})

	t.Run("repeated infix, same operator, left associative", func(t *testing.T) {
		result, errors := Parse("1 + 2 + 3")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.BinaryExpression{
				Operation: ast.OperationPlus,
				Left: &ast.BinaryExpression{
					Operation: ast.OperationPlus,
					Left: &ast.IntegerExpression{
						Value: big.NewInt(1),
						Base:  10,
					},
					Right: &ast.IntegerExpression{
						Value: big.NewInt(2),
						Base:  10,
					},
				},
				Right: &ast.IntegerExpression{
					Value: big.NewInt(3),
					Base:  10,
				},
			},
			result,
		)
	})

	t.Run("repeated infix, same operator, right associative", func(t *testing.T) {
		result, errors := Parse("1 ?? 2 ?? 3")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.BinaryExpression{
				Operation: ast.OperationNilCoalesce,
				Left: &ast.IntegerExpression{
					Value: big.NewInt(1),
					Base:  10,
				},
				Right: &ast.BinaryExpression{
					Operation: ast.OperationNilCoalesce,
					Left: &ast.IntegerExpression{
						Value: big.NewInt(2),
						Base:  10,
					},
					Right: &ast.IntegerExpression{
						Value: big.NewInt(3),
						Base:  10,
					},
				},
			},
			result,
		)
	})

	t.Run("mixed infix and prefix", func(t *testing.T) {
		result, errors := Parse("1 +- 2 ++ 3")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.BinaryExpression{
				Operation: ast.OperationPlus,
				Left: &ast.BinaryExpression{
					Operation: ast.OperationPlus,
					Left: &ast.IntegerExpression{
						Value: big.NewInt(1),
						Base:  10,
					},
					Right: &ast.UnaryExpression{
						Operation: ast.OperationMinus,
						Expression: &ast.IntegerExpression{
							Value: big.NewInt(2),
							Base:  10,
						},
					},
				},
				Right: &ast.UnaryExpression{
					Operation: ast.OperationPlus,
					Expression: &ast.IntegerExpression{
						Value: big.NewInt(3),
						Base:  10,
					},
				},
			},
			result,
		)
	})

	t.Run("nested expression", func(t *testing.T) {
		result, errors := Parse("(1 + 2) * 3")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.BinaryExpression{
				Operation: ast.OperationMul,
				Left: &ast.BinaryExpression{
					Operation: ast.OperationPlus,
					Left: &ast.IntegerExpression{
						Value: big.NewInt(1),
						Base:  10,
					},
					Right: &ast.IntegerExpression{
						Value: big.NewInt(2),
						Base:  10,
					},
				},
				Right: &ast.IntegerExpression{
					Value: big.NewInt(3),
					Base:  10,
				},
			},
			result,
		)
	})

	t.Run("less and greater", func(t *testing.T) {
		result, errors := Parse("1 < 2 > 3")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.BinaryExpression{
				Operation: ast.OperationGreater,
				Left: &ast.BinaryExpression{
					Operation: ast.OperationLess,
					Left: &ast.IntegerExpression{
						Value: big.NewInt(1),
						Base:  10,
					},
					Right: &ast.IntegerExpression{
						Value: big.NewInt(2),
						Base:  10,
					},
				},
				Right: &ast.IntegerExpression{
					Value: big.NewInt(3),
					Base:  10,
				},
			},
			result,
		)
	})

	t.Run("array expression", func(t *testing.T) {
		result, errors := Parse("[ 1,2 + 3, 4  ,  5 ]")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.ArrayExpression{
				Values: []ast.Expression{
					&ast.IntegerExpression{
						Value: big.NewInt(1),
						Base:  10,
					},
					&ast.BinaryExpression{
						Operation: ast.OperationPlus,
						Left: &ast.IntegerExpression{
							Value: big.NewInt(2),
							Base:  10,
						},
						Right: &ast.IntegerExpression{
							Value: big.NewInt(3),
							Base:  10,
						},
					},
					&ast.IntegerExpression{
						Value: big.NewInt(4),
						Base:  10,
					},
					&ast.IntegerExpression{
						Value: big.NewInt(5),
						Base:  10,
					},
				},
			},
			result,
		)
	})

	t.Run("dictionary expression", func(t *testing.T) {
		result, errors := Parse("{ 1:2 + 3, 4  :  5 }")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.DictionaryExpression{
				Entries: []ast.Entry{
					{
						Key: &ast.IntegerExpression{
							Value: big.NewInt(1),
							Base:  10,
						},
						Value: &ast.BinaryExpression{
							Operation: ast.OperationPlus,
							Left: &ast.IntegerExpression{
								Value: big.NewInt(2),
								Base:  10,
							},
							Right: &ast.IntegerExpression{
								Value: big.NewInt(3),
								Base:  10,
							},
						},
					},
					{
						Key: &ast.IntegerExpression{
							Value: big.NewInt(4),
							Base:  10,
						},
						Value: &ast.IntegerExpression{
							Value: big.NewInt(5),
							Base:  10,
						},
					},
				},
			},
			result,
		)
	})

	t.Run("identifier in addition", func(t *testing.T) {
		result, errors := Parse("a + 3")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.BinaryExpression{
				Operation: ast.OperationPlus,
				Left: &ast.IdentifierExpression{
					Identifier: ast.Identifier{
						Identifier: "a",
					},
				},
				Right: &ast.IntegerExpression{
					Value: big.NewInt(3),
					Base:  10,
				},
			},
			result,
		)
	})

	t.Run("path expression", func(t *testing.T) {
		result, errors := Parse("/foo/bar")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.PathExpression{
				Domain:     ast.Identifier{Identifier: "foo"},
				Identifier: ast.Identifier{Identifier: "bar"},
			},
			result,
		)
	})

	t.Run("conditional", func(t *testing.T) {
		result, errors := Parse("a ? b : c ? d : e")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.ConditionalExpression{
				Test: &ast.IdentifierExpression{
					Identifier: ast.Identifier{
						Identifier: "a",
					},
				},
				Then: &ast.IdentifierExpression{
					Identifier: ast.Identifier{
						Identifier: "b",
					},
				},
				Else: &ast.ConditionalExpression{
					Test: &ast.IdentifierExpression{
						Identifier: ast.Identifier{
							Identifier: "c",
						},
					},
					Then: &ast.IdentifierExpression{
						Identifier: ast.Identifier{
							Identifier: "d",
						},
					},
					Else: &ast.IdentifierExpression{
						Identifier: ast.Identifier{
							Identifier: "e",
						},
					},
				},
			},
			result,
		)
	})

	t.Run("boolean expressions", func(t *testing.T) {
		result, errors := Parse("true + false")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.BinaryExpression{
				Operation: ast.OperationPlus,
				Left: &ast.BoolExpression{
					Value: true,
				},
				Right: &ast.BoolExpression{
					Value: false,
				},
			},
			result,
		)
	})

	t.Run("move operator", func(t *testing.T) {
		result, errors := Parse("(<-x)")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.UnaryExpression{
				Operation: ast.OperationMove,
				Expression: &ast.IdentifierExpression{
					Identifier: ast.Identifier{
						Identifier: "x",
					},
				},
			},
			result,
		)
	})

	t.Run("invocation", func(t *testing.T) {
		result, errors := Parse("f(1,2)")
		require.Empty(t, errors)

		assert.Equal(t,
			&ast.InvocationExpression{
				InvokedExpression: &ast.IdentifierExpression{
					Identifier: ast.Identifier{
						Identifier: "f",
						Pos:        ast.Position{Offset: 0, Line: 0, Column: 1},
					},
				},
				Arguments: []*ast.Argument{
					{
						Label: "",
						Expression: &ast.IntegerExpression{
							Value: big.NewInt(1),
							Base:  10,
							Range: ast.Range{
								StartPos: ast.Position{Offset: 0, Line: 1, Column: 3},
								EndPos:   ast.Position{Offset: 0, Line: 1, Column: 3},
							},
						},
					},
					{
						Label: "",
						Expression: &ast.IntegerExpression{
							Value: big.NewInt(2),
							Base:  10,
							Range: ast.Range{
								StartPos: ast.Position{Offset: 0, Line: 1, Column: 5},
								EndPos:   ast.Position{Offset: 0, Line: 1, Column: 5},
							},
						},
					},
				},
				EndPos: ast.Position{Offset: 0, Line: 1, Column: 6},
			},
			result,
		)
	})
}
