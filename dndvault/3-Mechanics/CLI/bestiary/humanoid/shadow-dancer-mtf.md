---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/7
- monster/environment/forest
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Shadow Dancer"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 225
---
# [Shadow Dancer](3-Mechanics\CLI\bestiary\humanoid/shadow-dancer-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 225*  

## Shadow Dancer

Those who have fought shadow dancers describe the experience as similar to fighting a living darkness. Every dim alcove and darkened nook is a place from where the lithe and acrobatic shadow dancers can emerge to ambush their prey. Using this tactic, they attack their enemies from all angles with a flurry of entangling chains that hold fast and corrupt the flesh. When their quarry is helpless, others move in to help dispatch the prey. Then they loot the corpse for trinkets, anything colorful and lively to gaze at after they return to the gloom of the Shadowfell.

## Shadar-kai

In the perpetual gloom of the Shadowfell lives a society that serves the Raven Queen. They were brought into that dusky realm in ages past, so long ago that they're now perfectly adapted to that cheerless environment, both physically and mentally.

## Soul Custodians

Shadar-kai watch over both the Shadowfell and the material world, scouting out choice souls and tragedies that might please their deity. They are rumored to be able to coax worldly events along tragic paths for her amusement. The Raven Queen is famously cryptic even to her most devoted followers, however. The shadar-kai's efforts are rewarded only with vague omens they interpret as best they can.

## Blighted Elves

Shadar-kai were once elves, but eons of exposure to the debilitating influence of the Shadowfell has left them joyless and mournful. In that realm, they have the appearance of withered elves: pale hair, wrinkled gray skin, and swollen joints give them a corpselike aspect. They appear more youthful while on other planes, but their skin always retains its deathly pallor. They dress in dark cloaks and heavy veils, detest mirrors, and avoid keeping things that remind them of their age.

```statblock
"name": "Shadow Dancer (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Neutral"
"ac": !!int "15"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "71"
"hit_dice": "13d8 + 13"
"stats":
- !!int "12"
- !!int "16"
- !!int "13"
- !!int "11"
- !!int "12"
- !!int "12"
"speed": "30 ft."
"saves":
  "Charisma": !!int "4"
  "Dexterity": !!int "6"
"skillsaves":
  "Stealth": !!int "6"
"damage_resistances": "necrotic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Common, Elvish"
"cr": "7"
"traits":
- "desc": "The shadow dancer has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put it to sleep."
  "name": "Fey Ancestry"
- "desc": "As a bonus action, the shadow dancer can teleport up to 30 feet to an unoccupied\
    \ space it can see. Both the space it teleports from and the space it teleports\
    \ to must be in dim light or darkness. The shadow dancer can use this ability\
    \ between the weapon attacks of another action it takes."
  "name": "Shadow Jump"
"actions":
- "desc": "The shadow dancer makes three spiked chain attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) piercing damage, and the target must\
    \ succeed on a DC 14 Dexterity saving throw or suffer one additional effect of\
    \ the shadow dancer's choice:\n\n- The target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 14) if it is a Medium or smaller creature. Until the grapple ends,\
    \ the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the shadow dancer can't grapple another target.  \n- The target is knocked\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone).  \n- The target takes dice:4d10|text(22)\
    \ (4d10) necrotic damage.  "
  "name": "Spiked Chain"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Shadow%20Dancer.webp"
```
^statblock

```encounter-table
name: Shadow Dancer
creatures:
 - 1: Shadow Dancer
```

## Environment

forest, underdark, urban