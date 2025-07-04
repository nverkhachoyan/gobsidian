---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/13
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Narzugon"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 167
---
# [Narzugon](3-Mechanics\CLI\bestiary\fiend/narzugon-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 167*  

## Narzugon

Paladins who make deals with devils and carry their twisted sense of honor into the afterlife are especially valuable to the archdukes of the Nine Hells, who want unquestioning champions to lead their legions in war. These narzugons, wielding lances of hellfire and riding nightmare steeds like horrific perversions of knights errant, roam across the infernal layers and other planes to carry out the will of their masters.

### Death in Hellfire

 A narzugon's lances are forged in hellfire. The soul of anyone killed by such a lance is shunted to the River Styx for rebirth as a lemure. Each lance is unique to its owner, bearing the marks of both the narzugon and its master.

## Nightmare Riders

Each narzugon claims a nightmare as its mount. These nightmares are bound by [infernal tack](/3-Mechanics/CLI/items/infernal-tack-mtf.md) and must respond to summons and commands from the wearer of the spurs.

```statblock
"name": "Narzugon (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
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
  "Perception": !!int "7"
"damage_resistances": "acid; cold; bludgeoning, piercing, slashing from nonmagical\
  \ attacks that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 17"
"languages": "Common, Infernal, telepathy 120 ft."
"cr": "13"
"traits":
- "desc": "The narzugon has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks made to perceive good-aligned creatures."
  "name": "Diabolical Sense"
- "desc": "The narzugon wears spurs that are part of infernal tack, which allow it\
    \ to summon its [nightmare](/3-Mechanics/CLI/bestiary/fiend/nightmare.md) companion."
  "name": "Infernal Tack"
- "desc": "The narzugon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The narzugon uses its Infernal Command or Terrifying Command. It also makes\
    \ three hellfire lance attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 10 ft., one\
    \ target. Hit: dice:1d12 + 5|text(11) (1d12 + 5) piercing damage plus dice:3d10|text(16)\
    \ (3d10) fire damage. If this damage kills a creature, the creature's soul rises\
    \ from the River Styx as a lemure in Avernus in dice: 1d4|avg|noform (1d4)\
    \ hours.\n\nIf the creature isn't revived before then, only a [wish](/3-Mechanics/CLI/spells/wish.md)\
    \ spell or killing the lemure and casting true resurrection on the creature's\
    \ original body can restore it to life. Constructs and devils are immune to this\
    \ effect."
  "name": "Hellfire Lance"
- "desc": "Each ally of the narzugon within 60 feet of it can't be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) until the end\
    \ of the narzugon's next turn."
  "name": "Infernal Command"
- "desc": "Each creature that isn't a fiend within 60 feet of the narzugon that can\
    \ hear it must succeed on a DC 17 Charisma saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of it for 1 minute.\n\nA creature can repeat the saving throw at the end of\
    \ each of its turns, ending the effect on itself on a success. A creature that\
    \ makes a successful saving throw is immune to this narzugon's Terrifying Command\
    \ for 24 hours."
  "name": "Terrifying Command"
- "desc": "The narzugon, or one creature it touches, regains up to 100 hit points."
  "name": "Healing (1/Day)"
"source":
- "MTF"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Narzugon.webp"
```
^statblock

```encounter-table
name: Narzugon
creatures:
 - 1: Narzugon
```