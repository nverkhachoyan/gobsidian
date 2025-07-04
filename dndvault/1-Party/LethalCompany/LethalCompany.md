```dataview
TABLE WITHOUT ID link(file.name) AS "Character Name", Player, Class, Race, level, Role
from "1-Party/LethalCompany"
where (Role = "Player") 
where (Status = "Active") 
```

<br>

```dataview
TABLE WITHOUT ID link(file.name) AS "Character Name", Player, hp, ac, modifier, pasperc As "Passive Perception (WIS)"
from "1-Party/LethalCompany"
where (Role = "Player") 
where (Status = "Active") 
```