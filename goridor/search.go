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
	return math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))
}

func reconstructPath(cameFrom map[*Tile]*Tile, current *Tile) []*Tile {
	fullPath := []*Tile{current}
	for key, _ := range cameFrom {
		c := cameFrom[key]
		fullPath = append(fullPath, c)
	}
	return fullPath
}

func AStar(start *Tile, end *Tile) []*Tile {

	openQueue := make(PriorityQueue, 0)
	heap.Init(&openQueue)

	cameFrom := make(map[*Tile]*Tile)
	gScore := make(map[*Tile]float64)
	gScore[start] = 0

	startNode := &Node{tile: start, f: distance(start, end)}
	heap.Push(&openQueue, startNode)

	for openQueue.Len() > 0 {

		current := heap.Pop(&openQueue).(*Node)

		if current.tile == end {
			println("Found")
			return reconstructPath(cameFrom, current.tile)
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
	return []*Tile{}
}

/*
// TODO: optimize for searching all goal locations
func AStarSearch(startTile *Tile, endTile *Tile) []*Tile {

	var closedList []*Tile
	openQueue := make(PriorityQueue, 0)
	heap.Init(&openQueue)

	startNode := &Node{tile: startTile, f: 0}

	heap.Push(&openQueue, startNode)
	//openQueue.update(startNode, startNode.tile, 0)

	var current *Node

	SEARCHSTART:
	for openQueue.Len() > 0 {

		current = heap.Pop(&openQueue).(*Node)

		if current.tile == endTile {
			println("Found")
			return []*Tile{current.tile}
		}

		for _, tile := range current.tile.neighbor {
			if tile == nil {
				continue
			}
			if slices.Contains(closedList, tile) {
				continue
			}
			g := current.g + 1 // Distance between
			h := 1 // Distance to end
			f := g + h

			if openQueue.contains(tile) {
				if g >

				continue SEARCHSTART
			} else {
				heap.Push(&openQueue, &Node{tile: tile, f: f})
			}
		}

	}
	return []*Tile{}
}
*/
/**
// A* (star) Pathfinding
// Initialize both open and closed list
let the openList equal empty list of nodes
let the closedList equal empty list of nodes
// Add the start node
put the startNode on the openList (leave it's f at zero)
// Loop until you find the end
while the openList is not empty
    // Get the current node
    let the currentNode equal the node with the least f value
    remove the currentNode from the openList
    add the currentNode to the closedList
    // Found the goal
    if currentNode is the goal
        Congratz! You've found the end! Backtrack to get path
    // Generate children
    let the children of the currentNode equal the adjacent nodes

    for each child in the children
        // Child is on the closedList
        if child is in the closedList
            continue to beginning of for loop
        // Create the f, g, and h values
        child.g = currentNode.g + distance between child and current
        child.h = distance from child to end
        child.f = child.g + child.h
        // Child is already in openList
        if child.position is in the openList's nodes positions
            if the child.g is higher than the openList node's g
                continue to beginning of for loop
        // Add the child to the openList
        add the child to the openList

*/
