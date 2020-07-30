# GoLand Tips & Tricks

This repository is a collection of tips & tricks for GoLand that can be tried any time by users.

Feel free to contribute or open issues for various features that you'd like to see added.

## How to use?

Each file contains a small description of what to do.

There are two types of important comments:

- `// Step X.` -> this allows you to perform actions in a certain order
- `// E.g.` -> this allows you to know what to do at that specific point 

## List of Tips&Tricks

The tips and tricks can be found divided by a few major categories:
- [Completion](#completion)
- [Editing](#editing)
- [Go Modules](#go-modules)
- [Navigation](#navigation)
- [Refactoring](refactoring)
- [Running, testing and debugging](#running-testing-and-debugging)

### Completion

| Tip number | Contents | New in |
|---|----|-----|
| [001](completion/tip001) | Import a package without typing its name | |
| [002](completion/tip002) | Smart Type Completion | |
| [003](completion/tip003) | Postfix Completion | |
| [004](completion/tip004) | Method-like Completion | |
| [005](completion/tip005) | (Custom) Live Templates. Use of builtin completion templates helpers | |
| [006](completion/tip006) | Partial Match Completion | |
| [007](completion/tip007) | Completion with Tab | |
| [008](completion/tip008) | Completion for type-assertion | |
| [009](completion/tip009) | Parameter name auto-generation | |


### Editing

| Tip number | Contents | New in |
|---|----|-----|
| [001](editing/tip001) | Cyclic Expand Word / Cyclic Expand Word (Backward) | |
| [002](editing/tip002) | Parameter Info for functions and structs ||
| [003](editing/tip003) | Language Injections ||
| [004](editing/tip004) | Go Templates support ||
| [005](editing/tip005) | Create undefined type. Use multi-cursor for struct tag Live Template. ||
| [006](editing/tip006) | Extend/Shrink Selection ||
| [007](editing/tip007) | Add Selection for Next Occurrence ||
| [008](editing/tip008) | Completion in comments ||
| [009](editing/tip009) | Function/struct literals wrapping/chopping styles | 2020.2 |

### Inspections
| Tip number | Contents | New in |
|---|----|-----|
| [001](inspections/tip001) | Detect and navigate to duplicate tags of fields in a structure | 2020.2 |
| [001](inspections/tip001) | Inspection for int to string conversion (e.g. `string(int)`) | 2020.2 |

### Go Modules
| Tip number | Contents | New in |
|---|----|-----|
| [001](modules/tip001) | Go mod completions support. Inspection from the IDE for local paths. | |

### Navigation
| Tip number | Contents | New in |
|---|----|-----|
| [001](navigation/tip001) | Navigate to File. Switcher. Recent Files. Recent Locations. | |
| [002](navigation/tip002) | Navigate to File. Structure Pop-up. Select in. | |
| [003](navigation/tip003) | Type Hierarchy. Call Hierarchy | |
| [004](navigation/tip004) | Navigate to/from symbols referenced in documentation comments | 2020.2 |

### Refactoring
| Tip number | Contents | New in |
|---|----|-----|
| [001](refactoring/tip001) | Implement Interface | |
| [002](refactoring/tip002) | Change Signature refactoring. Move refactoring | |
| [003](refactoring/tip003) | Extract Interface refactoring | |
| [004](refactoring/tip004) | Introduce Constant/Variable refactoring. Inline refactoring | |

### Running, testing and debugging
| Tip number | Contents | New in |
|---|----|-----|
| [001](run-test-debug/tip001) | Completion in breakpoints | |
| [002](run-test-debug/tip002) | Generate test for function. Use auto-test for testing changes. | |
| [003](run-test-debug/tip003) | Debugger Labels | |
| [004](run-test-debug/tip004) | Benchmark support | |
| [005](run-test-debug/tip005) | Smart Step Into | |

## Setup the IDE for tips/comments/shortcut highlighting

To see the shortcuts in a nicer manner, setup the following TODOs:
- `// E.g. ` -> `\b(E\.g\. )\b.*`, Case Sensitive, Color #35C03B
- `// Step x. ` -> `\b(Step \d\. )\b.*`, Case Sensitive, Color #35C03B
- `// Shortcut: ` -> `\b(Shortcut: )\b.*`, Case Sensitive, Color #2B80C0

## Contributions welcomed

Do you have a favorite tip or trick, and you want to see it here?

Send a PR at any time, and we'll add it!

## License

Apache 2, see the [license](LICENSE.md).