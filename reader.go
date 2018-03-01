package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type Data struct {
	R, C, F, N, B, T uint64
	RS               []Ride
}

type Ride struct {
	S, F   [2]uint
	ES, LF uint
}

// Fill will read pizza input information to the Data struct
func (d *Data) Fill(src io.Reader) error {
	scanner := bufio.NewScanner(src)

	if ok := scanner.Scan(); !ok {
		err := scanner.Err()
		return errors.Wrap(err, "reading first line from input source")
	}

	n, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d %d %d", &d.R, &d.C, &d.F, &d.N, &d.B, &d.T)
	if err != nil || n != 6 {
		return errors.Wrapf(err, "first line missed some important valuest, should contain 6 separate number (%d was readed)", n)
	}

	d.RS = make([]Ride, 0, d.R)
	i := 0 // row index
	var x1, y1, x2, y2, es, lf uint
	for scanner.Scan() {
		n, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d %d %d", &x1, &y1, &x2, &y2, &es, &lf)
		if err != nil || n != 6 {
			return errors.Wrapf(err, "ride line missed some important valuest, should contain 6 separate number (%d was readed)", n)
		}
		d.RS[i] = Ride{S: [2]uint{x1, y1}, F: [2]uint{x2, y2}, ES: es, LF: lf}
		i++
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "scanning input rides data")
	}

	return nil
}
