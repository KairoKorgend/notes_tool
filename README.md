# Notes tool

The NotesTool is a command-line application designed for managing short single-line notes. It allows users to organize their notes into collections, providing essential functionalities such as viewing existing notes, adding new notes and removing existing notes.

### Usage of the tool

The NotesTool allows users to perform the following operations:

- **Create or Manage a Collection:**
  - `$ ./notestool [TAG]`
  - Example: `$ ./notestool coding_ideas`
  
  Creates or loads a collection named "coding_ideas," initiating a menu-based interface for note management.

- **Display Notes from the Collection:**
  - Selecting option `1` from the menu displays the existing notes within the collection.

- **Add a Note to the Collection:**
  - Selecting option `2` from the menu prompts the user to input a new note, which is then added to the collection.

- **Remove a Note from the Collection:**
  - Selecting option `3` from the menu allows the user to remove a specific note from the collection.

- **Exit the Program:**
  - Selecting option `4` exits the program.

### Data Storage

For each collection, the NotesTool creates a separate database in the form of a plain text file with the same name as the collection (e.g., "coding_ideas.txt"). Each note is stored as a separate row within this file. If the specified collection does not exist, a new one is created. Existing collections are loaded to ensure that notes persist between different runs of the tool.

## Examples

#### Create or Manage a Collection

```bash
$ ./notestool coding_ideas
```
This command creates or loads the "coding_ideas" collection and displays the menu.

#### Display Help Message

```bash
$ ./notestool help
Usage: ./notestool [TAG]
```

#### Tool Functionality

```bash
$ ./notestool coding_ideas
Welcome to the notes tool!

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
2

Enter the note text:
note three

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two
003 - note three

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
3

Enter the number of note to remove or 0 to cancel:
3

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
4
