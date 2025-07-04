`BUTTON[button_quest]`

```dataview
TABLE WITHOUT ID link(file.name) AS "Quest Name", questStatus AS "Status", questGiver AS "Quest Giver", questLocationObtained AS "Location", questSessionObtained AS "Session", questLootAvail AS "Available Rewards", questLookEarned AS "Acquired Rewards"
from "2-Campaign" AND #quest
where questStatus = "In Progress"
```

<br>

# Completed


```dataview
TABLE WITHOUT ID link(file.name) AS "Quest Name", questStatus AS "Status", questGiver AS "Quest Giver", questLocationObtained AS "Location", questSessionObtained AS "Session", questLootAvail AS "Available Rewards", questLookEarned AS "Acquired Rewards"
from "2-Campaign" AND #quest
where questStatus = "Complete"
```

# Not Started


```dataview
TABLE WITHOUT ID link(file.name) AS "Quest Name", questStatus AS "Status", questGiver AS "Quest Giver", questLocationObtained AS "Location", questSessionObtained AS "Session", questLootAvail AS "Available Rewards", questLookEarned AS "Acquired Rewards"
from "2-Campaign" AND #quest
where questStatus = "Not Started"
```