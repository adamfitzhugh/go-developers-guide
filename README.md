# go-developers-guide

Go: The Complete Developers Guide Udemy course

## Variables
* := syntax relies on the Go compiler to work out that the variable 'card' is going to be of type string
* := is only used when you are declaring a NEW variable. * You do not need the colon when assigning a new value to that variable

```
  card := newCard
```

## Arrays/Slices
* An array is a fixed list of values, a slice is a varied list of values.
* In order to declare a slice of string use the [] followed by the type and then the values within the {}:

```
cards := []string{"Ace of Diamonds", newCard()}
```


* To declare an array, you can add a value to the [] such as [5]


* In order to append to a list, you can define the new variable as below. This will take the existing list, and add "Six of Spades" to it:

```
cards = append(cards, "Six of Spades")
```


* To iterate over the above list, you can create a for loop like below. This will loop over 'cards' and print out each card value:

```
	for i, card := range cards {
		fmt.Println(i, card)
	}
```


## Custom Type Declarations
* To declare a new special type. The below code declares a slice of strings:

```
type deck []string
```

* To add the above 'deck' type to a variable, you can assign it as:

```
cards := deck{"Ace of Diamonds", newCard()}
```


## Reciever Functions
* Any variable of type deck now gets access to the print method
* The reciever sets up methods on variables that we create
* An example of a reciever function:

```
  func (d deck) print() {
  	for i, card := range d {
  		fmt.Println(i, card)
  	}
  }
```

* Back in the main.go file, a 'cards.print()' method is called. 'cards' correlates to the 'd' in the reciever function