// Copyright 2017 the original author or authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pbfparser_test

import (
	"fmt"
	"io"
	"log"
	"os"

	parser "github.com/maguro/pbfparser"
)

func Example() {
	in, err := os.Open("testdata/greater-london.osm.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	d, err := parser.NewDecoder(in)

	// use more memory from the start, it is faster
	d.SetBufferSize(33554432)

	// start decoding with several goroutines, it is faster
	d.Start(2)

	var nc, wc, rc uint64
	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			case *parser.Node:
				// Process Node v.
				nc++
			case *parser.Way:
				// Process Way v.
				wc++
			case *parser.Relation:
				// Process Relation v.
				rc++
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)
	// Output:
	// Nodes: 2729006, Ways: 459055, Relations: 12833
}