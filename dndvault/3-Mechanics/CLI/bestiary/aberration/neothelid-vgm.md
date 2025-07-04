---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/13
- monster/environment/underdark
- monster/size/gargantuan
- monster/type/aberration
aliases: ["Neothelid"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 181
---
# [Neothelid](3-Mechanics\CLI\bestiary\aberration/neothelid-vgm.md)
*Source: Volo's Guide to Monsters p. 181*  

A slime-covered worm of immense size, a neothelid is the result of the mind flayer reproductive cycle gone horribly wrong. On rare occasions, an illithid colony collapses, typically after an external assault, and the elder brain is killed. When that happens, the colony's tadpoles are suddenly freed from their fate. They no longer serve as food, and in turn are no longer fed by their caretakers. Driven by hunger, they turn to devouring one another. Only one tadpole survives out of the thousands in the colony's pool, and it emerges as a neothelid.

## Abhorrent to Illithids

Among the strongest taboos in illithid society is the idea of allowing a mature tadpole to survive without implanting it into a donor brain. Under normal circumstances, any tadpole that grows larger than a few inches in length is killed by the elder brain to be food for it or for less mature tadpoles. Any tadpole that survives beyond that state is perceived as a threat to the colony, and the mind flayers organize hunting parties to exterminate the abomination. Lacking enough intelligence to be detected by an elder brain's power to sense thoughts, neothelids warrant such precautions.

## Savage Behemoth

As a feral thing, a neothelid knows nothing beyond the predatory existence it has lived so far and struggles to comprehend its new psionic abilities. Neothelids prowl subterranean passages in search of more brains to sate their constant hunger, growing ever more vicious. These creatures can spray tissue-dissolving enzymes from their tentacle ducts, reducing victims to a puddle of slime and leaving only the pulsing brain unharmed. They have no knowledge of their link to illithids, so they're just as likely to prey on mind flayers as on anything else.

```statblock
"name": "Neothelid (VGM)"
"size": "Gargantuan"
"type": "aberration"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "325"
"hit_dice": "21d20 + 105"
"stats":
- !!int "27"
- !!int "7"
- !!int "21"
- !!int "3"
- !!int "16"
- !!int "12"
"speed": "30 ft."
"saves":
  "Charisma": !!int "6"
  "Wisdom": !!int "8"
  "Intelligence": !!int "1"
"senses": "blindsight 120 ft., passive Perception 13"
"languages": ""
"cr": "13"
"traits":
- "desc": "The neothelid's innate spellcasting ability is Wisdom (spell save DC 16).\
    \ It can innately cast the following spells, requiring no components:\n\nAt\
    \ will: [levitate](/3-Mechanics/CLI/spells/levitate.md)\n\n1/day each: [confusion](/3-Mechanics/CLI/spells/confusion.md),\
    \ [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md), [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "The neothelid is aware of the presence of creatures within 1 mile of it\
    \ that have an Intelligence score of 4 or higher. It knows the distance and direction\
    \ to each creature, as well as each creature's Intelligence score, but can't sense\
    \ anything else about it. A creature protected by a [mind blank](/3-Mechanics/CLI/spells/mind-blank.md)\
    \ spell, a [nondetection](/3-Mechanics/CLI/spells/nondetection.md) spell, or similar\
    \ magic can't be perceived in this manner."
  "name": "Creature Sense"
- "desc": "The neothelid has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 15 ft., one\
    \ target. Hit: dice:3d8 + 8|text(21) (3d8 + 8) bludgeoning damage plus dice:3d8|text(13)\
    \ (3d8) psychic damage. If the target is a Large or smaller creature, it must\
    \ succeed on a DC 18 Strength saving throw or be swallowed by the neothelid. A\
    \ swallowed creature is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ and [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), it has total\
    \ cover against attacks and other effects outside the neothelid, and it takes\
    \ dice:10d6|text(35) (10d6) acid damage at the start of each of the neothelid's\
    \ turns.\n\nIf the neothelid takes 30 damage or more on a single turn from a creature\
    \ inside it, the neothelid must succeed on a DC 18 Constitution saving throw at\
    \ the end of that turn or regurgitate all swallowed creatures, which fall [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ in a space within 10 feet of the neothelid. If the neothelid dies, a swallowed\
    \ creature is no longer [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by it and can escape from the corpse by using 20 feet of movement, exiting [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Tentacles"
- "desc": "The neothelid exhales acid in a 60-foot cone. Each creature in that area\
    \ must make a DC 18 Dexterity saving throw, taking dice:10d6|text(35) (10d6)\
    \ acid damage on a failed save, or half as much damage on a successful one."
  "name": "Acid Breath (Recharge 5-6)"
"source":
- "VGM"
- "WDMM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Neothelid.webp"
```
^statblock

```encounter-table
name: Neothelid
creatures:
 - 1: Neothelid
```

## Environment

underdark