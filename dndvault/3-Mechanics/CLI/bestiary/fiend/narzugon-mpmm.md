---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/13
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Narzugon"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 190, Mordenkainen's Tome of Foes p. 167
---
# [Narzugon](3-Mechanics\CLI\bestiary\fiend/narzugon-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 190, Mordenkainen's Tome of Foes p. 167*  

```statblock
"name": "Narzugon (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Typically  Lawful Evil"
"ac": !!int "20"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "112"
"hit_dice": "15d8 + 45"
"stats":
- !!int "20"
- !!int "10"
- !!int "17"
- !!int "16"
- !!int "14"
- !!int "19"
"speed": "30 ft."
"saves":
  "Charisma": !!int "9"
  "Dexterity": !!int "5"
  "Constitution": !!int "8"
"skillsaves":
  "Perception": !!int "12"
"damage_resistances": "acid; cold; bludgeoning, piercing, slashing from nonmagical\
  \ attacks that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 22"
"languages": "Common, Infernal, telepathy 120 ft."
"cr": "13"
"traits":
- "desc": "The narzugon wears spurs that are part of [infernal tack](/3-Mechanics/CLI/items/infernal-tack-mtf.md),\
    \ which allow it to summon its [nightmare](/3-Mechanics/CLI/bestiary/fiend/nightmare.md)\
    \ companion as an action."
  "name": "Infernal Tack"
- "desc": "The narzugon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The narzugon makes three Hellfire Lance attacks. It also uses Infernal\
    \ Command or Terrifying Command."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 10 ft., one\
    \ target. Hit: dice:1d12 + 5|text(11) (1d12 + 5) piercing damage plus dice:3d10|text(16)\
    \ (3d10) fire damage. If this damage kills a creature with a soul, the soul\
    \ rises from the River Styx as a [lemure](/3-Mechanics/CLI/bestiary/fiend/lemure.md)\
    \ in Avernus in dice: 1d4|avg|noform (1d4) hours. If the creature isn't revived\
    \ before then, only a [wish](/3-Mechanics/CLI/spells/wish.md) spell or killing\
    \ the lemure and casting true resurrection on the creature's original body can\
    \ restore it to life. Constructs and devils are immune to this effect."
  "name": "Hellfire Lance"
- "desc": "Each ally of the narzugon within 60 feet of it can't be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) until the end\
    \ of the narzugon's next turn."
  "name": "Infernal Command"
- "desc": "Each creature within 60 feet of the narzugon that isn't a Fiend must succeed\
    \ on a DC 17 Charisma saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the narzugon for 1 minute. A creature can repeat the saving throw at the\
    \ end of each of its turns, ending the effect on itself on a success. A creature\
    \ that makes a successful saving throw is immune to this narzugon's Terrifying\
    \ Command for 24 hours."
  "name": "Terrifying Command"
- "desc": "The narzugon, or one creature it touches, regains 100 hit points."
  "name": "Healing (1/Day)"
"source":
- "MPMM"
- "MTF"
- "AATM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Narzugon.webp"
```
^statblock

```encounter-table
name: Narzugon
creatures:
 - 1: Narzugon
```