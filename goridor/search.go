package goridor

import (
	"container/heap"
	"math"
)

type Node struct {
	tile  *Tile
	f     float64
	index int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

func (pq *PriorityQueue) update(node *Node, tile *Tile, f float64) {
	node.tile = tile
	node.f = f
	heap.Fix(pq, node.index)
}

func (pq *PriorityQueue) contains(tile *Tile) bool {
	for _, node := range *pq {
		if node.tile == tile {
			return true
		}
	}
	return false
}

func distance(a *Tile, b *Tile) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)) //math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))
}

func AStar(start *Tile, end *Tile) (bool, []*Tile) { //map[*Tile]*Tile

	openQueue := make(PriorityQueue, 0)
	heap.Init(&openQueue)

	cameFrom := make(map[*Tile]*Tile)
	cameFrom[start] = nil
	gScore := make(map[*Tile]float64)
	gScore[start] = 0

	startNode := &Node{tile: start, f: 0}
	heap.Push(&openQueue, startNode)

	for openQueue.Len() > 0 {

		current := heap.Pop(&openQueue).(*Node)

		if current.tile == end {
			c := end
			var path []*Tile
			for c != start {
				path = append([]*Tile{c}, path...)
				c = cameFrom[c]
			}
			return true, path
		}

		for _, neighbor := range current.tile.neighbor {
			if neighbor == nil {
				continue
			}
			if gScore[neighbor] == 0 {
				gScore[neighbor] = math.Inf(1)
			}
			newGScore := gScore[current.tile] + 1 // d cost

			if newGScore < gScore[neighbor] {
				cameFrom[neighbor] = current.tile
				gScore[neighbor] = newGScore
				if !openQueue.contains(neighbor) {
					heap.Push(&openQueue, &Node{tile: neighbor, f: newGScore + distance(neighbor, end)})
				}
			}
		}
	}
	return false, nil
}
