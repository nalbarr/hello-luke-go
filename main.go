package main

import (
	"fmt"
	"reflect"
)

// NA. Port original Java main()
/*
public class Main {
	public static void main(String args[]) {

		// 1D as array, list or linear algebra vector of scalars
		iterateScalars1DWithIndexes();

		// 1D as array, list or linear algebra vector of scalars
		iterateScalars1DWithCollections();

		// 1D as array, list or linear algebra vector of <Object>
		iterateObjects1DWithIndexes();

		// 1D as collection of interfaces
		iterateInterfaces1DWithCollections();

		// 1D as collection of interfaces with type casting
		iterateInterface1DWithSpecificTypeCasts();

		// 1D as different collection types for different operations (e.g., get)
		collectionsComplexity();

		// NAA. After some 2048 game development some new concepts creeped up.
		// - When to use events and event listeners
		// - Avoiding nested for loops if you can use better data structures (e.g., collections)
		// - Separation of concerns (e.g., view, models, controllers) can you code be flexible for
		// Java Swing, command line. How easy would it be to lift/move to Unity3D?
		// - Starting to think more functionally (e.g., for loops vs. matrices/vectorization)
	}

	public static void iterateScalars1DWithIndexes() {
		final int I_MAX = 5;
		for (int i = 0; i < I_MAX; i++) {
			System.out.println(String.format("i: %d", i));
		}
	}

	// NLA. Read below for reference.
	// i.e, https://stackoverflow.com/questions/24943663/how-to-iterate-through-an-arraylist-of-objects-of-arraylist-of-objects
	public static void iterateScalars1DWithCollections() {
		// CONCEPTS:
		// - What is List? (e.g., class, abstract class, interface)

		// NAA. NOTE: advanced in one line initialization
		//List<Integer> someInts = new ArrayList<Integer>(Arrays.asList(1,2,3,4,5));]
		// NAA. NOTE: below will warn you you need to specify parameterized type
		//List<Integer> someInts = new ArrayList();
		List<Integer> someInts = new ArrayList<Integer>();
		someInts.add(new Integer(1));
		someInts.add(new Integer(2));
		someInts.add(new Integer(3));
		someInts.add(new Integer(4));
		someInts.add(new Integer(5));

		for(Integer anInt: someInts) {
			System.out.println(String.format("anInt: %d", anInt));
		}

	}

	public static void iterateObjects1DWithIndexes() {
		// CONCEPTS:
		// - What is Item? (e.g., class, abstract class, interface)

		List<Item> someItems = new ArrayList<Item>();
		someItems.add(new Item("one", 1));
		someItems.add(new Item("two", 2));
		someItems.add(new Item("three", 3));
		someItems.add(new Item("four", 4));
		someItems.add(new Item("five", 5));

		for(int i=0; i < someItems.size(); i++) {
			System.out.println(String.format("item.getValue(): %d", someItems.get(i).getValue()));
		}
	}

	public static void iterateObjects1DWithCollections() {
		// CONCEPTS:
		// - What is Item? (e.g., class, abstract class, interface)

		List<Item> someItems = new ArrayList<Item>();
		someItems.add(new Item("one", 1));
		someItems.add(new Item("two", 2));
		someItems.add(new Item("three", 3));
		someItems.add(new Item("four", 4));
		someItems.add(new Item("five", 5));

		for(Item item: someItems) {
			System.out.println(String.format("item.getValue(): %d", item.getValue()));
		}
	}

	public static void iterateInterfaces1DWithCollections() {
		// CONCEPTS:
		// - What is BaseItem, IItem? (e.g., class, abstract class, interface)

		List<IItem> someItems = new ArrayList<IItem>();
		someItems.add(new Item("one", 1));
		someItems.add(new Item("two", 2));
		someItems.add(new Item("three", 3));
		someItems.add(new Item("four", 4));
		someItems.add(new Item("five", 5));

		for(IItem item: someItems) {
			System.out.println(String.format("item.getValue(): %d", item.getValue()));
		}
	}

	public static void iterateInterface1DWithSpecificTypeCasts() {
		// CONCEPTS:
		// - Why can we add both Human and KillerWhale instances
		// - Why can we iterate through collection without indexes?
		// - How would you print out name and value? i.e., Human and KillerWhale extend Item super class

		List<IItem> someItems = new ArrayList<IItem>();
		someItems.add(new Human("Luke", 100));
		someItems.add(new KillerWhale("Waffles", 100));

		for(IItem item: someItems) {
			System.out.println(item.getClass().getName());
			// NOTE: not ideal but, showing how to find a class instance type
			if (item instanceof Human) {
				System.out.println(String.format("item.swim(): %d", ((Human) item).walk()));
			} else if (item instanceof KillerWhale) {
				System.out.println(String.format("item.swim(): %d", ((KillerWhale) item).swim()));
			}
		}
	}

	public static void iterateDifferentCollections() {
		// CONCEPTS:
		// - Why do both someItems and otherItems compile?
		// - Why are there warnings
		List<Item> someItems = new ArrayList<Item>();
		List<Item> otherItems = new LinkedList<Item>();
	}

	public static void collectionsComplexity() {
		// CONCEPTS:
		// - ArrayList, LinkedList
		// - NOTE: populating each collection with 10 million integers and attempting to get an integer stored at a random index.
		// - Why does calling get(<index>) result in different duration times to complete (performance)?
		// - Hpw would this influence your design of a program?
		final int MAX_SIZE = 10000000;
		Random rand = new Random();

		List<Integer> intList1 = new ArrayList<Integer>();
		for (int i = 0; i < MAX_SIZE; i++) {
			intList1.add(rand.nextInt(100));
		}

		List<Integer> intList2 = new LinkedList<Integer>();
		for (int i = 0; i < MAX_SIZE; i++) {
			intList2.add(rand.nextInt(100));
		}

		long start = 0L;
		long stop = 0L;
		long duration = 0L;

		int index = rand.nextInt(MAX_SIZE);

		start = System.currentTimeMillis();
		int someInt1 = intList1.get(index);
		stop = System.currentTimeMillis();
		duration = stop - start;
		System.out.println(String.format("intList1.get(%d): %d took %d milliseconds ", index, someInt1, duration));

		start = System.currentTimeMillis();
		int someInt2 = intList2.get(index);
		stop = System.currentTimeMillis();
		duration = stop - start;
		System.out.println(String.format("intList2.get(%d): %d took %d milliseconds ", index, someInt2, duration));
	}
*/

func iterateScalars1DWithIndexes() {
	fmt.Println("*** iterateScalars1DWithIndexes")

	ints := make([]int, 5)
	ints[0] = 1
	ints[1] = 2
	ints[2] = 3
	ints[3] = 4
	ints[4] = 5

	for i := 0; i < len(ints); i++ {
		fmt.Printf("i: %d\n", i)
	}
	fmt.Println("")
}

func iterateScalars1DWithCollections() {
	fmt.Println("*** iterateScalars1DWithCollections")

	ints := [5]int{1, 2, 3, 4, 5}

	for i, e := range ints {
		fmt.Printf("ints[%d]: %d\n", i, e)
	}
	fmt.Println("")
}

// NA. Golang does not support classes !!!
// NA. Golang supports trailing comma for last element in array declaration?
// - https://dave.cheney.net/2014/10/04/that-trailing-comma
func iterateObjects1DWithIndexes() {
	fmt.Println("*** iterateObjects1DWithIndexes")

	type Point struct {
		X int
		Y int
	}
	points := [5]Point{
		Point{1, 1},
		Point{2, 2},
		Point{3, 3},
		Point{4, 4},
		Point{5, 5},
	}

	for i, e := range points {
		fmt.Printf("ints[%d]: %+v\n", i, e)
	}
	fmt.Println("")
}

// Adder ...
type Adder interface {
	Add() int
}

// Point2D ...
type Point2D struct {
	X int
	Y int
}

// Point3D ...
type Point3D struct {
	X int
	Y int
	Z int
}

// Add ...
func (p Point2D) Add() int {
	return p.X + p.Y
}

// Add ...
func (p Point3D) Add() int {
	return p.X + p.Y + p.Z
}
func iterateInterfaces1DWithCollections() {
	fmt.Println("*** iterateInterfaces1DWithCollections")

	point2Ds := [5]Point2D{
		Point2D{1, 1},
		Point2D{2, 2},
		Point2D{3, 3},
		Point2D{4, 4},
		Point2D{5, 5},
	}
	point3Ds := [5]Point3D{
		Point3D{1, 1, 1},
		Point3D{2, 2, 2},
		Point3D{3, 3, 3},
		Point3D{4, 4, 4},
		Point3D{5, 5, 5},
	}

	for i, e := range point2Ds {
		fmt.Printf("ints[%d]: %+v\n", i, e)
	}
	fmt.Println("")

	for i, e := range point3Ds {
		fmt.Printf("ints[%d]: %+v\n", i, e)
	}
	fmt.Println("")
}

func iterateInterface1DWithSpecificTypeCasts() {
	fmt.Println("*** iterateInterface1DWithSpecificTypeCasts")

	point2Dand3Ds := [5]Adder{
		Point2D{1, 1},
		Point3D{2, 2, 2},
		Point2D{3, 3},
		Point3D{4, 4, 4},
		Point2D{5, 5},
	}
	for i, e := range point2Dand3Ds {
		fmt.Printf("ints[%d]: %+v is a %s\n", i, e, reflect.TypeOf(e))
	}
	fmt.Println("")
}

func collectionsComplexity() {
	fmt.Println("*** collectionsComplexity")
}

func main() {

	// 1D as array, list or linear algebra vector of scalars
	iterateScalars1DWithIndexes()

	// 1D as array, list or linear algebra vector of scalars
	iterateScalars1DWithCollections()

	// 1D as array, list or linear algebra vector of <Object>
	iterateObjects1DWithIndexes()

	// 1D as collection of interfaces
	iterateInterfaces1DWithCollections()

	// 1D as collection of interfaces with type casting
	iterateInterface1DWithSpecificTypeCasts()

	// 1D as different collection types for different operations (e.g., get)
	collectionsComplexity()
}
