---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/12
- monster/environment/forest
- monster/environment/mountain
- monster/environment/swamp
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Archdruid"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 210
---
# [Archdruid](3-Mechanics\CLI\bestiary\humanoid/archdruid-vgm.md)
*Source: Volo's Guide to Monsters p. 210*  

Archdruids watch over the natural wonders of their domains. They seldom interact with civilized folk unless there is a great threat to the natural order. An archdruid typically has one or more pupils who are druids (see the Monster Manual for statistics), and the archdruid's lair is usually guarded by loyal beasts and fey creatures.

```statblock
"name": "Archdruid (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "16"
"ac_class": "[hide armor](/3-Mechanics/CLI/items/hide-armor.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "132"
"hit_dice": "24d8 + 24"
"stats":
- !!int "10"
- !!int "14"
- !!int "12"
- !!int "12"
- !!int "20"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "9"
  "Intelligence": !!int "5"
"skillsaves":
  "Medicine": !!int "9"
  "Nature": !!int "5"
  "Perception": !!int "9"
"senses": "passive Perception 19"
"languages": "Druidic plus any two languages"
"cr": "12"
"traits":
- "desc": "The archdruid is an 18th-level spellcaster. Its spellcasting ability is\
    \ Wisdom (spell save DC 17, dice: d20+9 (+9) to hit with spell attacks). It\
    \ has the following druid spells prepared:\n\nCantrips (at will): [druidcraft](/3-Mechanics/CLI/spells/druidcraft.md),\
    \ [mending](/3-Mechanics/CLI/spells/mending.md), [poison spray](/3-Mechanics/CLI/spells/poison-spray.md),\
    \ [produce flame](/3-Mechanics/CLI/spells/produce-flame.md)\n\n1st level (4\
    \ slots): [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md), [entangle](/3-Mechanics/CLI/spells/entangle.md),\
    \ [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md), [speak with animals](/3-Mechanics/CLI/spells/speak-with-animals.md)\n\
    \n2nd level (3 slots): [animal messenger](/3-Mechanics/CLI/spells/animal-messenger.md),\
    \ [beast sense](/3-Mechanics/CLI/spells/beast-sense.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md)\n\
    \n3rd level (3 slots): [conjure animals](/3-Mechanics/CLI/spells/conjure-animals.md),\
    \ [meld into stone](/3-Mechanics/CLI/spells/meld-into-stone.md), [water breathing](/3-Mechanics/CLI/spells/water-breathing.md)\n\
    \n4th level (3 slots): [dominate beast](/3-Mechanics/CLI/spells/dominate-beast.md),\
    \ [locate creature](/3-Mechanics/CLI/spells/locate-creature.md), [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md),\
    \ [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)\n\n5th level (3 slots):\
    \ [commune with nature](/3-Mechanics/CLI/spells/commune-with-nature.md), [mass\
    \ cure wounds](/3-Mechanics/CLI/spells/mass-cure-wounds.md), [tree stride](/3-Mechanics/CLI/spells/tree-stride.md)\n\
    \n6th level (1 slots): [heal](/3-Mechanics/CLI/spells/heal.md), [heroes' feast](/3-Mechanics/CLI/spells/heroes-feast.md),\
    \ [sunbeam](/3-Mechanics/CLI/spells/sunbeam.md)\n\n7th level (1 slots): [fire\
    \ storm](/3-Mechanics/CLI/spells/fire-storm.md)\n\n8th level (1 slots): [animal\
    \ shapes](/3-Mechanics/CLI/spells/animal-shapes.md)\n\n9th level (1 slots):\
    \ [foresight](/3-Mechanics/CLI/spells/foresight.md)"
  "name": "Spellcasting"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) slashing damage."
  "name": "Scimitar"
- "desc": "The archdruid magically polymorphs into a beast or elemental with a challenge\
    \ rating of 6 or less, and can remain in this form for up to 9 hours. The archdruid\
    \ can choose whether its equipment falls to the ground, melds with its new form,\
    \ or is worn by the new form. The archdruid reverts to its true form if it dies\
    \ or falls [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious). The\
    \ archdruid can revert to its true form using a bonus action on its turn.\n\n\
    While in a new form, the archdruid retains its game statistics and ability to\
    \ speak, but its AC, movement modes, Strength, and Dexterity are replaced by those\
    \ of the new form, and it gains any special senses, proficiencies, traits, actions,\
    \ and reactions (except class features, legendary actions, and lair actions) that\
    \ the new form has but that it lacks. It can cast its spells with verbal or somatic\
    \ components in its new form.\n\nThe new form's attacks count as magical for the\
    \ purpose of overcoming resistances and immunity to nonmagical attacks."
  "name": "Change Shape (2/Day)"
"source":
- "VGM"
- "WDMM"
- "GoS"
- "MOT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Archdruid.webp"
```
^statblock

```encounter-table
name: Archdruid
creatures:
 - 1: Archdruid
```

## Environment

forest, mountain, swamp