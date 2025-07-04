---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/6
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Conjurer Wizard"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 260, Volo's Guide to Monsters p. 212
---
# [Conjurer Wizard](3-Mechanics\CLI\bestiary\humanoid/conjurer-wizard-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 260, Volo's Guide to Monsters p. 212*  

```statblock
"name": "Conjurer Wizard (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "58"
"hit_dice": "13d8"
"stats":
- !!int "9"
- !!int "14"
- !!int "11"
- !!int "17"
- !!int "12"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "4"
  "Intelligence": !!int "6"
"skillsaves":
  "History": !!int "6"
  "Arcana": !!int "6"
"senses": "passive Perception 11"
"languages": "any four languages"
"cr": "6"
"traits":
- "desc": "The conjurer casts one of the following spells, using Intelligence as the\
    \ spellcasting ability (spell save DC 14):\n\nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n1/day each: [fly](/3-Mechanics/CLI/spells/fly.md), [stinking cloud](/3-Mechanics/CLI/spells/stinking-cloud.md),\
    \ [web](/3-Mechanics/CLI/spells/web.md)\n\n2/day each: [fireball](/3-Mechanics/CLI/spells/fireball.md),\
    \ [mage armor](/3-Mechanics/CLI/spells/mage-armor.md), [unseen servant](/3-Mechanics/CLI/spells/unseen-servant.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The conjurer makes three Arcane Burst attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Spell Attack: dice: d20+8 (+8) to hit, reach 5 ft.\
    \ or range 120 ft., one target. Hit: dice:3d10 + 3|text(19) (3d10 + 3) force\
    \ damage."
  "name": "Arcane Burst"
"bonus_actions":
- "desc": "The conjurer teleports, along with any equipment it is wearing or carrying,\
    \ up to 30 feet to an unoccupied space that it can see. If it instead chooses\
    \ a space within range that is occupied by a willing Small or Medium creature,\
    \ they both teleport, swapping places."
  "name": "Benign Transportation (Recharge 4-6)"
- "desc": "The conjurer magically summons an [air elemental](/3-Mechanics/CLI/bestiary/elemental/air-elemental.md),\
    \ an [earth elemental](/3-Mechanics/CLI/bestiary/elemental/earth-elemental.md),\
    \ a [fire elemental](/3-Mechanics/CLI/bestiary/elemental/fire-elemental.md), or\
    \ a [water elemental](/3-Mechanics/CLI/bestiary/elemental/water-elemental.md).\
    \ The elemental appears in an unoccupied space within 60 feet of the conjurer,\
    \ whom it obeys. It takes its turn immediately after the conjurer. It lasts for\
    \ 1 hour, until it or the conjurer dies, or until the conjurer dismisses it as\
    \ a bonus action."
  "name": "Summon Elemental (1/Day)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Conjurer%20Wizard.webp"
```
^statblock

```encounter-table
name: Conjurer Wizard
creatures:
 - 1: Conjurer Wizard
```

## Environment

urban