```dataview
TABLE WITHOUT ID link(file.name) AS "POI Name", type
where contains(file.path, this.file.folder) AND file.name != this.file.name AND file.name != "Attachments" AND file.name != "_attachments"
where (NoteIcon="location")
```


