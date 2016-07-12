# Hoard

A trivial commandline task list written in Go

---

I wrote Hoard as an exercise in Go. It's a trivial program that does four things: `add` a task, `finish` a task, `remove` a task, or `view` your current tasks. Also I needed a small task log which I can invoke from the commandline -- and what's better than to write my own?

Sure, it's not the *best* example of a Go app but it's my intention to tweak this every now and then when I come up with new features and to possibly improve my Go coding. ;)

---

## Installation

```bash
go get github.com/kixpanganiban/hoard
```

Run it with `$GOPATH/bin/hoard` or even better, symlink it to your local bin folder to run it from anywhere:

```bash
ln -s $GOPATH/bin/hoard /usr/local/bin/hoard
```

---

## Usage

`hoard [--action=<action>] [task1, task2...]`

Where `<action>` is `add`, `remove`, `finish`, or `view`.

If `add`, `task1`...`taskn` are strings to be added as individual tasks.

If `remove` or `finish`, `task1`...`taskn` are integers (index of the task) to be removed or marked finished.

`view` displays all your current tasks in the format `<index> <finished> <task>`.

---

## TODO
- [ ] Refactor error handling
- [ ] Add ability to export tasks to CSV
- [ ] Add task timestamps
- [ ] Add task tags
- [ ] Make tasks queryable
    - [ ] By partial string
    - [ ] By tags
    - [ ] By timestamp?
- [ ] Add optional colored output


---

## LICENSE

This is licensed under the [MIT License](https://opensource.org/licenses/MIT).

The MIT License (MIT)

Copyright (c) 2016 Kix Panganiban

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
