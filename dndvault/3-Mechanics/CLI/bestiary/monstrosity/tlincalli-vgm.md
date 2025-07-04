---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/desert
- monster/size/large
- monster/type/monstrosity
aliases: ["Tlincalli"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 193
---
# [Tlincalli](3-Mechanics\CLI\bestiary\monstrosity/tlincalli-vgm.md)
*Source: Volo's Guide to Monsters p. 193*  

Tlincallis, also called scorpion folk, are chitin-covered creatures, humanoid from the waist up with the lower body of an enormous scorpion, complete with a stinger at the end of a long tail.

## Desert Nomads

Tlincallis live austerely. They range across arid lands, hunting at dawn and dusk. In the hours between, they wait out the day's heat or the night's cold by burying themselves in loose sand or earth or, if the terrain proves too inflexible, lurking in ruins or shallow caves. A tribe of tlincallis stays in one place for only as long as the hunting is good in the immediate area, though they might visit the same way stations over and over during their wanderings. The tribe also settles down temporarily whenever it's time to lay eggs and hatch a new brood of young.

## Poisonous Eggs

Tlincallis deposit their eggs in warm places out of direct sunlight, often amid a stand of cacti near their present encampment. There the eggs lie protected by hard shells coated in paralytic poison similar to that produced by their stingers. A would-be predator that dares to break an egg is defenseless against the tlincallis that come to investigate.

## Horrid Kidnappers

Tlincallis eat what they kill, but they also take some of their prey alive when they have new mouths to feed. After using their stingers to paralyze victims and their spiked chains to bind them, tlincallis take these prisoners back to their encampment and tie them to cactus or rock formations. There, victims wait until the sun sets and the newly hatched young emerge from the lair to eat them alive.

## Prideful Hunters

Tlincallis see themselves as great hunters. If a tlincalli tribe encounters a more powerful hunter, such as a blue dragon, the tribe's leader must decide whether the group becomes obedient to the superior hunter, moves on, or fights to the death to defeat it.

## Makeshift Weapons and Objects

Tlincallis are uncivilized and don't build cities, make clothing, or mine metals. Instead, they scavenge what they need or want. They do, however, know how to melt down scavenged metal to forge crude weapons and tools.

```statblock
"name": "Tlincalli (VGM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "85"
"hit_dice": "10d10 + 30"
"stats":
- !!int "16"
- !!int "13"
- !!int "16"
- !!int "8"
- !!int "12"
- !!int "8"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "4"
  "Survival": !!int "4"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Tlincalli"
"cr": "5"
"actions":
- "desc": "The tlincalli makes two attacks: one with its longsword or spiked chain,\
    \ and one with its sting."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) slashing damage, or dice:1d10 + 3|text(8)\
    \ (1d10 + 3) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage, and the target is\
    \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 11) if\
    \ it is a Large or smaller creature. Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the tlincalli can't use the spiked chain against another target."
  "name": "Spiked Chain"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage plus dice:4d6|text(14)\
    \ (4d6) poison damage, and the target must succeed on a DC 14 Constitution saving\
    \ throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for 1\
    \ minute. If it fails the saving throw by 5 or more, the target is also [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ while [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned). The target\
    \ can repeat the saving throw at the end of each of its turns, ending the effect\
    \ on itself on a success."
  "name": "Sting"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Tlincalli.webp"
```
^statblock

```encounter-table
name: Tlincalli
creatures:
 - 1: Tlincalli
```

## Environment

desert