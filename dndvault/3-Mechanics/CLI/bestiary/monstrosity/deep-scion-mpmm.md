---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/coastal
- monster/environment/underwater
- monster/size/medium
- monster/type/monstrosity
aliases: ["Deep Scion"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 88, Volo's Guide to Monsters p. 135
---
# [Deep Scion](3-Mechanics\CLI\bestiary\monstrosity/deep-scion-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 88, Volo's Guide to Monsters p. 135*  

```statblock
"name": "Deep Scion (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "11"
"hp": !!int "67"
"hit_dice": "9d8 + 27"
"stats":
- !!int "18"
- !!int "13"
- !!int "16"
- !!int "10"
- !!int "12"
- !!int "14"
"speed": "30 ft. (20 ft. and swim 40 ft. in hybrid form)"
"saves":
  "Charisma": !!int "4"
  "Wisdom": !!int "3"
"skillsaves":
  "Sleight of Hand": !!int "3"
  "Deception": !!int "6"
  "Stealth": !!int "3"
  "Insight": !!int "3"
"senses": "darkvision 120 ft., passive Perception 11"
"languages": "Aquan, Common, thieves' cant"
"cr": "3"
"traits":
- "desc": "The deep scion can breathe air and water."
  "name": "Amphibious (Hybrid Form Only)"
"actions":
- "desc": "The deep scion makes two Battleaxe attacks, or it makes one Bite attack\
    \ and two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage, or dice:1d10 + 4|text(9)\
    \ (1d10 + 4) slashing damage if used with two hands."
  "name": "Battleaxe"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d4 + 4|text(6) (1d4 + 4) piercing damage."
  "name": "Bite (Hybrid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) slashing damage."
  "name": "Claw (Hybrid Form Only)"
- "desc": "The deep scion emits a terrible scream audible within 300 feet. Creatures\
    \ within 30 feet of the deep scion must succeed on a DC 13 Wisdom saving throw\
    \ or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) until the end\
    \ of the deep scion's next turn. In water, the psychic screech also telepathically\
    \ transmits the deep scion's memories of the last 24 hours to its master, regardless\
    \ of distance, so long as it and its master are in the same body of water."
  "name": "Psychic Screech (Hybrid Form Only; Recharges after a Short or Long Rest)"
"bonus_actions":
- "desc": "The deep scion transforms into a hybrid form (humanoid-piscine) or back\
    \ into its true form, which is humanlike. Its statistics, other than its speed,\
    \ are the same in each form. Any equipment it is wearing or carrying isn't transformed.\
    \ The deep scion reverts to its true form if it dies."
  "name": "Change Shape"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Deep%20Scion.webp"
```
^statblock

```encounter-table
name: Deep Scion
creatures:
 - 1: Deep Scion
```

## Environment

coastal, underwater