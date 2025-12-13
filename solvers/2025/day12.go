package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

type size struct {
	piece  []string
	hashes int
}

type board struct {
	width  int
	height int
	counts []int
}

func getSizesAndBoards(data string) ([]*size, []*board) {
	sizes := []*size{}
	boards := []*board{}

	var cpieces []string
	chash := 0
	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, "x") {
			elems := strings.Split(line, ":")
			selems := strings.Split(elems[0], "x")
			width, err := strconv.ParseInt(selems[0], 10, 64)
			if err != nil {
				log.Fatalf("Cannot parse width: %v", err)
			}
			height, err := strconv.ParseInt(selems[1], 10, 64)
			if err != nil {
				log.Fatalf("Cannot parse width: %v", err)
			}

			var counts []int
			for _, c := range strings.Fields(elems[1]) {
				nv, err := strconv.ParseInt(c, 10, 64)
				if err != nil {
					log.Fatalf("Cannot parse count: %v", err)
				}
				counts = append(counts, int(nv))
			}

			boards = append(boards, &board{
				width:  int(width),
				height: int(height),
				counts: counts,
			})
		} else {
			if len(strings.TrimSpace(line)) == 0 {
				sizes = append(sizes, &size{
					piece:  cpieces,
					hashes: chash,
				})
				cpieces = []string{}
				chash = 0
				continue
			} else if !strings.Contains(line, ":") {
				cpieces = append(cpieces, strings.TrimSpace(line))
				chash += strings.Count(line, "#")
			}
		}
	}

	return sizes, boards
}

func buildBoard(width, height int) [][]bool {
	board := make([][]bool, height)
	for i := 0; i < height; i++ {
		board[i] = make([]bool, width)
	}
	return board
}

func buildPieces(sizes []*size, counts []int) [][][]bool {
	var pieces [][][]bool
	for j, c := range counts {
		for i := 0; i < c; i++ {
			pieces = append(pieces, buildPiece(sizes[j]))
		}
	}
	return pieces
}

func buildPiece(s *size) [][]bool {
	var res [][]bool
	for _, line := range s.piece {
		row := []bool{}
		for _, c := range line {
			if c == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		res = append(res, row)
	}
	log.Printf("Built piece: %v", res)
	return res
}

func rotate(piece [][]bool, rotation int) [][]bool {
	if rotation == 0 {
		return piece
	}

	var newPiece [][]bool

	if rotation == 1 {
		for y := len(piece[0]) - 1; y >= 0; y-- {
			var nrow []bool
			for x := 0; x < len(piece); x++ {
				nrow = append(nrow, piece[y][x])
			}
			newPiece = append(newPiece, nrow)
		}
	} else if rotation == 2 {
		for x := len(piece) - 1; x >= 0; x-- {
			var nrow []bool
			for y := len(piece[0]) - 1; y >= 0; y-- {
				nrow = append(nrow, piece[y][x])
			}
			newPiece = append(newPiece, nrow)
		}
	} else if rotation == 3 {
		for y := len(piece[0]) - 1; y >= 0; y-- {
			var nrow []bool
			for x := len(piece) - 1; x >= 0; x-- {
				nrow = append(nrow, piece[y][x])
			}
			newPiece = append(newPiece, nrow)
		}
	}
	//log.Printf("Rotated piece: %v", newPiece)
	return newPiece
}

func copyBoard(board [][]bool) [][]bool {
	nboard := make([][]bool, len(board))
	for y := 0; y < len(board); y++ {
		nboard[y] = make([]bool, len(board[0]))
		for x := 0; x < len(board[0]); x++ {
			nboard[y][x] = board[y][x]
		}
	}
	return nboard
}

func canPlace(board [][]bool, piece [][]bool, xoff, yoff, rotation int) ([][]bool, bool) {
	rpiece := rotate(piece, rotation)
	cboard := copyBoard(board)

	//log.Printf("Trying to place %v in %v at %v,%v", rpiece, cboard, xoff, yoff)

	for y := 0; y < len(rpiece); y++ {
		for x := 0; x < len(rpiece[0]); x++ {
			if rpiece[y][x] {
				if !cboard[y+yoff][x+xoff] {
					cboard[y+yoff][x+xoff] = true
				} else {
					return cboard, false
				}
			}
		}
	}

	return cboard, true
}

func place(board [][]bool, pieces [][][]bool) bool {
	if len(pieces) == 0 {
		return true
	}
	piece := pieces[0]

	//log.Printf("Placing piece %v on board %v", piece, board)

	for y := 0; y <= len(board)-len(piece); y++ {
		for x := 0; x <= len(board[0])-len(piece[0]); x++ {
			// Try all rotations
			for r := 0; r < 4; r++ {
				if nboard, ok := canPlace(board, piece, x, y, r); ok {
					if place(nboard, pieces[1:]) {
						return true
					}
				}
			}
		}
	}

	return false
}

func (s *Server) Day12Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sizes, boards := getSizesAndBoards(req.GetData())

	solved := 0
	for _, board := range boards {
		// Can we just fit each piece on its own?
		sumPieces := 0
		for _, counts := range board.counts {
			sumPieces += counts
		}

		if (board.width/3)*(board.height/3) >= sumPieces {
			log.Printf("%v %v -> %v vs %v", board.width, board.height, (board.width/3)*(board.height/3), sumPieces)
			solved += 1
			continue
		}

		// Are there more hashes than spaces
		totalHashes := 0
		for i, counts := range board.counts {
			totalHashes += counts * sizes[i].hashes
		}

		if totalHashes > board.width*board.height {
			continue
		}

		// Now we do placement
		log.Printf("Trying to places")
		if place(buildBoard(board.width, board.height), buildPieces(sizes, board.counts)) {
			solved += 1
		}
	}

	return &pb.SolveResponse{Answer: int32(solved)}, nil
}
