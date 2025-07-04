---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/2
- monster/environment/desert
- monster/size/medium
- monster/type/aberration
aliases: ["Berbalang"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 120
---
# [Berbalang](3-Mechanics\CLI\bestiary\aberration/berbalang-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 120*  

## Berbalang

Berbalangs creep across the petrified remains of dead gods adrift on the Astral Plane. Obsessed with gathering secrets, both from the gods they inhabit and from the bones of dead creatures, they call forth the spirits of the dead and force them to divulge what they learned in life.

## Speakers of the Dead

Berbalangs prefer to speak only to dead things, and specifically only to the spirits they call forth in the hope of learning secrets. They record their stories on the bones that once belonged to these creatures, thus preserving the information they gain.

## Spectral Spy

The pursuit of knowledge drives everything berbalangs do. Although they mostly learn their secrets from the dead, they aren't above spying on the living to take knowledge from them as well. A berbalang can create a spectral duplicate of itself and send the duplicate out to gather information on other planes by watching places where the gods and their servants gather. When a berbalang is perceiving its environment through its duplicate, its actual body is [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious) and can't protect or nourish itself. Thus, a berbalang typically uses its duplicate for only a short time before returning its consciousness to its body.

## Weird Oracles

The knowledge that berbalangs accumulate makes them great sources of information for powerful people traveling the planes. Berbalangs ignore petitioners, however, unless they come bearing a choice secret or the bones of a particularly interesting creature. Githyanki have found a way to coexist with berbalangs, and sometimes use the creatures to spy on their enemies and to watch over their crèches on the Material Plane.

```statblock
"name": "Berbalang (MTF)"
"size": "Medium"
"type": "aberration"
"alignment": "Neutral Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "38"
"hit_dice": "11d8 - 11"
"stats":
- !!int "9"
- !!int "16"
- !!int "9"
- !!int "17"
- !!int "11"
- !!int "10"
"speed": "30 ft., fly 40 ft."
"saves":
  "Dexterity": !!int "5"
  "Intelligence": !!int "5"
"skillsaves":
  "Religion": !!int "5"
  "Insight": !!int "2"
  "Perception": !!int "2"
  "History": !!int "5"
  "Arcana": !!int "5"
"senses": "truesight 120 ft., passive Perception 12"
"languages": "all but rarely speaks"
"cr": "2"
"traits":
- "desc": "The berbalang's innate spellcasting ability is Intelligence (spell save\
    \ DC 13). The berbalang can innately cast the following spells, requiring no material\
    \ components:\n\nAt will: [speak with dead](/3-Mechanics/CLI/spells/speak-with-dead.md)\n\
    \n1/day: [plane shift](/3-Mechanics/CLI/spells/plane-shift.md) (self only)"
  "name": "Innate Spellcasting"
- "desc": "As a bonus action, the berbalang creates one spectral duplicate of itself\
    \ in an unoccupied space it can see within 60 feet of it. While the duplicate\
    \ exists, the berbalang is [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious).\
    \ A berbalang can have only one duplicate at a time. The duplicate disappears\
    \ when it or the berbalang drops to 0 hit points or when the berbalang dismisses\
    \ it (no action required).\n\nThe duplicate has the same statistics and knowledge\
    \ as the berbalang, and everything experienced by the duplicate is known by the\
    \ berbalang. All damage dealt by the duplicate's attacks is psychic damage."
  "name": "Spectral Duplicate (Recharges after a Short or Long Rest)"
"actions":
- "desc": "The berbalang makes two attacks: one with its bite and one with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 3|text(8) (2d4 + 3) slashing damage."
  "name": "Claws"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Berbalang.webp"
```
^statblock

```encounter-table
name: Berbalang
creatures:
 - 1: Berbalang
```

## Environment

desert