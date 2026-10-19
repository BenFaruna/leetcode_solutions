# Min Stack

## Problem Description

Design a stack data structure that supports retrieving the minimum element in constant time, \(O(1)\), while maintaining the standard stack operations.

A regular stack follows the **Last In, First Out (LIFO)** principle, meaning the last element added is the first one removed. However, finding the minimum element in a regular stack may require traversing all its elements.

The goal is to implement a stack that supports `push()`, `pop()`, `top()`, and `min()` efficiently, ensuring that **every operation takes \(O(1)\) time complexity**.

## Requirements

Implement a stack with the following operations:

- **`push(val)`**: Push an integer `val` onto the top of the stack.
- **`pop()`**: Remove the element at the top of the stack.
- **`top()`**: Return the element at the top of the stack without removing it.
- **`min()`**: Return the minimum element currently in the stack.

## Example

**Input**

```text
push(-2)
push(0)
push(-3)
min()
pop()
top()
min()
```

**Output**

```text
-3
0
-2
```

**Explanation**

1. Push `-2`, `0`, and `-3` onto the stack.
2. `min()` returns `-3`, the smallest element.
3. `pop()` removes `-3` from the stack.
4. `top()` returns `0`, the current top element.
5. `min()` returns `-2`, the smallest remaining element.

## Constraints

- The stack may contain positive, negative, or zero-valued integers.
- Duplicate values must be handled correctly.
- Minimum values must be updated correctly after elements are removed.
- The stack operations are assumed to be called according to their valid usage, such as calling `top()` or `min()` only when the stack is non-empty.

## Complexity Requirements

Each operation must run in constant time, regardless of the number of elements in the stack.

| Operation | Required Time Complexity |
|---|---|
| `push(val)` | \(O(1)\) |
| `pop()` | \(O(1)\) |
| `top()` | \(O(1)\) |
| `min()` | \(O(1)\) |

The implementation should also aim for \(O(n)\) auxiliary space complexity, where \(n\) is the number of elements stored in the stack.

## Challenge

Implement the Min Stack in your preferred programming language without using AI to develop the solution.

Think carefully about how to maintain access to the minimum element as values are pushed onto and removed from the stack, without scanning the entire stack on every `min()` call.

## Source

Inspired by [Chinedu's post on X](https://x.com/chinedu_10/status/2102498682482336171?s=20).

