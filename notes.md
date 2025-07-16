# Advanced Go

## Duration
- 3 Days

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
| What | When |
| ----- | ---- |
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins) |
| Lunch Break | 12:30 PM (1 hour) |
| Tea Break | 3:00 PM (20 mins) |
| Wind up | 5:00 PM |

## Methodology
- No powerpoint
- Code & Discussion

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code (or any other editor)

## Repository
- https://github.com/tkmagesh/cisco-advgo-jul-2025

## Prerequisites
- Data Types
- Variabales, Constants, iota
- if else, switch case, for
- Functions
    - Named results
    - Anonymous Functions
    - Higher Order Functions
        - Function types
        - Passing functions as arguments
        - Returning functions as return values
    - Error Handling
    - Panic & Recovery
    - Structs & Methods
    - Struct Composition
    - Interfaces
    - Modules & Packages

## Coverage
- Recap
    - iota
    - Higher Order Functions
    - Panic & Recovery
    - Interfaces
- Concurrency 
- Http Services
- GRPC
- Testing
- Profiling

## Recap
### iota
    - Auto generation of constants
### Higher Order Functions
- Can also treat functions as `data`
    - Assign functions as values to variables
    - Pass functions as arguments
    - Return functions as return values

### Panic & Recovery
#### Panic
- State of the application where the application execution cannot proceed further
- Use `panic()` function to raise application specific panics
- Typically, the runtime attempts to shutdown the app when a panic occurs AFTER executing all the deferred functions
#### Recover
- Use `recover()` function to recover from the panic (to prevent the application from shutting down)
- Apt to use the `recover()` in `deferred` functions

### Interface

