# GoLand Tips & Tricks

This repository is a collection of tips & tricks for GoLand that can be tried any time by users.

Feel free to contribute or open issues for various features that you'd like to see added.

## How to use?

Each file contains a small description of what to do.

There are two types of important comments:

- `// Step X.` -> this allows you to perform actions in a certain order
- `// E.g.` -> this allows you to know what to do at that specific point 

## List of Tips&Tricks

The tips and tricks are split into a few major categories:
- [Completion](#completion)
- [Editing](#editing)
- [Go Modules](#go-modules)
- [Navigation](#navigation)
- [Refactoring](refactoring)
- [Running, testing and debugging](#running-testing-and-debugging)

### Completion

| Tip number | Contents |
|---|----|
| 001 | Import a package without typing its name |
| 002 | Smart Type Completion |
| 003 | Postfix Completion |
| 004 | Method-like Completion |
| 005 | (Custom) Live Templates. Use of builtin completion templates helpers |
| 006 | Partial Match Completion |
| 007 | Completion with Tab |
| 008 | Completion for type-assertion |
| 009 | Parameter name auto-generation |


### Editing

| Tip number | Contents |
|---|----|
| 001 | Cyclic Expand Word / Cyclic Expand Word (Backward) |
| 002 | Parameter Info for functions and structs |
| 003 | Language Injections |
| 004 | Go Templates support |
| 005 | Create undefined type. Use multi-cursor for struct tag Live Template. |
| 006 | Extend/Shrink Selection |
| 007 | Add Selection for Next Occurrence |
| 008 | Completion in comments |

### Go Modules
| Tip number | Contents |
|---|----|
| 001 | Go mod completions support. Inspection from the IDE for local paths. |

### Navigation
| Tip number | Contents |
|---|----|
| 001 | Navigate to File. Switcher. Recent Files. Recent Locations. |
| 002 | Navigate to File. Structure Pop-up. Select in. |
| 003 | Type Hierarchy. Call Hierarchy |

### Refactoring
| Tip number | Contents |
|---|----|
| 001 | Implement Interface |
| 002 | Change Signature refactoring. Move refactoring |
| 003 | Extract Interface refactoring |
| 004 | Introduce Constant/Variable refactoring. Inline refactoring |

### Running, testing and debugging
| Tip number | Contents |
|---|----|
| 001 | Completion in breakpoints |
| 002 | Generate test for function. Use auto-test for testing changes. |
| 003 | Debugger Labels |
| 004 | Benchmark support |
| 005 | Smart Step Into |

## Setup the IDE for tips/comments/shortcut highlighting

To see the shortcuts in a nice manner, setup the following TODOs:
- `// E.g. ` -> `\b(E\.g\. )\b.*`, Case Sensitive, Color #35C03B
- `// Step x. ` -> `\b(Step \d\. )\b.*`, Case Sensitive, Color #35C03B
- `// Shortcut: ` -> `\b(Shortcut: )\b.*`, Case Sensitive, Color #2B80C0

## Contributions welcomed

Do you have a favorite tip or trick, and you want to see it here?

Send a PR at any time, and we'll add it!

## License

Apache 2, see the [license](LICENSE.md).