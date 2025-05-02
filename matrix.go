package matrix

import (
  "strings"
  "errors"
  "fmt"
  "strconv"
)

// Define the Matrix type here.
type Matrix interface {
  Cols() [][]int
  Rows() [][]int
  Set(int, int, int) bool
}

type matrix struct {
  data [][]int
}

func (m *matrix) width() int {
  if len(m.data) == 0 {
    return 0
  }
  return len(m.data[0])
}

func (m *matrix) addRow(s string) error {
  elements := strings.Split(strings.TrimSpace(s), " ")  
  if m.width() > 0 && len(elements) != m.width() {
    return fmt.Errorf(
      "expected %d elements in row, found %d in '%s'",
      m.width(),
      len(elements),
      s,
    )
  }

  row := make([]int, 0, len(elements))
  for _, ele := range elements {
    if val, err := strconv.Atoi(ele); err != nil {
      return fmt.Errorf("%s is not an int: %v", ele, err)
    } else {
      row = append(row, val)
    }
  }
  m.data = append(m.data, row)
  return nil
}

func New(s string) (Matrix, error) {
  if strings.TrimSpace(s) == "" {
    return nil, errors.New("empty input")
  }
 
  rows := strings.Split(s, "\n")
  data := make([][]int, 0, len(rows))

  m := matrix{ data }  
  for _, row := range rows {
    if err := m.addRow(row); err != nil {
      return nil, err
    }
  }

  return &m, nil
}
  
// Cols and Rows must return the results without affecting the matrix.
func (m *matrix) Cols() [][]int {
  transposed := make([][]int, 0, m.width())

  width := len(m.data)
  for i := 0; i < cap(transposed); i++ {
    transposed = append(transposed, make([]int, width)) 
  }

  for r, row := range m.data {
    for c, val := range row {
      transposed[c][r] = val
    }
  }
	return transposed
}

func (m *matrix) Rows() [][]int {
  data := make([][]int, len(m.data))
   
  width := m.width()

  for i, row := range m.data {
    data[i] = make([]int, width)
    copy(data[i], row)
  }
  return data
}

func (m *matrix) Set(row, col, val int) bool {
  if row < 0 || row >= len(m.data) {
    return false
  }
  if col < 0 || col >= m.width() {
    return false
  }
  m.data[row][col] = val
  return true
}
