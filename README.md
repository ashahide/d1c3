# d1c3

`d1c3` is a lightweight command-line tool written in Go for rolling dice, with support for *advantage*, *disadvantage*, and optional *verbose output* designed for tabletop gaming.

---

## Features

- Roll custom dice expressions like `2d6+1d8-3`
- Supports **advantage** and **disadvantage**
- Verbose mode explains individual dice rolls
- Gracefully handles combined expressions like `2d6+3d4-1`

---

## Installation

```bash
go install github.com/ashahide/d1c3@latest
```

Ensure that your `$GOPATH/bin` is in your `PATH`.

---

## Usage

```bash
d1c3 [dice_expression] [flags]
```

### Dice Expression Format

The expression must be passed as a **single string**, either:
- Without spaces:  
  ```bash
  d1c3 2d4+1d6-3
  ```
- Or quoted if including spaces:
  ```bash
  d1c3 '2d4 + 1d6 - 3'
  ```

---

## Flags

| Flag             | Description                                           |
|------------------|-------------------------------------------------------|
| `--advantage`    | Roll two dice and keep the higher result              |
| `--disadvantage` | Roll two dice and keep the lower result               |
| `--verbose`      | Print full breakdown of each roll and subtotal        |

> If both `--advantage` and `--disadvantage` are set, both will be ignored.

---

## Examples

### Basic roll

```bash
$ d1c3 1d5
2
```

### Verbose output

```bash
$ d1c3 1d5 -verbose
Dice Rolls:
  Added      1d5 -> [2] -> subtotal: 2
Total: 2
```

### Multi-term roll with subtraction

```bash
$ d1c3 2d6+3d5-7 -verbose
Dice Rolls:
  Added      2d6 -> [1 6] -> subtotal: 7
  Added      3d5 -> [5 1 5] -> subtotal: 11
  Subtracted 7d1 -> [1 1 1 1 1 1 1] -> subtotal: 7
Total: 11
```

---

## Output Behavior

- If `--verbose` is **not used**, only the final total is printed.
- If `--verbose` is used, each step of the roll is printed in a readable format showing:
  - the operation (`added` / `subtracted`)
  - the dice expression
  - individual roll results
  - subtotal per roll group

---

## Project Structure

- `main.go` – CLI argument parsing and execution
- `internal/roll` – Dice parsing and rolling logic
- `internal/logtools` – Logger initialization and debug output

---

## License

MIT License. See `LICENSE` for full details.

---

## Author

Developed by [@ashahide](https://github.com/ashahide)