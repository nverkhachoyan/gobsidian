```dataview
TABLE WITHOUT ID link(file.name) AS "Town Name", type, Alignment, population
where contains(file.path, this.file.folder) AND file.name != this.file.name AND file.name != "Attachments" AND file.name != "_attachments"
where (NoteIcon="settlement")
```











