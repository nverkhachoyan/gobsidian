---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/9
- monster/environment/desert
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/cleric
aliases: ["War Priest"]
NoteIcon: monster
BestiaryType: humanoid (cleric)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 254, Volo's Guide to Monsters p. 218
---
# [War Priest](3-Mechanics\CLI\bestiary\humanoid/war-priest-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 254, Volo's Guide to Monsters p. 218*  

```statblock
"name": "War Priest (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "cleric"
"alignment": "Any alignment"
"ac": !!int "18"
"ac_class": "[plate](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "117"
"hit_dice": "18d8 + 36"
"stats":
- !!int "16"
- !!int "10"
- !!int "14"
- !!int "11"
- !!int "17"
- !!int "13"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "7"
  "Constitution": !!int "6"
"skillsaves":
  "Intimidation": !!int "5"
  "Religion": !!int "4"
"senses": "passive Perception 13"
"languages": "any two languages"
"cr": "9"
"traits":
- "desc": "The war priest casts one of the following spells, using Wisdom as the spellcasting\
    \ ability (spell save DC 15):\n\nAt will: [light](/3-Mechanics/CLI/spells/light.md),\
    \ [spare the dying](/3-Mechanics/CLI/spells/spare-the-dying.md), [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\
    \n1/day each: [banishment](/3-Mechanics/CLI/spells/banishment.md), [command](/3-Mechanics/CLI/spells/command.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [flame strike](/3-Mechanics/CLI/spells/flame-strike.md),\
    \ [guardian of faith](/3-Mechanics/CLI/spells/guardian-of-faith.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [lesser restoration](/3-Mechanics/CLI/spells/lesser-restoration.md), [revivify](/3-Mechanics/CLI/spells/revivify.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The war priest makes two Maul attacks, and it uses Holy Fire."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) bludgeoning damage  plus Hit: dice:3d6|text(10)\
    \ (3d6) radiant damage."
  "name": "Maul"
- "desc": "The war priest targets one creature it can see within 60 feet of it. The\
    \ target must make a DC 15 Wisdom saving throw. On a failed save, the target takes\
    \ dice:2d8 + 3|text(12) (2d8 + 3) radiant damage, and it is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ until the start of the war priest's next turn. On a successful save, the target\
    \ takes half as much damage and isn't [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)."
  "name": "Holy Fire"
"bonus_actions":
- "desc": "The war priest or one creature of its choice within 60 feet of it regains\
    \ dice:2d8 + 3|text(12) (2d8 + 3) hit points."
  "name": "Healing Light (Recharge 4-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/War%20Priest.webp"
```
^statblock

```encounter-table
name: War Priest
creatures:
 - 1: War Priest
```

## Environment

desert, urban