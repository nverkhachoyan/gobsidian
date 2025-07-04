---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1
- monster/environment/arctic
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/humanoid/gnoll
aliases: ["Gnoll Flesh Gnawer"]
NoteIcon: monster
BestiaryType: humanoid (gnoll)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 154
---
# [Gnoll Flesh Gnawer](3-Mechanics\CLI\bestiary\humanoid/gnoll-flesh-gnawer-vgm.md)
*Source: Volo's Guide to Monsters p. 154*  

If any group of gnolls could be said to be more feral than the others, that distinction would go to the flesh gnawers. These gnolls eschew the use of ranged weapons in favor of short blades that they wield with speed and efficiency. In the thick of a fight, they are capable of dashing across the field, slashing and snarling as they run down stragglers and finish off wounded foes.

```statblock
"name": "Gnoll Flesh Gnawer (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "gnoll"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "22"
"hit_dice": "4d8 + 4"
"stats":
- !!int "12"
- !!int "14"
- !!int "12"
- !!int "8"
- !!int "10"
- !!int "8"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "4"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Gnoll"
"cr": "1"
"traits":
- "desc": "When the gnoll reduces a creature to 0 hit points with a melee attack on\
    \ its turn, the gnoll can take a bonus action to move up to half its speed and\
    \ make a bite attack."
  "name": "Rampage"
"actions":
- "desc": "The gnoll makes three attacks: one with its bite and two with its shortsword."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Shortsword"
- "desc": "Until the end of the turn, the gnoll's speed increases by 60 feet and it\
    \ doesn't provoke opportunity attacks."
  "name": "Sudden Rush"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Gnoll%20Flesh%20Gnawer.webp"
```
^statblock

```encounter-table
name: Gnoll Flesh Gnawer
creatures:
 - 1: Gnoll Flesh Gnawer
```

## Environment

grassland, forest, hill, arctic