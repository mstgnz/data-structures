package graph

import (
	"fmt"

	graph "github.com/mstgnz/data-structures/Graph"
)

// RunExamples demonstrates various graph algorithms
func RunExamples() {
	// Example 1: Topological Sort
	fmt.Println("Topological Sort Example:")
	g1 := graph.NewGraph(6, true)
	g1.AddEdge(5, 2, 1)
	g1.AddEdge(5, 0, 1)
	g1.AddEdge(4, 0, 1)
	g1.AddEdge(4, 1, 1)
	g1.AddEdge(2, 3, 1)
	g1.AddEdge(3, 1, 1)

	ts := graph.NewTopologicalSort(g1)
	order := ts.Sort()
	fmt.Printf("Topological Sort: %v\n\n", order)

	// Example 2: Strongly Connected Components (Tarjan)
	fmt.Println("Strongly Connected Components Example:")
	g2 := graph.NewGraph(5, true)
	g2.AddEdge(1, 0, 1)
	g2.AddEdge(0, 2, 1)
	g2.AddEdge(2, 1, 1)
	g2.AddEdge(0, 3, 1)
	g2.AddEdge(3, 4, 1)

	tarjan := graph.NewTarjanSCC(g2)
	components := tarjan.FindComponents()
	fmt.Printf("Strongly Connected Components: %v\n\n", components)

	// Example 3: Articulation Points and Bridges
	fmt.Println("Articulation Points and Bridges Example:")
	g3 := graph.NewGraph(5, false)
	g3.AddEdge(1, 0, 1)
	g3.AddEdge(0, 2, 1)
	g3.AddEdge(2, 1, 1)
	g3.AddEdge(0, 3, 1)
	g3.AddEdge(3, 4, 1)

	ap := graph.NewArticulationPoints(g3)
	points := ap.FindArticulationPoints()
	bridges := ap.FindBridges()
	fmt.Printf("Articulation Points: %v\n", points)
	fmt.Printf("Bridges: %v\n\n", bridges)

	// Example 4: Euler Path and Circuit
	fmt.Println("Euler Path and Circuit Example:")
	g4 := graph.NewGraph(4, false)
	g4.AddEdge(0, 1, 1)
	g4.AddEdge(1, 2, 1)
	g4.AddEdge(2, 3, 1)
	g4.AddEdge(3, 0, 1)

	ep := graph.NewEulerPath(g4)
	if ep.HasEulerCircuit() {
		circuit := ep.FindEulerCircuit()
		fmt.Printf("Euler Circuit: %v\n", circuit)
	} else if ep.HasEulerPath() {
		path := ep.FindEulerPath()
		fmt.Printf("Euler Path: %v\n", path)
	} else {
		fmt.Println("Euler path or circuit not found")
	}
	fmt.Println()

	// Example 5: Hamiltonian Path and Circuit
	fmt.Println("Hamiltonian Path and Circuit Example:")
	g5 := graph.NewGraph(4, false)
	g5.AddEdge(0, 1, 1)
	g5.AddEdge(1, 2, 1)
	g5.AddEdge(2, 3, 1)
	g5.AddEdge(3, 0, 1)

	hp := graph.NewHamiltonianPath(g5)
	if hp.HasHamiltonianCircuit() {
		circuit := hp.FindHamiltonianCircuit()
		fmt.Printf("Hamiltonian Circuit: %v\n", circuit)
	} else if hp.HasHamiltonianPath() {
		path := hp.FindHamiltonianPath()
		fmt.Printf("Hamiltonian Path: %v\n", path)
	} else {
		fmt.Println("Hamiltonian path or circuit not found")
	}
	fmt.Println()

	// Example 6: Shortest Path Algorithms
	fmt.Println("Shortest Path Algorithms Example:")
	g6 := graph.NewGraph(5, true)
	g6.AddEdge(0, 1, 4)
	g6.AddEdge(0, 2, 2)
	g6.AddEdge(1, 2, 3)
	g6.AddEdge(1, 3, 2)
	g6.AddEdge(1, 4, 3)
	g6.AddEdge(2, 1, 1)
	g6.AddEdge(2, 3, 4)
	g6.AddEdge(2, 4, 5)
	g6.AddEdge(4, 3, 1)

	// Bellman-Ford
	bf := graph.NewBellmanFord(g6, 0)
	if bf.ComputeShortestPaths() {
		fmt.Println("Bellman-Ford Results:")
		fmt.Printf("Distances: %v\n", bf.GetAllDistances())
	}

	// Floyd-Warshall
	fw := graph.NewFloydWarshall(g6)
	fw.ComputeShortestPaths()
	fmt.Println("\nFloyd-Warshall Results:")
	fmt.Printf("Distance Matrix: %v\n\n", fw.GetAllPairsDistances())

	// Example 7: Minimum Spanning Tree
	fmt.Println("Minimum Spanning Tree Example:")
	g7 := graph.NewGraph(5, false)
	g7.AddEdge(0, 1, 2)
	g7.AddEdge(0, 3, 6)
	g7.AddEdge(1, 2, 3)
	g7.AddEdge(1, 3, 8)
	g7.AddEdge(1, 4, 5)
	g7.AddEdge(2, 4, 7)
	g7.AddEdge(3, 4, 9)

	// Prim
	prim := graph.NewPrimMST(g7)
	if prim.FindMST() {
		fmt.Println("Prim MST Results:")
		fmt.Printf("MST Edges: %v\n", prim.GetMSTEdges())
		fmt.Printf("MST Cost: %v\n", prim.GetMSTCost())
	}

	// Kruskal
	kruskal := graph.NewKruskalMST(g7)
	if kruskal.FindMST() {
		fmt.Println("\nKruskal MST Results:")
		fmt.Printf("MST Edges: %v\n", kruskal.GetMSTEdges())
		fmt.Printf("MST Cost: %v\n", kruskal.GetMSTCost())
	}
}
