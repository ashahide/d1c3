# d1c3

`d1c3` is a lightweight command-line tool written in Go for rolling dice, with support for *advantage* and *disadvantage* mechanics (such as in tabletop RPGs like Dungeons & Dragons).

---

## Features

- Roll custom dice expressions like `2d6+1d8-3`
- Support for advantage (`--advantage`) and disadvantage (`--disadvantage`)
- Input parsing with detailed logging

---

## Installation

```bash
go install github.com/ashahide/d1c3@latest
```

Ensure that your `$GOPATH/bin` directory is in your system's `PATH`.

---

## Usage

```bash
d1c3 [dice_expression] [flags]
```

### Dice Expression Format

The dice expression must be passed as a **single string**, either:
- Without spaces:  
  ```bash
  d1c3 2d4+1d6-3
  ```
- Or quoted to preserve spaces:
  ```bash
  d1c3 '2d4 + 1d6 - 3'
  ```

Unquoted expressions with spaces will not be parsed correctly unless there are no whitespaces
like 2d4-6.

---

## Flags

| Flag             | Description                                           |
|------------------|-------------------------------------------------------|
| `--advantage`    | Roll two dice for each roll and keep the higher value |
| `--disadvantage` | Roll two dice for each roll and keep the lower value  |

If both `--advantage` and `--disadvantage` are provided, neither will be applied.

---

## Examples

```bash
d1c3 1d20
d1c3 3d6+2
d1c3 '2d10 + 1d4' --advantage
d1c3 '1d20 + 5' --disadvantage
```

---

## Example Output

```bash
$ go run cmd/main.go '2d6 + 2d3 - 7' --disadvantage
Dice Rolls (with disadvantage):  [{+ 2d6 [4 2] 6} {+ 2d3 [1 1] 2} {- 7d1 [1 1 1 1 1 1 1] 7}]
Total:  1

$ go run cmd/main.go '2d6 + 2d3 - 7'
Dice Rolls:  [{+ 2d6 [1 2] 3} {+ 2d3 [3 3] 6} {- 7d1 [1 1 1 1 1 1 1] 7}]
Total:  2

$ go run cmd/main.go '2d6 + 2d3'
Dice Rolls:  [{+ 2d6 [2 4] 6} {+ 2d3 [2 1] 3}]
Total:  9

$ go run cmd/main.go "2d4"
Dice Rolls:  [{+ 2d4 [1 3] 4}]
Total:  4

$ go run cmd/main.go "2d4-6"
Dice Rolls:  [{+ 2d4 [3 4] 7} {- 6d1 [1 1 1 1 1 1] 6}]
Total:  1

$ go run cmd/main.go "2d4" --advantage
Dice Rolls (with advantage):  [{+ 2d4 [4 3] 7}]
Total:  7

$ go run cmd/main.go "2d4" --disadvantage
Dice Rolls (with disadvantage):  [{+ 2d4 [4 1] 5}]
Total:  5
```

---

## Project Structure

- `main.go` – CLI argument parsing and execution
- `internal/roll` – Dice parsing and rolling logic
- `internal/logtools` – Logger initialization and debug output

---

## Logging

Logs include detailed information about:
- Argument parsing
- Flag interpretation
- Roll breakdowns
- Final totals

This helps with debugging or inspecting how the dice rolls were computed.

---

## License

MIT License. See `LICENSE` for full details.

---

## Author

Developed by [@ashahide](https://github.com/ashahide)