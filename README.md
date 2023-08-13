# Monkey Interpreter in Go

This project is an implementation of the Monkey programming language using Go. It's an experimental project to learn how to build an interpreter for a language and serves as an exercise for the book ["Writing An Interpreter In Go"](https://www.amazon.es/gp/product/3982016118).

## About Monkey Language

Monkey is a programming language designed for learning purposes. It's a simple language with a syntax that is easy to understand.

### Example of Monkey Language

Here is a quick glimpse into what Monkey code looks like:

```monkey
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
let myArray = [1, 2, 3, 4, 5];
let thorsten = {"name": "Thorsten", "age": 28};
myArray[0] // => 1
thorsten["name"] // => "Thorsten"

let add = fn(a, b) { return a + b; };
let add = fn(a, b) { a + b; };
add(1, 2);

let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

let twice = fn(f, x) { 
    return f(f(x));
};

let addTwo = fn(x) { 
    return x + 2;
};

twice(addTwo, 2); // => 6
```

## Building and Running

This section will guide you through the process of building and running the Monkey interpreter.

### Prerequisites

Make sure you have Go installed on your system. You can download it from [here](https://golang.org/).

### Building

Clone the repository:

```bash
git clone https://github.com/hlastras/monkeylang.git
cd monkeylang
```

Build the project:

```bash
go build
```

### Running

Execute the interpreter:

```bash
./monkey
```

You can now enter Monkey code directly into the interpreter!

## Contributing

This is an experimental project, and contributions are welcome. Please feel free to fork the repository, create feature branches, and submit pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

* Inspired by and an exercise of the book ["Writing An Interpreter In Go"](https://www.amazon.es/gp/product/3982016118).

