---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/3
- monster/environment/coastal
- monster/size/medium
- monster/type/fiend/yugoloth
aliases: ["Merrenoloth"]
NoteIcon: monster
BestiaryType: fiend (yugoloth)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 250
---
# [Merrenoloth](3-Mechanics\CLI\bestiary\fiend/merrenoloth-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 250*  

## Merrenoloth

The grim, gaunt captains of the ferries on the River Styx, merrenoloths have total command of their vessels, ensuring that their passengers reach their destinations safely. Sometimes merrenoloths can be coaxed away from the Lower Planes to captain other vessels, affording those ships and crews the same protection.

Whenever a merrenoloth takes on a contract to captain a ship, it bonds with the vehicle to make sure nothing goes awry with it during the journey. A merrenoloth can navigate its ship safely through the worst storms, always stays on course, and never runs afoul of the myriad hazards that can thwart lesser captains.

A merrenoloth can hold its own in a fight, but it prefers to avoid combat when possible. In fact, it typically specifies in its contracts that it is under no obligation to fight. A merrenoloth's first duty is always to its vessel.

## Yugoloths

Mercenaries that ply their trade throughout the Lower Planes and in other realms, yugoloths have a reputation for effectiveness that is matched only by their desire for ever more wealth. Although yugoloths aren't especially loyal and typically try to exploit every potential loophole in a contract, they undertake any task for which they are hired, no matter how despicable. Yugoloths come in a wide variety of forms, including those described in the Monster Manual and the six creatures presented here.

```statblock
"name": "Merrenoloth (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "yugoloth"
"alignment": "Neutral Evil"
"ac": !!int "13"
"hp": !!int "40"
"hit_dice": "9d8"
"stats":
- !!int "8"
- !!int "17"
- !!int "10"
- !!int "17"
- !!int "14"
- !!int "11"
"speed": "30 ft., swim 40 ft."
"saves":
  "Dexterity": !!int "5"
  "Intelligence": !!int "5"
"skillsaves":
  "Nature": !!int "5"
  "Perception": !!int "4"
  "History": !!int "5"
  "Survival": !!int "4"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 60 ft., passive Perception 14"
"languages": "Abyssal, Infernal, telepathy 60 ft."
"cr": "3"
"traits":
- "desc": "The merrenoloth's innate spellcasting ability is Intelligence (spell save\
    \ DC 13). It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [gust of wind](/3-Mechanics/CLI/spells/gust-of-wind.md)\n\n1/day: [control\
    \ weather](/3-Mechanics/CLI/spells/control-weather.md)\n\n3/day: [control\
    \ water](/3-Mechanics/CLI/spells/control-water.md)"
  "name": "Innate Spellcasting"
- "desc": "The merrenoloth has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "The merrenoloth's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "As a bonus action, the merrenoloth magically teleports, along with any\
    \ equipment it is wearing or carrying, up to 60 feet to an unoccupied space it\
    \ can see."
  "name": "Teleport"
"actions":
- "desc": "The merrenoloth uses Fear Gaze once and makes one oar attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 3|text(8) (2d4 + 3) slashing damage."
  "name": "Oar"
- "desc": "The merrenoloth targets one creature it can see within 60 feet of it. The\
    \ target must succeed on a DC 13 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. The [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Fear Gaze"
"lair_actions":
- "desc": "Any ship a merrenoloth is contracted to captain becomes the creature's\
    \ lair. When fighting on the ship, the merrenoloth can invoke its ability to take\
    \ lair actions. On initiative count 20 (losing initiative ties), the merrenoloth\
    \ can take one lair action to cause one of the following effects; it can't use\
    \ the same effect two rounds in a row:"
  "name": ""
- "desc": "- The ship regains dice:4d10|text(22) (4d10) hit points.  \n- A strong\
    \ wind propels the ship, increasing its speed by 30 feet until initiative count\
    \ 20 on the next round.  \n- The air within 60 feet of the ship is filled with\
    \ howling wind. Until initiative count 20 on the next round, that area is difficult\
    \ terrain, and when a Medium or smaller creature flies into that area or starts\
    \ its turn flying there, it must succeed on a DC 13 Strength saving throw or be\
    \ knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).  "
  "name": ""
"regional_effects":
- "desc": "A merrenoloth imbues its vessel with powerful magic that creates one or\
    \ more of the following effects:"
  "name": ""
- "desc": "- The ship doesn't sink even if its hull is breached.  \n- The ship always\
    \ stays on course to the destination the merrenoloth names.  \n- Creatures the\
    \ merrenoloth chooses to take on the ship aren't discomfited by wind or weather,\
    \ though this effect doesn't protect against damage.  "
  "name": ""
- "desc": "If the merrenoloth dies, these effects fade over the course of dice: 1d6|avg|noform\
    \ (1d6) hours."
  "name": ""
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Merrenoloth.webp"
```
^statblock

```encounter-table
name: Merrenoloth
creatures:
 - 1: Merrenoloth
```

## Environment

coastal