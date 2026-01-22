# River Crossing Puzzle 🚣

A command-line implementation of the classic river crossing logic puzzle in Go. Help the farmer (HS) transport a chicken, fox, and corn across the river without leaving incompatible items together.

## The Puzzle

The classic river crossing puzzle constraints:
- The farmer can only take one item in the boat at a time
- The chicken cannot be left alone with the corn (it will eat it)
- The fox cannot be left alone with the chicken (it will eat it)
- The farmer must be present to operate the boat

## Technical Implementation

This project demonstrates clean Go architecture with:
- **Separation of concerns** using multiple packages
- **State management** pattern for tracking puzzle state
- **Event-driven design** for handling moves
- **Test-driven development** with unit tests

```

## Getting Started

### Prerequisites

- Go 1.13 or higher

### Installation

```bash
git clone https://github.com/yourusername/rivercrossing.git
cd rivercrossing
go build
```

### Running the Program

```bash
go run main.go
```

## Commands
```
status
```
Displays the current state of all items (boat, west bank, east bank)

### Movement Commands

**Chicken:**
- `PutChickenBoat` - Put chicken in boat
- `PutChickenWest` - Move chicken to west bank
- `PutChickenEast` - Move chicken to east bank

**Fox (Rev):**
- `PutFoxBoat` - Put fox in boat
- `PutFoxWest` - Move fox to west bank
- `PutFoxEast` - Move fox to east bank

**Corn (Korn):**
- `PutCornBoat` - Put corn in boat
- `PutCornWest` - Move corn to west bank
- `PutCornEast` - Move corn to east bank

**Farmer (HS - Human/Menneske):**
- `PutHsBoat` - Put farmer in boat
- `PutHsWest` - Move farmer to west bank
- `PutHsEast` - Move farmer to east bank

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

## Key Features

- **State Tracking**: Real-time tracking of all entities across three locations (west bank, boat, east bank)
- **Boolean State Management**: Each entity has three boolean flags representing its possible locations
- **Immutable State Transitions**: State changes are handled through dedicated event functions
- **Interactive CLI**: Simple command-line interface for puzzle interaction
- **Extensible Design**: Easy to add validation rules or additional entities

## Learning Outcomes

This project demonstrates:
- Go package organization and module management
- State pattern implementation
- Command pattern for user actions
- Unit testing in Go
- Clean code principles and separation of concerns

## 👤 Author

**Robin**
- GitHub: [@robinfins](https://github.com/robinfins)
