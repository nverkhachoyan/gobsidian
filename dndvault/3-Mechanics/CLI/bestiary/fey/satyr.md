---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/environment/forest
- monster/size/medium
- monster/type/fey
aliases: ["Satyr"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Monster Manual p. 267. Available in the SRD and the Basic Rules.
---
# [Satyr](3-Mechanics\CLI\bestiary\fey/satyr.md)
*Source: Monster Manual p. 267. Available in the SRD and the Basic Rules.*  

Satyrs are raucous fey that frolic in wild forests, driven by curiosity and hedonism in equal measure.

Satyrs resemble stout male humans with the furry lower bodies and cloven hooves of goats. Horns sprout from their heads, ranging in shape from a pair of small nubs to large, curling rams' horns. They typically sport facial hair.

## Hedonistic Revelers

Satyrs crave the strongest drink, the most fragrant spices, and the most dizzying dances. A satyr feels starved when it can't indulge itself, and it goes to great lengths to sate its desires. It might kidnap a fine minstrel to hear lovely songs, sneak through a well-defended garden to gaze upon a beautiful lad or lass, or infiltrate a palace to taste the finest food in the land. Satyrs allow no festivity to pass them by. They partake in any holiday they've heard of. Civilizations of the world have enough festivals and holy days among them to justify nonstop celebration.

Inebriated on drink and pleasure, satyrs give no thought to the consequences of the hedonism they incite in others. They leave such creatures mystified at their own behavior. Such revelers might have to scrounge for excuses to explain their disordered state to parents, employers, family, or friends.

> [!quote]- A quote from Sonnet of a Naughty Satyr  
> 
> Twixt day and night the spirits goad me on
> 
> They pine for ages past when hearts were pure
> 
> Against all reason now they seem unsure
> 
> They laugh and scream between mine ears anon.
> 
> Now fill my cup not once, not twice, but thrice
> 
> With flag's brim upon my lips I dance
> 
> Let unseen pixies toss their gowns askance
> 
> While I, the Carnal King, indulge my vice.
> 
> With folded boughs, the treants take their leave
> 
> As merry damsels' corsets come undone
> 
> My song doth stir them like a summer breeze
> 
> They fill mine empty cup without reprieve.
> 
> The sun becomes the moon becomes the sun
> 
> I while away the hours as I please.


```statblock
"name": "Satyr"
"size": "Medium"
"type": "fey"
"alignment": "Chaotic Neutral"
"ac": !!int "14"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "31"
"hit_dice": "7d8"
"stats":
- !!int "12"
- !!int "16"
- !!int "11"
- !!int "12"
- !!int "10"
- !!int "14"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "2"
  "Performance": !!int "6"
"senses": "passive Perception 12"
"languages": "Common, Elvish, Sylvan"
"cr": "1/2"
"traits":
- "desc": "The satyr has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 1|text(6) (2d4 + 1) bludgeoning damage."
  "name": "Ram"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Shortsword"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Shortbow"
"source":
- "MM"
- "SKT"
- "IMR"
- "CM"
- "WBtW"
- "KftGV"
- "SatO"
- "ToFW"
- "BMT"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Satyr.webp"
```
^statblock

```encounter-table
name: Satyr
creatures:
 - 1: Satyr
```

## Environment

forest