package matrix

import (
  "strings"
  "errors"
  "fmt"
  "strconv"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
  if strings.TrimSpace(s) == "" {
    return nil, errors.New("empty input")
  }
 
  rows := strings.Split(s, "\n")
  matrix := Matrix{}
  
  for _, str := range rows {
    cols := strings.Split(strings.TrimSpace(str), " ")
    if len(matrix) > 0 && len(matrix[0]) != len(cols) {
      return nil, fmt.Errorf(
        "expected %d elements in row, found %d in '%s'",
        len(matrix[0]),
        len(cols),
        str,
      )
    }

    row := make([]int, 0, len(cols))
    for _, ele := range cols {
      if val, err := strconv.Atoi(ele); err != nil {
        return nil, fmt.Errorf("'%s' is not an int: %v", ele, err)
      } else {
        row = append(row, val)
      }
    }
    matrix = append(matrix, row)
  }

  return matrix, nil
}
  
// Cols and Rows must return the results without affecting the matrix.
func (matrix Matrix) Cols() [][]int {
  transposed := make([][]int, 0, len(matrix[0]))

  width := len(matrix)
  for i := 0; i < cap(transposed); i++ {
    transposed = append(transposed, make([]int, width)) 
  }

  for r, row := range matrix {
    for c, val := range row {
      transposed[c][r] = val
    }
  }
	return transposed
}

func (matrix Matrix) Rows() [][]int {
  data := make([][]int, len(matrix))
   
  for r, row := range matrix {
    data[r] = make([]int, len(matrix[0]))
    copy(data[r], row)
  }
  return data
}

func (matrix Matrix) Set(row, col, val int) bool {
  if row < 0 || row >= len(matrix) {
    return false
  }
  if col < 0 || col >= len(matrix[0]) {
    return false
  }
  matrix[row][col] = val
  return true
}
